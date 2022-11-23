package edgescrs

import (
	"context"
	"fmt"
	"log"
	"os"

	"difioperator/utils"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"

	edgev1 "github.com/difinative/Edge-MonitoringOperator/api/v1"
)

func CreateEdge(dynClient *dynamic.Interface, namespace string, edgeToCreate edgev1.Edge) error {

	resource := schema.GroupVersionResource{
		Group:    utils.EDGE_CR_GROUP,
		Version:  utils.EDGE_CR_VERSION,
		Resource: utils.EDGE_CR_RESOURCE,
	}
	log.Printf("Edge resource >> Group: %s, Version: %s, Kind: %s\n", resource.Group, resource.Version, resource.Resource)
	log.Println("Edge resource namespace: ", namespace)
	edgeObject, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&edgeToCreate)

	if err != nil {
		log.Printf("!!!! Error while trying to type cast the edge: %s to it's resource type !!!!\n", edgeToCreate.Name)
		log.Println(err)
	}
	unStructEdge := &unstructured.Unstructured{Object: edgeObject}
	_, err = (*dynClient).Resource(resource).Namespace(namespace).Create(context.TODO(), unStructEdge, metav1.CreateOptions{})

	return err
}

func GetEdge(dynClient *dynamic.Interface, namespace string, edgeName string) edgev1.Edge {

	resource := schema.GroupVersionResource{
		Group:    utils.EDGE_CR_GROUP,
		Version:  utils.EDGE_CR_VERSION,
		Resource: utils.EDGE_CR_RESOURCE,
	}
	log.Printf("Edge resource >> Group: %s, Version: %s, Kind: %s\n", resource.Group, resource.Version, resource.Resource)
	log.Println("Edge resource namespace: ", namespace)
	// edgeObject, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&edgeToCreate)

	// if err != nil {
	// 	log.Printf("!!!! Error while trying to type cast the edge: %s to it's resource type !!!!\n", edgeToCreate.Name)
	// 	log.Println(err)
	// }
	// unStructEdge := &unstructured.Unstructured{Object: edgeObject}
	// unStructEdge, err = (*dynClient).Resource(resource).Namespace(namespace).Create(context.TODO(), unStructEdge, metav1.CreateOptions{})
	fmt.Println("name", edgeName)
	unstructedge, err := (*dynClient).Resource(resource).Namespace(namespace).Get(context.TODO(), edgeName, metav1.GetOptions{})
	if err != nil {
		log.Println("Error while tring to get the edge: ", edgeName)
		log.Println("Error", err)
		os.Exit(1)
	}
	var edge edgev1.Edge
	err = runtime.DefaultUnstructuredConverter.FromUnstructured(unstructedge.Object, &edge)
	if err != nil {
		log.Println("Error while tring to parse the unstructed object of edge: ", edgeName)
		os.Exit(1)
	}
	return edge
}

func UpdateEdge(dynClient *dynamic.Interface, namespace string, edgeToUpdate edgev1.Edge) error {

	resource := schema.GroupVersionResource{
		Group:    utils.EDGE_CR_GROUP,
		Version:  utils.EDGE_CR_VERSION,
		Resource: utils.EDGE_CR_RESOURCE,
	}
	log.Printf("Edge resource >> Group: %s, Version: %s, Kind: %s\n", resource.Group, resource.Version, resource.Resource)
	log.Println("Edge resource namespace: ", namespace)
	edgeObject, err := runtime.DefaultUnstructuredConverter.ToUnstructured(&edgeToUpdate)

	if err != nil {
		log.Printf("!!!! Error while trying to type cast the edge: %s to it's resource type !!!!\n", edgeToUpdate.Name)
		log.Println(err)
	}
	unStructEdge := &unstructured.Unstructured{Object: edgeObject}
	_, err = (*dynClient).Resource(resource).Namespace(namespace).Update(context.TODO(), unStructEdge, metav1.UpdateOptions{})

	return err
}
