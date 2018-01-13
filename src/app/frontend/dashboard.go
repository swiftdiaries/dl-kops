package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/swiftdiaries/dl-kops/src/app/backend/controller"
	"github.com/swiftdiaries/dl-kops/src/app/backend/worker"
	"github.com/swiftdiaries/dl-kops/src/app/frontend/jobs"
)

var (
	port = os.Getenv("PORT_1")
)

func main() {

	fileServerIndex := http.FileServer(http.Dir("./src/app/frontend/index/"))
	http.Handle("/", fileServerIndex)
	http.HandleFunc("/installcontroller", controller.InstallController)
	http.HandleFunc("/setupcontroller", controller.SetupController)
	http.HandleFunc("/installworker", worker.InstallWorker)
	http.HandleFunc("/setupworker", worker.SetupWorker)
	http.HandleFunc("/gettoken", controller.GetToken)
	http.HandleFunc("/jobs", jobs.JobSubmitHandler)
	http.HandleFunc("/submit", jobs.RunJobs)
	go open("http://localhost:" + port + "/")
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
