package cmd
import (
	"github.com/spf13/cobra"
	"github.com/is-ahmed/command-map/types"
)


var (
	InsertCmd = &cobra.Command{
		Use: "insert",
		Short: "Create a new command to insert into the map",
		Run: func(cmd *cobra.Command, args []string) {

			var command string = args[0]
			var description string = args[1]
			
			var c types.Command
			c.Command = command
			c.Description = description

			commandList := types.GetCommands()
			
			commandList.Commands = append(commandList.Commands, c)
			types.WriteCommands(commandList)

		},
			
	}
)
