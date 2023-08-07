package main

import (
	"log"
	"github.com/is-ahmed/cmap/cmd"
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
