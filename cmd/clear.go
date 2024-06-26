package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/spf13/cobra"
)

var (
	ClearCmd = &cobra.Command {
		Use: "clear",
		Short: "Clear all entries from ~/.comandmap",
		Run: func(cmd *cobra.Command, args []string){
			user, _ := user.Current()
			filePath := user.HomeDir + "/.commandmap"
			if err := os.Truncate(filePath, 0); err != nil {
				log.Fatal("Failed to clear ~/.commandmap")
			} else {
				fmt.Println("All commands have been cleared")
			}
		},
	}
)
