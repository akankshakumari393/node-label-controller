package main

import (
	"context"
	"flag"
	"os"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
)

func main() {
	kubeconfigFile := os.Getenv("HOME") + "/.kube/config"
	kubeconfig := flag.String("kubeconfig", kubeconfigFile, "Kubeconfig File location")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		klog.Infof("erorr %s building config from flags.. trying from inside the cluster", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			klog.Fatalf("Error building kubeconfig: %s", err.Error())
		}
	}
	kubeclient, err := kubernetes.NewForConfig(config)
	if err != nil {
		klog.Fatalf("Error building kubernetes clientset: %s", err.Error())
	}

	klog.Info("List all nodes ")
	nodes, err := kubeclient.CoreV1().Nodes().List(context.Background(), v1.ListOptions{})
	if err != nil {
		klog.Fatalf("error %s, listing nodes\n", err.Error())
	}

	for _, node := range nodes.Items {
		klog.Info(node.Name)
	}
}
