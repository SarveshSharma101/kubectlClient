/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var namespace string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Example: `- kubectl squirrel addedge <nameOfEdge> [flags]
- kubectl squirrel addcamera <nameOfCamera> <nameOfEdgeWhereCameraBelongs> [flags]`,
	Use:   "squirrel",
	Short: "'kubectl squirrel' helps you create instance of Custom resource in your k8 cluster",
	Long: `Following is the list of custom resources supported:
1. Edge
Note: You should have the CRD's of the above listed resources installed in your k8 cluster`,
	Version: "0.1.0",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.difioperator.yaml)")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "", "specifies the namespace of the edges")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
