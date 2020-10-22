package pods

import (
  //"encoding/json"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  //"k8s.io/apimachinery/pkg/api/errors"
  //"k8s.io/api/core/v1"
  //"k8s.io/client-go/rest"

  "scale/pkg/globals"
)

/*https://stackoverflow.com/questions/53852530/how-to-get-logs-from-kubernetes-using-golang?rq=1 
func podLogs(namespace, container string, clientset *kubernetes.Clientset) {
	lgs, err := clientset.CoreV1().Pods(namespace).GetLogs(container, *v1.PodLogOptions{})
	if err != nil {
		logErrorf("Probelm getting logs", err)
	}
	f.Println(lgs)
}*/

//func Pods(namespace, podName string){
//func ListPods(clientset *kubernetes.Clientset) []map[string]interface{} {
func ListPods() []map[string]interface{} {
  allPods := make([]map[string]interface{}, 0)

  ps, err := globals.Clientset.CoreV1().Pods(globals.NAMESPACE).List(metav1.ListOptions{})
  //pod, err := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
  if err != nil {
    panic(err.Error())
  }


  for p, _ := range ps.Items {
        newPod := make(map[string]interface{}, 0)
	/*f.Println("**********************************************")
	f.Println(ps.Items[p].Name)
	f.Println(ps.Items[p].Namespace)
	f.Println(ps.Items[p].Spec.ServiceAccountName)
	f.Println(ps.Items[p].Spec.RestartPolicy)*/
	newPod["name"] = ps.Items[p].Name
	newPod["namespace"] = ps.Items[p].Namespace
	newPod["service_account"] = ps.Items[p].Spec.ServiceAccountName
	newPod["restart_policy"] = ps.Items[p].Spec.RestartPolicy

	cont := make(map[string]string, 0)

	containers := ps.Items[p].Spec.Containers
	for c, _ := range containers {
		cont["name"] = containers[c].Name
		cont["image"] = containers[c].Image
		/*f.Println(containers[c].Name)
		f.Println(containers[c].Image)*/
		//f.Println(containers[c].Resources.)
	}
	newPod["containers"] = cont

	allPods = append(allPods, newPod)
  }
  return allPods

}
