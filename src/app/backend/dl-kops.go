package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"os"
	exec "os/exec"
	"os/user"
	"strings"

	"github.com/spf13/viper"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

//Config is used to load yaml config for ssh
type Config struct {
	Info struct {
		Username    string `json:"username"`
		Keyfilepath string `json:"keyfilepath"`
		HostIP      string `json:"hostip"`
	} `json:"info"`
}

//KeyPair creates a keypair with a path file string
func KeyPair(keyFile string) (ssh.AuthMethod, error) {
	pem, err := ioutil.ReadFile(keyFile)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(pem)
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeys(key), nil
}

//Connect establishes a connection betwen a host and the program
func Connect(host string, methods ...ssh.AuthMethod) (*ssh.Client, error) {
	cfg := ssh.ClientConfig{
		User: "SETCloud",
		Auth: methods,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	return ssh.Dial("tcp", host, &cfg)
}

//SSHAgent creates a SSH Agent
func SSHAgent() (ssh.AuthMethod, error) {
	agentSock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, err
	}

	return ssh.PublicKeysCallback(agent.NewClient(agentSock).Signers), nil
}

func main() {

	viperInstance := viper.New()

	viperInstance.SetConfigName("config")
	viperInstance.AddConfigPath(".")
	err := viperInstance.ReadInConfig()
	if err != nil {
		log.Printf("Error in reading config %s", err)
	}

	var c Config
	err = viperInstance.Unmarshal(&c)
	if err != nil {
		log.Printf("Error in deconding config %s", err)
	}

	//Key File location+creation
	keyfilepath := c.Info.Keyfilepath
	if keyfilepath == "" {
		usr, _ := user.Current()
		keyfilepath = usr.HomeDir + os.Getenv("kubekeyconfig")
	}
	//SSH Agent creation
	agent, err := SSHAgent()
	if err != nil {
		log.Printf("Error in creating SSH Agent %s", err)
	}

	//keypair decoding
	keyPair, err := KeyPair(keyfilepath)
	if err != nil {
		log.Printf(" error creating file: %s", err)
	}

	//Client creation
	client, err := Connect(c.Info.HostIP+":22", agent, keyPair)
	if err != nil {
		log.Printf("Error creating client %s", err)
	}
	defer client.Close()

	//Session creation
	session, err := client.NewSession()
	if err != nil {
		log.Fatalf("session failed:%v", err)
	}
	defer session.Close()

	session.Stdout = os.Stdout
	session.Setenv("LS_COLORS", os.Getenv("LS_COLORS"))

	shcmd := "sh"
	var args []string
	var output []string
	args = []string{"./src/app/backend/trial.sh"}
	args = append(args, "128.96.110.19")
	cmd := exec.Command(shcmd, args...)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		log.Fatalf("exec Error: %s", err)
	}
	fmt.Printf("%s", out.String())
	output = append(output, out.String())
	fmt.Println(output)

	err = session.Run("./setup_cluster.sh")

	if err != nil {
		log.Fatalf("Run failed:%v", err)
	}
}
