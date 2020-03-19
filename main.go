package main

import (
	"fmt"
  "github.com/kmdrn7/desu/cmd/pod"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {

  kubeconfig := SetupKubeConfig()
  ns := SetupNamespace()
  listOptions := SetupListOptions()

  parseFlag()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

  // list all pods using pods package
  pods.List()

  // create the api object
  api := clientset.CoreV1()

  pods, err := api.Pods(ns).List(listOptions)
  if err != nil {
    panic(err.Error())
  }

  fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))

}
