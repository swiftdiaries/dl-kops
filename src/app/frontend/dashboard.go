package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

var (
	port = os.Getenv("PORT_1")
)

func main() {

	fileServerIndex := http.FileServer(http.Dir("./src/app/frontend/index/"))
	http.Handle("/", fileServerIndex)
	http.HandleFunc("/result", Output)
	http.HandleFunc("/setup", SetupController)
	//fmt.Print("Serving on http://localhost:" + port + "/\n")
	go open("http://localhost:" + port + "/")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

//SetupController is used to setup kubernetes on the controller node
func SetupController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		hostip := r.Form["hostip"][0]
		shcmd := "sh"
		var args []string
		var output []string
		args = []string{"./src/app/backend/trial.sh", hostip}
		//args = []string{"./src/app/backend/controllerkubeup.sh", hostip}
		fmt.Printf("Args: %s", args)
		//args = append(args, tempargs)
		cmd := exec.Command(shcmd, args...)
		cmd.Stdin = strings.NewReader("")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatalf("exec Error: %s", err)
		}
		//fmt.Printf("%s", out.String())
		output = append(output, out.String())

		fmt.Fprintf(w, "%s", output)
	}
}

//Output is used to display the :port/result call
func Output(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()

		//fmt.Printf("name:%s,\nip:%s\nkey:%s\n", r.Form["hostname"][0], r.Form["hostip"][0], r.Form["keyfile"][0])
		hostname := r.Form["hostname"][0]
		hostip := r.Form["hostip"][0]
		keyfile := r.Form["keyfile"][0]
		//filename := "./src/app/backend/trial.sh"
		shcmd := "sh"
		var args []string
		var output []string
		args = []string{"./src/app/backend/trial.sh", hostname, keyfile, hostip}
		//args = []string{"./src/app/backend/setup_cluster.sh", hostname, keyfile, hostip}
		fmt.Printf("Args: %s", args)
		//args = append(args, tempargs)
		cmd := exec.Command(shcmd, args...)
		cmd.Stdin = strings.NewReader("")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		if err != nil {
			log.Fatalf("exec Error: %s", err)
		}
		//fmt.Printf("%s", out.String())
		output = append(output, out.String())

		fmt.Fprintf(w, "%s", output)
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
