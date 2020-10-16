package pods

import (
  f "fmt"
  "encoding/json"
  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  //"k8s.io/apimachinery/pkg/api/errors"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/rest"
  "os"
)

func Pods(namespace, podName string){
  // creates the in-cluster config
  config, err := rest.InClusterConfig()
  if err != nil {
    panic(err.Error())
  }


  clientset, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }

  //https://godoc.org/k8s.io/api/core/v1#Pod
  //https://godoc.org/k8s.io/api/core/v1#PodLogOptions
  //pods, err := clientset.CoreV1()Pods(namespace).List(metav1.ListOptions{})
  pod, err := clientset.CoreV1().Pods(namespace).Get(podName, metav1.GetOptions{})
  if err != nil {
    panic(err.Error())
  }

  f.Prinln(pod)

}

func logErrorExitf(msg string, args ...interface{}) {
  f.Fprintf(os.Stderr, msg+"\n", args...)
  os.Exit(1)
}

func logErrorf(msg string, args ...interface{}) {
  f.Fprintf(os.Stderr, msg+"\n", args...)
}
