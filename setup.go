package main

import (
  "fmt"
  "flag"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
  "os"
  "path/filepath"
)

func SetupKubeConfig() string {

  fmt.Println("./Desu setup kubeconfig")

  var kubeconfig string
  flag.StringVar(&kubeconfig, "kubeconfig", filepath.Join(os.Getenv("HOME"), ".kube", "config"), "/home/kmdr7/.kube/config")
  return kubeconfig
}

func SetupListOptions() v1.ListOptions {

  fmt.Println("./Desu setup list options")

  var label, field string
  flag.StringVar(&label, "label", "", "Resource Label")
  flag.StringVar(&field, "field", "", "Resource Field")

  listOptions := v1.ListOptions {
    LabelSelector: label,
    FieldSelector: field,
  }
  return listOptions
}

func SetupNamespace() string {

  fmt.Println("./Desu setup namespace")

  var ns string
  flag.StringVar(&ns, "namespace", "default", "Resource Namespace")
  return ns
}

func parseFlag() {
  flag.Parse()
}

