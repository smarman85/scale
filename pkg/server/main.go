package server

import (
	"net/http"
	"encoding/json"
	"scale/pkg/deployments"
	"scale/pkg/pods"
	"scale/pkg/globals"
	"scale/pkg/logging"
	"github.com/gorilla/mux"
)

func Run() {
	router := mux.NewRouter()

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	router.HandleFunc("/", home)
	//router.HandleFunc("/pods/{podName}", podsHandler)
	router.HandleFunc("/api/v1/deployments", apiHome)
	router.HandleFunc("/api/v1/deployment/{namespace}/{deployName}", apiDeployment)
	router.HandleFunc("/api/v1/pods", apiPods)
	http.ListenAndServe(":8080", router)
}

func home(w http.ResponseWriter, r *http.Request) {
	deps := deployments.Deployments()
	err := globals.Tpl.ExecuteTemplate(w, "index.gohtml", deps)
	if err != nil {
		logging.LogErrorf("template didn't execute: %v", err)
	}
}

func apiHome(w http.ResponseWriter, r *http.Request) {
	deps := deployments.Deployments()
	err := json.NewEncoder(w).Encode(deps)
	//out, err := json.Marshal(deps)
	if err != nil {
		logging.LogErrorf("Error encoding data %v", err)
	}
	//f.Println(string(out))
}

func apiDeployment(w http.ResponseWriter, r * http.Request) {
	vars := mux.Vars(r)
	dep := deployments.GetDeployment(vars["namespace"], vars["deployName"])
	err := json.NewEncoder(w).Encode(dep)
	if err != nil {
		logging.LogErrorf("Error encoding data from deployment %v", err)
	}
}

func apiPods(w http.ResponseWriter, r *http.Request) {
	pds := pods.ListPods()
	err := json.NewEncoder(w).Encode(pds)
	if err != nil {
		logging.LogErrorf("Error encoding data %v", err)
	}
}
