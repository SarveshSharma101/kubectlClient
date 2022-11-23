/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	k8client "difioperator/k8-client"
	edgescrs "difioperator/k8-client/Edges-CRs"
	"log"
	"os"

	edgev1 "github.com/difinative/Edge-MonitoringOperator/api/v1"
	"github.com/spf13/cobra"
)

var resolution string

// var cameraNamespace string

// var workingStatus string
var ip string

// addcameraCmd represents the addcamera command
var addcameraCmd = &cobra.Command{
	Example: `- kubectl squirrel addcamera <nameOfCamera> <nameOfEdgeWhereCameraBelongs> [flags]`,
	Use:     "addcamera",
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(2),

	Run: func(cmd *cobra.Command, args []string) {

		dynClient := k8client.GetClient()
		nameOfCamera := args[0]
		nameOfEdge := args[1]
		if namespace == "" {
			namespace = "default"
		}
		edge := edgescrs.GetEdge(&dynClient, namespace, nameOfEdge)
		var camera edgev1.Camera
		if resolution == "" {
			resolution = "1600x1200"
		}
		camera = edgev1.Camera{
			Resolution:    resolution,
			Workingstatus: "UP",
			IP:            ip,
		}
		if edge.Spec.Cameras == nil {
			cmap := map[string]edgev1.Camera{nameOfCamera: camera}
			edge.Spec.Cameras = cmap
		} else {
			_, isPresent := edge.Spec.Cameras[nameOfCamera]
			if isPresent {
				log.Println("Camera is already present in the edge")
				os.Exit(1)
			}

			edge.Spec.Cameras[nameOfCamera] = camera
		}
		err := edgescrs.UpdateEdge(&dynClient, namespace, edge)
		if err != nil {
			log.Println("Error while trying to add camera to edge: ", nameOfEdge)
			log.Println("Error: ", err)
			os.Exit(1)
		}

		log.Println("addcamera called")
	},
}

func init() {
	rootCmd.AddCommand(addcameraCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addcameraCmd.PersistentFlags().String("foo", "", "A help for foo")
	addcameraCmd.PersistentFlags().StringVarP(&resolution, "resolution", "r", "", "camera resolution ")
	addcameraCmd.PersistentFlags().StringVarP(&ip, "ip", "i", "", "camera ip")
	// addedgeCmd.PersistentFlags().StringVarP(&cameraNamespace, "cnamespace", "", "", "specifies in which namespace the edge should be created")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addcameraCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
