package main

import (
        //"context"
        "time"
        f "fmt"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        "k8s.io/apimachinery/pkg/api/errors"
        "k8s.io/client-go/kubernetes"
        "k8s.io/client-go/rest"
)

func main() {
        f.Println("Hello")
        // creates the in-cluster config
        config, err := rest.InClusterConfig()
        if err != nil {
                panic(err.Error())
        }


        clientset, err := kubernetes.NewForConfig(config)
        if err != nil {
                panic(err.Error())
        }

        for {
                // get deployments
                deps, err := clientset.AppsV1().Deployments("").List(metav1.ListOptions{})
                if err != nil {
                        panic(err.Error())
                }
                f.Printf("There are %d deployments running in the cluster\n", len(deps.Items))
                for k, _ := range deps.Items {
                        f.Println("NEW")
                        //f.Printf("%T\n", k)
                        f.Printf("%T\n", deps.Items[k])
                        //f.Println("%T\n", deps.Items[k].Spec.Template.Spec.Containers)
                        f.Println(deps.Items[k].Name)
                }

                time.Sleep(10 * time.Second)
        }


        for {
                // get pods in all the namespaces by omitting namespace
                // Or specify namespace to get pods in particular namespace
                //pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
                pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
                if err != nil {
                        panic(err.Error())
                }
                f.Printf("There are %d pods in the cluster\n", len(pods.Items))
                // Examples for error handling:
                // - Use helper functions e.g. errors.IsNotFound()
                // - And/or cast to StatusError and use its properties like e.g. ErrStatus.Message
                _, err = clientset.CoreV1().Pods("default").Get("example-xxxxx", metav1.GetOptions{})
                if errors.IsNotFound(err) {
                        f.Printf("Pod example-xxxxx not found in default namespace\n")
                } else if statusError, isStatus := err.(*errors.StatusError); isStatus {
                        f.Printf("Error getting pod %v\n", statusError.ErrStatus.Message)
                } else if err != nil {
                        panic(err.Error())
                } else {
                        f.Printf("Found example-xxxxx pod in default namespace\n")
                }

                time.Sleep(10 * time.Second)
        }
}

