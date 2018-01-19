package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/viper"
)

var (
	gopath = os.Getenv("GOPATH")
	//HomeDir contains the project's home directory
	HomeDir = gopath + "/src/github.com/swiftdiaries/dl-kops/"
)

//ClusterConfig describes a cluster.
type ClusterConfig struct {
	Controller struct {
		Hostname    string `json:"hostname"`
		Hostip      string `json:"hostip"`
		Keyfilepath string `json:"keyfilepath"`
	} `json:"Controller"`
	Worker struct {
		Hostname    string `json:"hostname"`
		Hostip      string `json:"hostip"`
		Keyfilepath string `json:"keyfilepath"`
	} `json:"Worker"`
}

//ExecuteScriptFile executes a script file with the given arguments
func ExecuteScriptFile(filepath string, arguments []string) []string {
	shcmd := "sh"
	var args []string
	var output []string
	args = append(args, filepath)
	for _, argument := range arguments {
		args = append(args, argument)
	}
	fmt.Printf("Args: %s", args)
	cmd := exec.Command(shcmd, args...)
	cmd.Stdin = strings.NewReader("")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatalf("exec Error: %s", err)
	}
	return output
}

//GetCreds gets creds from config
func GetCreds(role string) (string, string, string) {
	viper.SetConfigName("config")
	viper.AddConfigPath(HomeDir)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading in config")
	}
	var cluster ClusterConfig
	err = viper.Unmarshal(&cluster)
	if err != nil {
		fmt.Printf("Error in unmarshalling: %s", err)
	}
	if role == "controller" {
		return cluster.Controller.Hostname, cluster.Controller.Hostip, cluster.Controller.Keyfilepath
	}
	return cluster.Worker.Hostname, cluster.Worker.Hostip, cluster.Worker.Keyfilepath
}

//ResetConfig resets config to defaults
func ResetConfig(w http.ResponseWriter, r *http.Request) {
	hostname := "ubuntu"
	hostip := "0.0.0.1"
	keyfile := "/Users/Name/key.key"
	viper.SetConfigName("config")
	viper.AddConfigPath(HomeDir)
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Printf("Error reading in config")
	}
	var cluster ClusterConfig
	err = viper.Unmarshal(&cluster)
	if err != nil {
		fmt.Printf("Error in unmarshalling: %s", err)
	}
	cluster.Controller.Hostname = hostname
	cluster.Controller.Hostip = hostip
	cluster.Controller.Keyfilepath = keyfile
	cluster.Worker.Hostname = hostname
	cluster.Worker.Hostip = hostip
	cluster.Worker.Keyfilepath = keyfile
	b, err := json.Marshal(cluster)
	if err != nil {
		fmt.Printf("Error in marshalling: %s", err)
	}
	err = ioutil.WriteFile("config.yaml", b, 0644)
	if err != nil {
		fmt.Printf("Error in writing:%s", err)
	}
	fmt.Fprintf(w, "Reset successfully")
}
