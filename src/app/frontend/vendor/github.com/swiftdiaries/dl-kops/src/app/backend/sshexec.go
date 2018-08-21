package backend

import (
	"bytes"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/user"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

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

//Connect establishes a connection between a host and the program
func Connect(hostname string, host string, methods ...ssh.AuthMethod) (*ssh.Client, error) {
	cfg := ssh.ClientConfig{
		User: hostname,
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

//ExecuteSSHCommand used to run a command on a host through SSH
func ExecuteSSHCommand(hostname string, hostip string, keyfilepath string, command string) []string {

	//Key File location+creation
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
	client, err := Connect(hostname, hostip+":22", agent, keyPair)
	if err != nil {
		log.Printf("Error creating client %s", err)
	}
	defer client.Close()

	//Session creation
	session, err := client.NewSession()
	if err != nil {
		log.Printf("session failed:%v", err)
	}
	defer session.Close()

	var out bytes.Buffer
	var output []string
	session.Stdout = &out
	session.Setenv("LS_COLORS", os.Getenv("LS_COLORS"))

	err = session.Run(command)

	if err != nil {
		log.Printf("Run failed:%v", err)
	}
	output = append(output, out.String())

	return output
}
