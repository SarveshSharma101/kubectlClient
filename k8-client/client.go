package k8client

import (
	"flag"
	"log"

	"difioperator/utils"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/tools/clientcmd"
)

func GetClient() dynamic.Interface {

	log.Println("Loading kubeconfig file...")
	kubeconfig := flag.String(utils.KUBECONFIG, utils.KUBECONFIG_PATH, "path to load kubeconfig")
	log.Println("Loading the configs from kubeconfig")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Println("!!!!!!!!! Error while trying to load config file using kube config file !!!!!!!!!!!!")
		log.Print("!!! Error >>", err)
	}

	log.Println("Initiating a dynamic k8 client")
	dynClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Println("!!!!!!!!! Error while trying to crete dynamic client !!!!!!!!!!!!")
		log.Println("!!! Error >> ", err)
	}

	return dynClient
}
