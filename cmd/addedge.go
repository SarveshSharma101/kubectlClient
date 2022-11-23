/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	k8client "difioperator/k8-client"
	"difioperator/utils"
	"log"
	"strings"

	edgescrs "difioperator/k8-client/Edges-CRs"

	edgev1 "github.com/difinative/Edge-MonitoringOperator/api/v1"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// var cameras string
var zip string
var latitude string
var longitude string
var city string
var availableMemory string
var inferenceLastUpdate int16

// addedgeCmd represents the addedge command
var addedgeCmd = &cobra.Command{
	Example: `- kubectl squirrel addedge <nameOfEdge> [flags]`,
	Use:     "addedge",
	Short:   "creates a new instance of the custom resource -> edge",
	Long: `It is used to create a new instance of the edge in your k8 cluster.
Pass the name of the edge as an argument to the command.`,
	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		// cameraArr := []edgev1.Camera{}

		nameOfEdge := args[0]
		name := strings.ToLower(nameOfEdge)
		if namespace == "" {
			namespace = "default"
		}
		// if cameras != "" {
		// 	cameraArr = strings.Split(cameras, ",")
		// }
		if availableMemory == "" {
			availableMemory = "5G"
		}
		if inferenceLastUpdate == 0 {
			inferenceLastUpdate = 20
		}
		dynClient := k8client.GetClient()
		edgeToCreate := edgev1.Edge{
			TypeMeta: v1.TypeMeta{
				Kind:       utils.EDGE_CR_KIND,
				APIVersion: utils.EDGE_CR_GROUP + "/" + utils.EDGE_CR_VERSION,
			},
			ObjectMeta: v1.ObjectMeta{
				Name:      name,
				Namespace: namespace,
			},
			Spec: edgev1.EdgeSpec{
				Vitals: edgev1.Vitals{
					Working:                   "UP",
					AvailableMemory:           availableMemory,
					InferenceServerLastUpdate: int(inferenceLastUpdate),
				},
				Region: edgev1.Region{
					Zip:       zip,
					Latitude:  latitude,
					Longitude: longitude,
					City:      city,
				},
				Edgename: nameOfEdge,
				// Cameras:  cameraArr,
			},
		}
		err := edgescrs.CreateEdge(&dynClient, namespace, edgeToCreate)
		if err != nil {
			log.Println(err)
		} else {
			log.Println("Created an edge with name: ", edgeToCreate.ObjectMeta.Name)
		}
	},
}

func init() {
	rootCmd.AddCommand(addedgeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addedgeCmd.PersistentFlags().StringVarP(&cameras, "cameras", "c", "", "specifies the cameras associated with the edge")
	addedgeCmd.PersistentFlags().StringVarP(&zip, "zip", "z", "", "Zip code of the edge site")
	addedgeCmd.PersistentFlags().StringVarP(&latitude, "latitude", "l", "", "latitude ")
	addedgeCmd.PersistentFlags().StringVarP(&longitude, "longitude", "L", "", "longitude")
	addedgeCmd.PersistentFlags().StringVarP(&city, "city", "C", "", "city")
	addedgeCmd.PersistentFlags().StringVarP(&availableMemory, "availablememory", "m", "", "Available memory")
	addedgeCmd.PersistentFlags().Int16VarP(&inferenceLastUpdate, "inferencelastupdate", "i", 0, "Last update to inference server")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addedgeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
