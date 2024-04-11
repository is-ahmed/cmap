package cmd

import (
	"log"
	"os"
	"os/user"
	"github.com/spf13/cobra"
)

var (
	InitCmd = &cobra.Command{
		Use: "init",
		Short: "Intialize the mapping",
		Run: func(cmd *cobra.Command, args []string) {
			user, err := user.Current()	
			filePath := user.HomeDir + "/.commandmap"
			mapFile, err := os.Create(filePath)
			if err != nil {
				log.Fatal(err)
			}
			mapFile.Close()
		},
	}
)
