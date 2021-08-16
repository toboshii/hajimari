package kube

import (
	"os"

	"github.com/toboshii/hajimari/internal/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	logger = log.New()
)

// GetClient returns a k8s clientset
func GetClient() kubernetes.Interface {
	var kubeClient kubernetes.Interface
	_, err := rest.InClusterConfig()
	if err != nil {
		kubeClient = getClientOutOfCluster()
	} else {
		kubeClient = getClientInCluster()
	}

	return kubeClient
}

// GetClientInCluster returns a k8s clientset to the request from inside of cluster
func getClientInCluster() kubernetes.Interface {
	config, err := rest.InClusterConfig()
	if err != nil {
		logger.Fatalf("Can not get kubernetes config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Fatalf("Can not create kubernetes client: %v", err)
	}

	return clientset
}

func buildOutOfClusterConfig() (*rest.Config, error) {
	kubeconfigPath := os.Getenv("KUBECONFIG")
	if kubeconfigPath == "" {
		kubeconfigPath = os.Getenv("HOME") + "/.kube/config"
	}
	return clientcmd.BuildConfigFromFlags("", kubeconfigPath)
}

// GetClientOutOfCluster returns a k8s clientset to the request from outside of cluster
func getClientOutOfCluster() kubernetes.Interface {
	config, err := buildOutOfClusterConfig()
	if err != nil {
		logger.Fatalf("Cannot get kubernetes config: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		logger.Fatalf("Cannot create new kubernetes client from config: %v", err)
	}

	return clientset
}
