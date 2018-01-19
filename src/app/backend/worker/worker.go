package worker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/alecthomas/template"
	"github.com/spf13/viper"
	"github.com/swiftdiaries/dl-kops/src/app/backend"
	"github.com/swiftdiaries/dl-kops/src/app/backend/utils"
)

//RegisterWorker Registers worker into config
func RegisterWorker(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		hostname := r.Form["hostname"][0]
		hostip := r.Form["hostip"][0]
		keyfile := r.Form["keyfile"][0]
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
	}
}

//SetupWorker is used to setup kubernetes on the controller node
func SetupWorker(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		hostname, hostip, keyfile := utils.GetCreds("worker")
		masterip := r.Form["masterip"][0]
		jointoken := r.Form["jointoken"][0]
		certs := r.Form["certs"][0]
		var output []string
		command := "./workerup.sh " + jointoken + " " + masterip + " " + certs
		fmt.Println("Join worker: \n" + hostname + "\n" + hostip + "\n" + keyfile + masterip + "\n" + jointoken + "\n" + certs)
		output = backend.ExecuteSSHCommand(hostname, hostip, keyfile, command)
		fmt.Fprintf(w, "%s", output)
	}
}

//InstallWorker is used to display the :port/installcontroller call
func InstallWorker(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		hostname, hostip, keyfile := utils.GetCreds("worker")
		filepath := utils.HomeDir + "/scripts/setup_worker.sh"
		var arguments []string
		var output []string
		arguments = []string{hostname, keyfile, hostip}
		output = utils.ExecuteScriptFile(filepath, arguments)
		fmt.Fprintf(w, "%s", output)
	} else {
		resultfilepath := utils.HomeDir + "/result/result.html"
		t, _ := template.ParseFiles(resultfilepath)
		t.Execute(w, nil)
	}
}
