package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/is-ahmed/cmap/types"
	"github.com/spf13/cobra"
)


var (
	InsertCmd = &cobra.Command{
		Use: "insert",
		Short: "Create a new command to insert into the map",
		Run: func(cmd *cobra.Command, args []string) {

			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Pass a command: ")
			command, _ := reader.ReadString('\n')

			command = strings.TrimSpace(command)

			fmt.Print("Pass a description: ")
			description, _ := reader.ReadString('\n')
		
			var c types.Command
			c.Command = command
			c.Description = description

			commandList := types.GetCommands()
			
			commandList.Commands = append(commandList.Commands, c)
			types.WriteCommands(commandList)

		},
			
	}
)
