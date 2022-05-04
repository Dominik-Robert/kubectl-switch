/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	listFlag       bool
	kubeConfigPath string
)

type KubeConfig struct {
	APIVersion string `yaml:"apiVersion"`
	Clusters   []struct {
		Cluster struct {
			CertificateAuthorityData string `yaml:"certificate-authority-data"`
			Server                   string `yaml:"server"`
		} `yaml:"cluster"`
		Name string `yaml:"name"`
	} `yaml:"clusters"`
	Contexts []struct {
		Context struct {
			Namespace string `yaml:"namespace"`
			Cluster   string `yaml:"cluster"`
			User      string `yaml:"user"`
		} `yaml:"context"`
		Name string `yaml:"name"`
	} `yaml:"contexts"`
	CurrentContext string `yaml:"current-context"`
	Kind           string `yaml:"kind"`
	Preferences    struct {
	} `yaml:"preferences"`
	Users []struct {
		Name string `yaml:"name"`
		User struct {
			ClientCertificateData string `yaml:"client-certificate-data"`
			ClientKeyData         string `yaml:"client-key-data"`
		} `yaml:"user"`
	} `yaml:"users"`
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kubectl-switch",
	Short: "A little tool which helps you to switch faster to a namespace or a context",
	Long: `The binary provides two simple commands and a help function:
	
	the ns command provides a way to switch or list all namespaces, so if you want to list all namespaces you specify no more parameter just the ns argument.
	When you want to switch you add the namespaces right after the ns command
	
	the ctx command provides a way to switch or list all contexts, so if you want to list all contexts you specify no more parameter just the ctx argument.
	When you want to switch you add the contexts right after the ctx command`,
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

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kubectl-switch.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.

	dir, err := homedir.Dir()

	if err != nil {
		fmt.Println("error: homedir cannot resolved. " + err.Error())
		os.Exit(1)
	}

	dir += string(os.PathSeparator) + ".kube" + string(os.PathSeparator) + "config"

	rootCmd.PersistentFlags().StringVarP(&kubeConfigPath, "kubeConfig", "c", dir, "Prints a list of the selected resources")
}
