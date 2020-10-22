package globals

import (
	"os"
        "html/template"
        "k8s.io/client-go/kubernetes"
        "k8s.io/client-go/rest"
	"scale/pkg/logging"
)

var Tpl *template.Template
var NAMESPACE string = ""
var Clientset *kubernetes.Clientset

//func initGlobals() (*template.Template, string, *kubernetes.Clientset) {
func initGlobals() {

        if os.Getenv("NAMESPACE") != "" {
                NAMESPACE = os.Getenv("NAMESPACE")
        }

        Tpl = template.Must(template.ParseGlob("tmpl/*.gohtml"))
        config, err := rest.InClusterConfig()
        if err != nil {
          logging.LogErrorExitf("Error creating config: %v", err)
        }

        Clientset, err = kubernetes.NewForConfig(config)
        if err != nil {
          logging.LogErrorExitf("Error creating config: %v", err)
        }


}

func init() {
	initGlobals()
}
