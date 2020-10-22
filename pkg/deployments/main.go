package deployments

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        //"k8s.io/apimachinery/pkg/api/errors"
        //"k8s.io/client-go/kubernetes"
	"scale/pkg/globals"
	"scale/pkg/logging"
)

/*type Deployments struct {
	Name string `json:"name"`
	Namespace string `json:"namespace"`
	Replicas int `json:"replicas"`
	Containers map[string]string `json:"containers"`
}*/

//func Deployments(clientset *kubernetes.Clientset) []map[string]interface{} {
func Deployments() []map[string]interface{} {

	dplymts := make([]map[string]interface{}, 0)

	deps, err := globals.Clientset.AppsV1().Deployments(globals.NAMESPACE).List(metav1.ListOptions{})
	if err != nil {
		logging.LogErrorf("Error getting deployments in namespace %s: %v", globals.NAMESPACE, err)
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
