/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// nsCmd represents the ns command
var nsCmd = &cobra.Command{
	Use:   "ns",
	Short: "Switch or list all available namespaces from a valid kubeconfig",
	Long: `With the ns subcommand you can list all available namespaces when you specify nothing but a ns argument. 
	If you specify after the ns argument one more parameter you immediately switch to it.
	More than 1 Parameter is not supported`,
	Args: func(cmd *cobra.Command, args []string) error {
		if !listFlag && len(args) > 1 {
			return errors.New("requires one or no parameters")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			yamlFile, err := ioutil.ReadFile(kubeConfigPath)

			if err != nil {
				fmt.Println("Error while loading kubeconfig: " + err.Error())
			}

			var kubeConfig KubeConfig
			err = yaml.Unmarshal(yamlFile, &kubeConfig)
			if err != nil {
				log.Fatalf("Unmarshal: %v", err)
			}

			for i := range kubeConfig.Contexts {
				if kubeConfig.Contexts[i].Name == kubeConfig.CurrentContext {
					fmt.Println("Setting Context")
					kubeConfig.Contexts[i].Context.Namespace = args[0]

					data, err := yaml.Marshal(kubeConfig)
					if err != nil {
						log.Fatalf("Cannot write kubeConfig file: %v", err)
					}
					ioutil.WriteFile(kubeConfigPath, data, 0)
				}
			}
		} else {
			config, err := clientcmd.BuildConfigFromFlags("", kubeConfigPath)
			if err != nil {
				panic(err.Error())
			}

			// create the clientset
			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				panic(err.Error())
			}

			namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})

			if err != nil {
				panic(err.Error())
			}

			for _, namespace := range namespaces.Items {
				fmt.Println(namespace.Name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(nsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
