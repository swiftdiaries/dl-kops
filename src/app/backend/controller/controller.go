package controller

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swiftdiaries/dl-kops/src/app/backend"
)

//SetupController is used to setup kubernetes on the controller node
func SetupController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		hostname := r.Form["hostname"][0]
		hostip := r.Form["hostip"][0]
		keyfile := r.Form["keyfile"][0]
		var output []string
		//args = []string{"./scripts/trial.sh", hostip}
		command := "./controllerkubeup.sh " + hostip
		output = backend.ExecuteSSHCommand(hostname, hostip, keyfile, command)
		/*
			shcmd := "sh"
			var args []string
			args = []string{"./scripts/controllerkubeup.sh", hostip}
			cmd := exec.Command(shcmd, args...)
			cmd.Stdin = strings.NewReader("")
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				log.Fatalf("exec Error: %s", err)
			}
			output = append(output, out.String())
		*/
		fmt.Fprintf(w, "%s", output)
	}
}

//InstallController is used to display the :port/installcontroller call
func InstallController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		//fmt.Printf("name:%s,\nip:%s\nkey:%s\n", r.Form["hostname"], r.Form["hostip"], r.Form["keyfile"])
		hostname := r.Form["hostname"][0]
		hostip := r.Form["hostip"][0]
		keyfile := r.Form["keyfile"][0]
		shcmd := "sh"
		var args []string
		var output []string
		//args = []string{"./scripts/trial.sh", hostname, keyfile, hostip}
		args = []string{"./scripts/setup_controller.sh", hostname, keyfile, hostip}
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
		t, _ := template.ParseFiles("./result/result.html")
		t.Execute(w, nil)
	}
}

//GetToken is used to display the :port/gettoken call
func GetToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		//fmt.Printf("name:%s,\nip:%s\nkey:%s\n", r.Form["hostname"], r.Form["hostip"], r.Form["keyfile"])
		hostname := r.Form["hostname"][0]
		hostip := r.Form["hostip"][0]
		keyfile := r.Form["keyfile"][0]
		var output []string
		output = backend.ExecuteSSHCommand(hostname, hostip, keyfile, "kubeadm token generate")
		/*
			shcmd := "sh"
			var args []string
			//args = []string{"./scripts/trial.sh", hostname, keyfile, hostip}
			args = []string{"kubeadm generate token", hostname, keyfile, hostip}
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
		*/

		fmt.Fprintf(w, "%s", output)
	} else {
		t, _ := template.ParseFiles("./result/result.html")
		t.Execute(w, nil)
	}
}
