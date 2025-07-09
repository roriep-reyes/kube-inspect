/*
Copyright © 2025 RORIE REYES <roriep.reyes@gmail.com>
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

var namespace string

// podsCmd represents the pods command
var podsCmd = &cobra.Command{
	Use:   "pods",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config, err := rest.InClusterConfig()
		if err != nil {
			config, err = clientcmd.BuildConfigFromFlags("", filepath.Join(os.Getenv("HOME"), ".kube", "config"))
			if err != nil {
				fmt.Printf("Error creating config: %v\n", err)
				os.Exit(1)
			}
		}

		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			fmt.Printf("Error creating clientset: %v\n", err)
			os.Exit(1)
		}

		pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			fmt.Printf("Error listing pods: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Found %d pods:\n", len(pods.Items))
		for _, pod := range pods.Items {
			fmt.Printf("- %s (%s)\n", pod.Name, pod.Status.Phase)
		}
	},
}

func init() {
	rootCmd.AddCommand(podsCmd)

	podsCmd.Flags().StringVarP(&namespace, "namespace", "n", "", "Namespace to list pods from (default is all namespaces)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
