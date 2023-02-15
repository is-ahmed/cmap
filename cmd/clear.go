package cmd

import (
	"log"
	"github.com/spf13/cobra"
	"os"
)

var (
	ClearCmd = &cobra.Command {
		Use: "clear",
		Short: "Clear all entries from ~/.comandmap",
		Run: func(cmd *cobra.Command, args []string){
			if err := os.Truncate("/home/isahmed/.commandmap", 0); err != nil {
				log.Fatal("Failed to clear ~/.commandmap")
			}
		},
	}
)
