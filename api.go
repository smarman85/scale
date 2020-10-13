package main

import (
	"os"
	f "fmt"
	//"log"
	"net/http"
	"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        //"k8s.io/apimachinery/pkg/api/errors"
        "k8s.io/client-go/kubernetes"
        "k8s.io/client-go/rest"
	"html/template"
	"github.com/gorilla/mux"
)

type Deployments struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Replicas int `json:"replicas"`
	Containers map[string]string `json:"containers"`
}

var tpl *template.Template
var NAMESPACE string = ""

func init(){
	tpl = template.Must(template.ParseGlob("tmpl/*.gohtml"))
}

func deployment(deploy string) {
	if os.Getenv("NAMESPACE") != "" {
		NAMESPACE = os.Getenv("NAMESPACE")
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		logErrorExitf("Error creating config: %v", err)
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logErrorExitf("Error creating config: %v", err)
		panic(err.Error())
	}

	dep, err := clientset.AppsV1().Deployments("gosite").Get(deploy, metav1.GetOptions{})
	if err != nil {
		logErrorf("Error getting deployment: %v", err)
	}
	f.Println(dep.Name)
	d := Deployments{
		Name: dep.Name,
		Namespace: dep.Namespace,
		Replicas: int(dep.Status.Replicas),
	}
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(d)
	//return d
}

func deployments() []Deployments {
	if os.Getenv("NAMESPACE") != "" {
		NAMESPACE = os.Getenv("NAMESPACE")
	}

	config, err := rest.InClusterConfig()
	if err != nil {
		logErrorExitf("Error creating config: %v", err)
		panic(err.Error())
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logErrorExitf("Error creating config: %v", err)
		panic(err.Error())
	}

	deps, err := clientset.AppsV1().Deployments(NAMESPACE).List(metav1.ListOptions{})
	if err != nil {
		logErrorf("Error getting deployments in namespace %s: %v", NAMESPACE, err)
	}

	dplymts := make([]Deployments, 0)

	for d, _ := range deps.Items {
		cont := make(map[string]string, 0)
		containers := deps.Items[d].Spec.Template.Spec.Containers
		for c, _ := range containers{
			cont["image"] = containers[c].Image
			cont["name"] = containers[c].Name
		}
		d := Deployments{
			Name: deps.Items[d].Name,
			Namespace: deps.Items[d].Namespace,
			Replicas: int(deps.Items[d].Status.Replicas),
			Containers: cont,
		}
		dplymts = append(dplymts, d)
	}
	return dplymts
}

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", home)
	router.HandleFunc("/deployment/{deployment}", dep)
	http.ListenAndServe(":8080", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	deps := deployments()
	err := json.NewEncoder(w).Encode(deps)
	if err != nil {
		logErrorExitf("Error encoding data %v", err)
	}
}

func dep(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	dep := vars["deployment"]
	dplyment := deployment(dep)
	//f.Println(dep)
}

func logErrorExitf(msg string, args ...interface{}) {
	f.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func logErrorf(msg string, args ...interface{}) {
	f.Fprintf(os.Stderr, msg+"\n", args...)
}
