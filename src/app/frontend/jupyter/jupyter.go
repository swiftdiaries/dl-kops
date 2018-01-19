package jupyter

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swiftdiaries/dl-kops/src/app/backend/utils"
)

var (
	jupyterhtmlpath = utils.HomeDir + "/src/app/frontend/jupyter/jupyter.html"
)

//Resources describes notebook resources
type Resources struct {
	NvidiaGPU string
	CPU       string
}

//LaunchJupyter handles /jupyter calls and launches a jupyter notebook
func LaunchJupyter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:" + r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles(jupyterhtmlpath)
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		resourceTemplate := &Resources{
			NvidiaGPU: r.FormValue("gpu"),
			CPU:       r.FormValue("cpu"),
		}
		jupyteryamlpath := utils.HomeDir + "templates/jupyter.yaml"
		t, err := template.ParseFiles(jupyteryamlpath)
		if err != nil {
			fmt.Printf("Error in parsing yaml:%s", err)
		}
		filename := "jobfiles/jupyter-depl-svc.yaml"
		file, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Error in creating files: %s", err)
		}
		err = t.Execute(file, &resourceTemplate)
		if err != nil {
			fmt.Printf("Error in creating file: %s", err)
		}
		jobfilepath := utils.HomeDir + filename
		output := utils.KubectlExecuteYaml(jobfilepath)
		podslog := utils.KubectlExecuteCommand([]string{"get", "pods", "-o go-template --template '{{range .items}}{{.metadata.name}}{{'\n'}}{{end}}'"})
		output = append(output, strings.Join(podslog, "\n"))
		output = append(output, "Wait for the pod to start running and then execute:\n")
		output = append(output, "$kubectl port-forward ${JUPYTER_POD_NAME} 8888:8888")
		fmt.Fprintf(w, strings.Join(output, " "))
	}
}

//GetPodLogs fetches a pods logs
func GetPodLogs(podName string) []string {
	return utils.KubectlExecuteCommand([]string{"logs", podName})
}

//PortForward enable port-forward on a given pod, local port and containerport
func PortForward(podName string, localport string, containerport string) []string {
	return utils.KubectlExecuteCommand([]string{"port-forward", podName, localport + ":" + containerport})
}
