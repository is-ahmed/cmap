package cmd

import (
	"github.com/is-ahmed/cmap/types"
	"github.com/spf13/cobra"
)

var (
	ListCmd = &cobra.Command{
		Use: "list",
		Short: "Prints all commands w/descriptions to standard out",
		Run: func(cmd *cobra.Command, args []string) {
			var commandList types.Commands = types.GetCommands()
			commandList.Print()	
		},
	}
)
