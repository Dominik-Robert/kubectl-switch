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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if !listFlag && len(args) != 1 {
			return errors.New("requires at least one arg")
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

		if listFlag {
			for _, value := range kubeConfig.Contexts {
				fmt.Println(value.Name)
			}
			return
		}

		kubeConfig.CurrentContext = args[0]
		data, err := yaml.Marshal(kubeConfig)
		if err != nil {
			log.Fatalf("Cannot write kubeConfig file: %v", err)
		}
		ioutil.WriteFile(kubeConfigPath, data, 0)

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
