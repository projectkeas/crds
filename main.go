package main

import (
	"context"
	"flag"
	"fmt"
	"path/filepath"

	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"

	client "github.com/projectkeas/crds/pkg/client/clientset/versioned"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {

	namespace := "keas"

	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	// use the current context in kubeconfig
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err.Error())
	}

	client, err := client.NewForConfig(config)
	if err != nil {
		fmt.Println("Unable to create our clientset")
		return
	}

	policies, err := client.KeasV1alpha1().IngestionPolicies(namespace).List(context.Background(), metav1.ListOptions{})
	if err == nil {
		for _, policy := range policies.Items {
			fmt.Printf("%s - %s\n", policy.Name, policy.Spec.Description)
		}
	} else {
		fmt.Println(err.Error())
	}
}
