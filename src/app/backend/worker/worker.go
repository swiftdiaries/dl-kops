package worker

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

//SetupWorker is used to setup kubernetes on the controller node
func SetupWorker(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		hostname := r.Form["hostname"][0]
		hostip := r.Form["hostip"][0]
		keyfile := r.Form["keyfile"][0]
		masterip := r.Form["masterip"][0]
		jointoken := r.Form["jointoken"][0]
		var output []string
		command := "./workerup.sh " + jointoken + " " + masterip
		fmt.Println("Join worker: \n" + hostname + "\n" + hostip + "\n" + keyfile + masterip + " " + jointoken)
		output = backend.ExecuteSSHCommand(hostname, hostip, keyfile, command)
		/*
			shcmd := "sh"
			var args []string
			//args = []string{"./scripts/trial.sh", hostip, jointoken}
			args = []string{"./scripts/workerup.sh", jointoken, hostip}
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

//InstallWorker is used to display the :port/installcontroller call
func InstallWorker(w http.ResponseWriter, r *http.Request) {
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
		args = []string{"./scripts/setup_worker.sh", hostname, keyfile, hostip}
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
