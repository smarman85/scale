package pods

import (
  f "fmt"
  //"encoding/json"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  //"k8s.io/apimachinery/pkg/api/errors"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/rest"
  "os"
)

//func Pods(namespace, podName string){
func ListPods(namespace string) []map[string]interface{} {
  // creates the in-cluster config
  config, err := rest.InClusterConfig()
  if err != nil {
    panic(err.Error())
  }

  allPods := make([]map[string]interface{}, 0)

  clientset, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }

  //https://godoc.org/k8s.io/api/core/v1#Pod
  //https://godoc.org/k8s.io/api/core/v1#PodLogOptions
  ps, err := clientset.CoreV1().Pods(namespace).List(metav1.ListOptions{})
  //pod, err := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
  if err != nil {
    panic(err.Error())
  }


  for p, _ := range ps.Items {
        newPod := make(map[string]interface{}, 0)
	f.Println("**********************************************")
	f.Println(ps.Items[p].Name)
	newPod["name"] = ps.Items[p].Name
	newPod["namespace"] = ps.Items[p].Namespace
	newPod["service_account"] = ps.Items[p].Spec.ServiceAccountName
	newPod["restart_policy"] = ps.Items[p].Spec.RestartPolicy
	f.Println(ps.Items[p].Namespace)
	f.Println(ps.Items[p].Spec.ServiceAccountName)
	f.Println(ps.Items[p].Spec.RestartPolicy)

	cont := make(map[string]string, 0)

	containers := ps.Items[p].Spec.Containers
	for c, _ := range containers {
		cont["name"] = containers[c].Name
		cont["image"] = containers[c].Image
		f.Println(containers[c].Name)
		f.Println(containers[c].Image)
		//f.Println(containers[c].Resources.)
	}
	newPod["containers"] = cont

	allPods = append(allPods, newPod)
  }
  return allPods

}

func logErrorExitf(msg string, args ...interface{}) {
  f.Fprintf(os.Stderr, msg+"\n", args...)
  os.Exit(1)
}

func logErrorf(msg string, args ...interface{}) {
  f.Fprintf(os.Stderr, msg+"\n", args...)
}
