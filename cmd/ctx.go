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
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

// ctxCmd represents the ctx command
var ctxCmd = &cobra.Command{
	Use:   "ctx",
	Short: "Switch or list all available contexts from a valid kubeconfig",
	Long: `With the ctx subcommand you can list all available contexts when you specify nothing but a ctx argument. 
	If you specify after the ctx argument one more parameter you immediately switch to it.
	More than 1 Parameter is not supported`,
	Args: func(cmd *cobra.Command, args []string) error {
		if !listFlag && len(args) > 1 {
			return errors.New("requires one or no parameters")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		yamlFile, err := ioutil.ReadFile(kubeConfigPath)

		if err != nil {
			fmt.Println("Error while loading kubeconfig: " + err.Error())
		}

		var kubeConfig KubeConfig

		err = yaml.Unmarshal(yamlFile, &kubeConfig)
		if err != nil {
			log.Fatalf("Unmarshal: %v", err)
		}

		if len(args) != 1 {
			for _, value := range kubeConfig.Contexts {
				fmt.Println(value.Name)
			}
			return
		} else {
			kubeConfig.CurrentContext = args[0]
			data, err := yaml.Marshal(kubeConfig)
			if err != nil {
				log.Fatalf("Cannot write kubeConfig file: %v", err)
			}
			ioutil.WriteFile(kubeConfigPath, data, 0)
		}
	},
}

func init() {
	rootCmd.AddCommand(ctxCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// ctxCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// ctxCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
