package deployments

import (
	"os"
	f "fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  //"k8s.io/apimachinery/pkg/api/errors"
  "k8s.io/client-go/kubernetes"
)

/*type Deployments struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Replicas int `json:"replicas"`
	Containers map[string]string `json:"containers"`
}*/

func Deployments(namespace string, clientset *kubernetes.Clientset) []map[string]interface{} {

	dplymts := make([]map[string]interface{}, 0)

	deps, err := clientset.AppsV1().Deployments(namespace).List(metav1.ListOptions{})
	if err != nil {
		logErrorf("Error getting deployments in namespace %s: %v", namespace, err)
	}

	for d, _ := range deps.Items {
		dep := make(map[string]interface{}, 0)
		cont := make(map[string]string, 0)
		containers := deps.Items[d].Spec.Template.Spec.Containers
		for c, _ := range containers{
			cont["image"] = containers[c].Image
			cont["name"] = containers[c].Name
		}
		dep["name"] = deps.Items[d].Name
		dep["namespace"] = deps.Items[d].Namespace
		dep["replicas"] = int(deps.Items[d].Status.Replicas)
		dep["containers"] = cont

		dplymts = append(dplymts, dep)
	}
	return dplymts
}


func logErrorExitf(msg string, args ...interface{}) {
	f.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func logErrorf(msg string, args ...interface{}) {
	f.Fprintf(os.Stderr, msg+"\n", args...)
}
