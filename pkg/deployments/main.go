package deployments

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"scale/pkg/globals"
	"scale/pkg/logging"
	"encoding/json"
	//f "fmt"
)

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

func GetDeployment(namespace, deployName string) map[string]interface{} {
//func GetDeployment(namespace, deployName string) {
	dep, err := globals.Clientset.AppsV1().Deployments(namespace).Get(deployName, metav1.GetOptions{})
	if err != nil {
		logging.LogErrorf("Error getting deployment %s in namespace %s: %v", deployName, namespace, err)
	}
	/*f.Println(dep.Name)
	f.Println(dep.Namespace)
	f.Println(dep.CreationTimestamp)
	for k, v := range dep.Labels {
		f.Printf("%s:%s\n", k, v)
	}
	for k, v := range dep.Annotations {
		f.Printf("%s:%s\n", k, v)
	}*/
	config := dep.Annotations["kubectl.kubernetes.io/last-applied-configuration"]
	configMap := make(map[string]interface{})

	err = json.Unmarshal([]byte(config), &configMap)
	if err != nil {
		logging.LogErrorf("Error converting config to map namespace: %s: app: %s error: %v", namespace, deployName, err)
	}

	return configMap

	/*f.Println(configMap["apiVersion"])
	f.Println(configMap["kind"])
	f.Println(configMap["metadata"])
	f.Println(configMap["spec"])*/


}
