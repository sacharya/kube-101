/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
  "flag"
  "fmt"
  "time"

  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/kubernetes"
  "k8s.io/client-go/tools/clientcmd"
  // Uncomment the following line to load the gcp plugin (only required to authenticate against GKE clusters).
  // _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
  "k8s.io/client-go/pkg/api/v1"
  "k8s.io/client-go/tools/cache"
  "k8s.io/apimachinery/pkg/fields"
  //"k8s.io/client-go/rest"
)

func watchPods(clientset *kubernetes.Clientset) {
  watchlist := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "pods", v1.NamespaceAll, fields.Everything())
  _, controller := cache.NewInformer(
    watchlist,
    &v1.Pod{},
    time.Second * 0,
    cache.ResourceEventHandlerFuncs{
      AddFunc: func(obj interface{}) {
        fmt.Printf("add: %s \n", obj)
      },
      DeleteFunc: func(obj interface{}) {
        fmt.Printf("delete: %s \n", obj)
      },
      UpdateFunc:func(oldObj, newObj interface{}) {
        fmt.Printf("old: %s, new: %s \n", oldObj, newObj)
      },
    },
  )
  stop := make(chan struct{})
  go controller.Run(stop)

  for {
    time.Sleep(10 * time.Second)
  }
}
func watchEvents(clientset *kubernetes.Clientset) {
  watchlist := cache.NewListWatchFromClient(clientset.CoreV1().RESTClient(), "events", v1.NamespaceAll, fields.Everything())
  _, controller := cache.NewInformer(
    watchlist,
    &v1.Event{},
    time.Second * 0,
    cache.ResourceEventHandlerFuncs{
      AddFunc: func(obj interface{}) {
        fmt.Printf("add: %s \n", obj)
      },
      DeleteFunc: func(obj interface{}) {
        fmt.Printf("delete: %s \n", obj)
      },
      UpdateFunc:func(oldObj, newObj interface{}) {
        fmt.Printf("old: %s, new: %s \n", oldObj, newObj)
      },
    },
  )
  stop := make(chan struct{})
  go controller.Run(stop)

  for {
    time.Sleep(10 * time.Second)
  }
}

func listPods(clientset *kubernetes.Clientset) {
  for {
    pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
    if err != nil {
      panic(err.Error())
    }
    fmt.Printf("There are %d pods in the cluster\n", len(pods.Items))
    time.Sleep(10 * time.Second)
  }
}

func main() {
  kubeconfig := flag.String("kubeconfig", "/root/.kube/config", "absolute path to the kubeconfig file")
  flag.Parse()
  // uses the current context in kubeconfig
  config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
  if err != nil {
    panic(err.Error())
  }
  // creates the clientset
  clientset, err := kubernetes.NewForConfig(config)
  if err != nil {
    panic(err.Error())
  }
  //watchlist := cache.NewListWatchFromClient(clientset.Core().RESTClient(), "pods", v1.NamespaceDefault, fields.Everything())
  //watchPods(clientset)
  //listPods(clientset)
  //watchEvents(clientset)

  namespaces, err := clientset.CoreV1().Namespaces().List(metav1.ListOptions{})
  if err != nil {
    panic(err.Error())
  }
  fmt.Printf("There are %d namespaces in the cluster\n", len(namespaces.Items))
  // _, err := clientset.CoreV1().Namespaces().Create(&v1.Namespace{ObjectMeta: meta_v1.ObjectMeta{Name: name}})

}
