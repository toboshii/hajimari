package kube

import (
	"os"

	"github.com/toboshii/hajimari/internal/log"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	logger = log.New()
)

func getConfig() *rest.Config {
	var config *rest.Config

	config, _ = rest.InClusterConfig()

	configPath := os.Getenv("KUBECONFIG")
	if configPath == "" {
		configPath = os.Getenv("HOME") + "/.kube/config"
	}
	config, _ = clientcmd.BuildConfigFromFlags("", configPath)

	return config
}

// GetClient returns a k8s clientset
func GetClient() kubernetes.Interface {
	var client kubernetes.Interface

	config := getConfig()

	client, err := kubernetes.NewForConfig(config)
	if err != nil {
		logger.Fatalf("Cannot create new kubernetes client from config: %v", err)
	}

	return client
}

// GetClient returns a k8s clientset
func GetDynamicClient() dynamic.Interface {
	var client dynamic.Interface

	config := getConfig()

	client, err := dynamic.NewForConfig(config)
	if err != nil {
		logger.Fatalf("Cannot create new kubernetes client from config: %v", err)
	}

	return client
}
