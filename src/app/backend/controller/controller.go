package controller

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

	"github.com/alecthomas/template"
	"github.com/spf13/viper"
	"github.com/swiftdiaries/dl-kops/src/app/backend"
	"github.com/swiftdiaries/dl-kops/src/app/backend/utils"
)

//RegisterController and write to file again.
func RegisterController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		hostname := r.FormValue("hostname")
		hostip := r.FormValue("hostip")
		keyfile := r.FormValue("keyfile")
		keyfile = strings.Replace(keyfile, "~", os.Getenv("HOME"), 1)
		viper.SetConfigName("config")
		viper.AddConfigPath(utils.HomeDir)
		err := viper.ReadInConfig()
		if err != nil {
			fmt.Printf("Error reading in config")
		}
		var cluster utils.ClusterConfig
		err = viper.Unmarshal(&cluster)
		if err != nil {
			fmt.Printf("Error in unmarshalling: %s", err)
		}
		cluster.Controller.Hostname = hostname
		cluster.Controller.Hostip = hostip
		cluster.Controller.Keyfilepath = keyfile
		b, err := json.Marshal(cluster)
		if err != nil {
			fmt.Printf("Error in marshalling: %s", err)
		}
		err = ioutil.WriteFile("config.yaml", b, 0644)
		if err != nil {
			fmt.Printf("Error in writing:%s", err)
		}
	}
}

//SetupController is used to setup kubernetes on the controller node
func SetupController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		hostname, hostip, keyfile := utils.GetCreds("controller")
		var output []string
		command := utils.HomeDir + "/controllerkubeup.sh " + hostip
		output = backend.ExecuteSSHCommand(hostname, hostip, keyfile, command)
		fmt.Fprintf(w, "%s", output)
	}
}

//InstallController is used to display the :port/installcontroller call
func InstallController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		hostname, hostip, keyfile := utils.GetCreds("controller")
		shcmd := "sh"
		var args []string
		var output []string
		scriptfilepath := utils.HomeDir + "/scripts/setup_controller.sh "
		args = []string{scriptfilepath, hostname, keyfile, hostip}
		fmt.Printf("Args: %s", args)
		cmd := exec.Command(shcmd, args...)
		cmd.Stdin = strings.NewReader("")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatalf("exec Error: %s", err)
		}
		output = append(output, out.String())
		fmt.Fprintf(w, "%s", output)
	} else {
		htmlfilepath := utils.HomeDir + "/result/result.html"
		t, _ := template.ParseFiles(htmlfilepath)
		t.Execute(w, nil)
	}
}

//GetToken is used to display the :port/gettoken call
func GetToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		hostname, hostip, keyfile := utils.GetCreds("controller")
		var output []string
		output = backend.ExecuteSSHCommand(hostname, hostip, keyfile, "kubeadm token generate")
		fmt.Fprintf(w, "%s", output)
	} else {
		htmlfilepath := utils.HomeDir + "/result/result.html"
		t, _ := template.ParseFiles(htmlfilepath)
		t.Execute(w, nil)
	}
}
