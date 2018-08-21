package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/swiftdiaries/dl-kops/src/app/backend/controller"
	"github.com/swiftdiaries/dl-kops/src/app/backend/utils"
	"github.com/swiftdiaries/dl-kops/src/app/backend/worker"
	"github.com/swiftdiaries/dl-kops/src/app/frontend/jobs"
	"github.com/swiftdiaries/dl-kops/src/app/frontend/jupyter"
)

var (
	port = os.Getenv("PORT_1")
)

func main() {
	fileServerIndex := http.FileServer(http.Dir("index/"))
	http.Handle("/", fileServerIndex)
	http.HandleFunc("/registercontroller", controller.RegisterController)
	http.HandleFunc("/installcontroller", controller.InstallController)
	http.HandleFunc("/setupcontroller", controller.SetupController)
	http.HandleFunc("/registerworker", worker.RegisterWorker)
	http.HandleFunc("/installworker", worker.InstallWorker)
	http.HandleFunc("/setupworker", worker.SetupWorker)
	http.HandleFunc("/gettoken", controller.GetToken)
	http.HandleFunc("/jupyter", jupyter.LaunchJupyter)
	http.HandleFunc("/jobs", jobs.JobSubmitHandler)
	http.HandleFunc("/submit", jobs.RunJobs)
	http.HandleFunc("/resetconfig", utils.ResetConfig)
	//go open("http://localhost:" + port + "/")
	log.Println("Listening on", ":"+port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func open(url string) error {

	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
		//args = []string{"-a", "'Google Chrome'"}
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	fmt.Println(cmd, args)
	return exec.Command(cmd, args...).Start()

}
