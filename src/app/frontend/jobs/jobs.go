package jobs

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"

	"github.com/swiftdiaries/dl-kops/src/app/backend/utils"
)

//JobTemplate to hold job parameters
type JobTemplate struct {
	Name      string
	Imagename string
	Command   string
}

var (
	jobshtmlpath = utils.HomeDir + "/src/app/frontend/jobs/jobs.html"
)

//RunJobs submitted jobs are executed on the cluster
func RunJobs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:" + r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles(jobshtmlpath)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		device := r.FormValue("device")
		jobTemplate := &JobTemplate{
			Name:      r.Form["name"][0],
			Imagename: r.Form["imagename"][0],
		}
		for _, command := range r.Form["command"] {
			jobTemplate.Command += " " + command
		}
		cpuyamlfilepath := utils.HomeDir + "/templates/cpu-job-template.yaml"
		gpuyamlfilepath := utils.HomeDir + "/templates/gpu-job-template.yaml"
		t, err := template.ParseFiles(cpuyamlfilepath)
		if device == "gpu" {
			t, err = template.ParseFiles(gpuyamlfilepath)
			if err != nil {
				fmt.Printf("Error in templating: %s", err)
			}
		}
		filename := "jobfiles/" + jobTemplate.Name + ".yaml"
		file, err := os.Create(filename)
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
		jobfilepath := utils.HomeDir + filename
		output := utils.KubectlExecuteYaml(jobfilepath)
		output = append(output, string(b))
		fmt.Fprintf(w, strings.Join(output, " "))
	}
}

//JobSubmitHandler handles /jobs request
func JobSubmitHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:" + r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles(jobshtmlpath)
		t.Execute(w, nil)
	}
}
