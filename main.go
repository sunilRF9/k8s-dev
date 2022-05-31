package main


import (
	"flag"
	"fmt"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func main() {
	kubeconfig := flag.String("kubeconfig", "/home/sunil/.kube/config", "location to your config file")
	config, err := clientcmd.BuildConfigFromFlags("",*kubeconfig)
	if err != nil {
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
	}
	ctx := context.Background()
	pods, err := clientset.CoreV1().Pods("default").List(ctx, metav1.ListOptions{})
	if err != nil {
	}
	fmt.Println("Pods from default ns")
	for _, pod := range pods.Items {
		fmt.Println(pod.Name)
	}
	ingress, err := clientset.NetworkingV1().Ingresses("default").List(ctx, metav1.ListOptions{})
	if err != nil {
	}
	fmt.Println("Listing Ingresses")
	for _, ing := range ingress.Items {
		fmt.Println(ing.Name)
	}

}
