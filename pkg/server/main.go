package server

import (
	"os"
	f "fmt"
	"log"
	"net/http"
	"encoding/json"
	"html/template"
	"scale/pkg/deployments"
	"scale/pkg/pods"
	"github.com/gorilla/mux"
)

var tpl *template.Template
var NAMESPACE string = ""

func init(){
	tpl = template.Must(template.ParseGlob("tmpl/*.gohtml"))
}


func Run() {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", home)
	//router.HandleFunc("/pods/{podName}", podsHandler)
	router.HandleFunc("/api/v1/deployments", apiHome)
	router.HandleFunc("/api/v1/pods", apiPods)
	http.ListenAndServe(":8080", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	deps := deployments.Deployments(NAMESPACE)
	err := tpl.ExecuteTemplate(w, "index.gohtml", deps)
	if err != nil {
		log.Println("template didn't execute: ", err)
	}
}

func apiHome(w http.ResponseWriter, r *http.Request) {
	deps := deployments.Deployments(NAMESPACE)
	err := json.NewEncoder(w).Encode(deps)
	//out, err := json.Marshal(deps)
	if err != nil {
		logErrorExitf("Error encoding data %v", err)
	}
	//f.Println(string(out))
}

func apiPods(w http.ResponseWriter, r *http.Request) {
	pds := pods.ListPods(NAMESPACE)
	err := json.NewEncoder(w).Encode(pds)
	if err != nil {
		logErrorExitf("Error encoding data %v", err)
	}
}

func logErrorExitf(msg string, args ...interface{}) {
	f.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func logErrorf(msg string, args ...interface{}) {
	f.Fprintf(os.Stderr, msg+"\n", args...)
}
