package jobs

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

//JobTemplate to hold job parameters
type JobTemplate struct {
	Name      string
	Imagename string
	Command   string
}

//RunJobs submitted jobs are executed on the cluster
func RunJobs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:" + r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./src/app/frontend/jobs/jobs.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		jobTemplate := &JobTemplate{
			Name:      r.Form["name"][0],
			Imagename: r.Form["imagename"][0],
		}
		for _, command := range r.Form["command"] {
			jobTemplate.Command += " " + command
		}
		t, err := template.ParseFiles("./templates/cpu-job-template.yaml")
		if err != nil {
			fmt.Printf("Error in templating: %s", err)
		}
		file, err := os.Create("jobfiles/" + jobTemplate.Name + ".yaml")
		if err != nil {
			fmt.Printf("Error in creating files: %s", err)
		}
		err = t.Execute(file, &jobTemplate)
		if err != nil {
			fmt.Printf("Error in executing template: %s", err)
		}
		b, err := json.Marshal(jobTemplate)
		if err != nil {
			fmt.Printf("Error in marshalling: %s", err)
		}

		fmt.Fprintf(w, string(b))
	}
}

//JobSubmitHandler handles /jobs request
func JobSubmitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:" + r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./src/app/frontend/jobs/jobs.html")
		t.Execute(w, nil)
	}
}
