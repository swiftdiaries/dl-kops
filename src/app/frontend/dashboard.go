package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"

	"github.com/swiftdiaries/dl-kops/src/app/backend"
)

var (
	port = os.Getenv("PORT_1")
)

func main() {

	fileServerIndex := http.FileServer(http.Dir("./src/app/frontend/index/"))
	http.Handle("/", fileServerIndex)
	http.HandleFunc("/result", Output)
	fileServerResult := http.FileServer(http.Dir("./result/"))
	http.Handle("/display", fileServerResult)
	fmt.Print("Serving on http://localhost:" + port + "/\n")
	go open("http://localhost:" + port + "/")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//Output is used to display the :port/result call
func Output(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()

		fmt.Printf("name:%s,\nip:%s\nkey:%s\n", r.Form["hostname"][0], r.Form["hostip"][0], r.Form["keyfile"][0])
		hostname := r.Form["hostname"][0]
		hostip := r.Form["hostip"][0]
		keyfile := r.Form["keyfile"][0]
		filename := "./src/app/backend/setupkubernetes.sh"
		outputlogs := backend.ExecuteThroughSSH(hostname, hostip, keyfile, filename)
		//outlogs := backend.PrintLines(filename)
		//fmt.Fprintf(w, "%s", outlogs)
		fmt.Fprintf(w, "%s", outputlogs)
	} else {
		t, _ := template.ParseFiles("./result/result.html")
		t.Execute(w, nil)
	}
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
