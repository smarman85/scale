package main

import (
        //"context"
        //"time"
        //"strconv"
        f "fmt"
        "encoding/json"
        metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
        //"k8s.io/apimachinery/pkg/api/errors"
        "k8s.io/client-go/kubernetes"
        "k8s.io/client-go/rest"
        "os"
)

type Deployment struct {
        Namespace string `json:"namespace"`
        Name string `json:"name"`
        Replicas int `json:"replicas"`
        Containers []containers `json:"containers"`
        Nodeselector string `json:"nodeselector"`
        Nodename string `json:"nodename"`
}

type containers struct {
        Image string `json:"image"`
        Name string `json:"name"`
}

func main() {

        NAMESPACE := ""

        if os.Getenv("NAMESPACE") != "" {
                NAMESPACE = os.Getenv("NAMESPACE")
        }
        // creates the in-cluster config
        config, err := rest.InClusterConfig()
        if err != nil {
                panic(err.Error())
        }


        clientset, err := kubernetes.NewForConfig(config)
        if err != nil {
                panic(err.Error())
        }

        // get deployments
        deps, err := clientset.AppsV1().Deployments(NAMESPACE).List(metav1.ListOptions{})
        if err != nil {
                panic(err.Error())
        }
        f.Printf("There are %d deployments running in the cluster\n", len(deps.Items))
        //https://godoc.org/k8s.io/api/apps/v1#DeploymentList
        for k, _ := range deps.Items {
                dep := &Deployment{
                        Namespace: deps.Items[k].Namespace,
                        Name: deps.Items[k].Name,
                        Replicas: int(deps.Items[k].Status.Replicas),
                        //Containers:
                        //Nodeselector: deps.Items[k].Spec.Template.Spec.NodeSelector,
                        Nodename: deps.Items[k].Spec.Template.Spec.NodeName,
                }

                output, err := json.Marshal(dep)
                if err != nil {
                        panic(err)
                }
                f.Println(string(output))
                //f.Println("%T\n", deps.Items[k].Spec.Template.Spec.Containers)
                f.Printf("Name: %s\n", deps.Items[k].Name)
                f.Printf("Namespace: %s\n", deps.Items[k].Namespace)
                f.Printf("Created: %v\n", deps.Items[k].CreationTimestamp)
                //reaplicas := deps.Items[k].Spec.Replicas
                f.Printf("Replicas: %v\n", deps.Items[k].Status.Replicas)
                f.Println("POD INFO")
                f.Println("Containers")
                //https://godoc.org/k8s.io/api/core/v1#PodTemplateSpec
                containers := deps.Items[k].Spec.Template.Spec.Containers
                for container, _ := range containers {
                        //https://godoc.org/k8s.io/api/core/v1#Container
                        f.Printf("ContainerName: %s\n", containers[container].Name)
                        f.Printf("ContainerImage: %s\n", containers[container].Image)
                        //f.Printf("ContainerENV: %s\n", containers[container].Env)
                }
                f.Println("Nodeselector")
                f.Println(deps.Items[k].Spec.Template.Spec.NodeSelector)
                f.Println("NodeName")
                f.Println(deps.Items[k].Spec.Template.Spec.NodeName)
                f.Println("*******")
                //f.Println(deps.Items[k])
                //f.Printf("%T\n", deps.Items[k])
                //f.Println(deps.Items[k])
        }


        /*for {
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
        }*/
}

