package cmd

import (
	"log"
	"os"
	"github.com/spf13/cobra"
)

var (
	InitCmd = &cobra.Command{
		Use: "init",
		Short: "Intialize the mapping",
		Run: func(cmd *cobra.Command, args []string) {
			mapFile, err := os.Create("/home/isahmed/.commandmap")
			if err != nil {
				log.Fatal(err)
			}
			mapFile.Close()
		},
	}
)
