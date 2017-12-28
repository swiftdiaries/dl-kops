package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
)

var (
	port = os.Getenv("PORT_1")
)

func main() {

	fileServerIndex := http.FileServer(http.Dir("./src/app/frontend/index/"))
	http.Handle("/", fileServerIndex)
	http.HandleFunc("/result", output)
	fileServerResult := http.FileServer(http.Dir("./result/"))
	http.Handle("/display", fileServerResult)
	fmt.Print("Serving on http://localhost:" + port + "/")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func output(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "POST" {
		r.ParseForm()
		fmt.Println(r.Form["username"], r.Form["phonenumber"])
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
