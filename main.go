/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import "github.com/dominik-robert/kubectl-switch/cmd"

var (
	Version   = "1.0.0"
	BuildTime = "2015-08-01 UTC"
	GitHash   = ""
)

func main() {
	cmd.Execute()
}
