package main

import (
	"context"
	"flag"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/sat/.kube/config", "location of kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	checkErr(err)
	clientset, err := kubernetes.NewForConfig(config)
	checkErr(err)
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("clamav").List(ctx, metav1.ListOptions{})
	for _, p := range pods.Items {
		fmt.Println("Pods", p.Name)
	}
	deployments, err := clientset.AppsV1().Deployments("clamav").List(ctx, metav1.ListOptions{})
	checkErr(err)
	for _, d := range deployments.Items {
		fmt.Println("Deployment", d.Name)
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err) //stop the program execution
	}
}
