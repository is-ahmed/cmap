package main

import (
	"log"
	cmd "github.com/is-ahmed/command-map/cmd"
)


func init() {

	cmd.RootCmd.AddCommand(cmd.ListCmd)
	cmd.RootCmd.AddCommand(cmd.InitCmd)
	cmd.RootCmd.AddCommand(cmd.InsertCmd)
	cmd.RootCmd.AddCommand(cmd.SearchCmd)
	cmd.RootCmd.AddCommand(cmd.ClearCmd)
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
