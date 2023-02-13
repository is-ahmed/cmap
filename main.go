package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
//	"strings"
//	"github.com/lithammer/fuzzysearch/fuzzy"
	"github.com/spf13/cobra"
)

type Commands struct {
	Commands []Command
}

type Command struct {
	command string
	description string
}


func (c Command) print() {
	fmt.Println(c.command + " : " + c.description)
}

func (c Commands) print() {
	for i := 0; i < len(c.Commands); i++ {
		c.Commands[i].print()
	}
}

func checkInit() {
	// Check if application has been initialized
			
	
}

func getCommands() Commands {
	var commandList Commands
	commandMap, err := os.Open("/home/isahmed/.commandmap")
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(commandMap)
	json.Unmarshal(byteValue, &commandList)
	return commandList
}

func writeCommands(commandList Commands) {
	//fmt.Println(commandList)
	sampleCommand := Commands{
		Commands: []Command{
			{
				command: "sudo apt update",
				description: "refresh repos",
			},
		},
	}
    byteValue, _ := json.MarshalIndent(sampleCommand, "", " ")
	_ = ioutil.WriteFile("/home/isahmed/.commandmap", byteValue, 0644)
}

var (
	persistRootFlag bool;
	localRootFlag bool;
    rootCmd = &cobra.Command{
	 Use: "Command Map",
	 Short: "Map for common commands you use but not often enough to commit to memory",
    }
	listCmd = &cobra.Command{
		Use: "list",
		Short: "Prints all commands w/descriptions to standard out",
		Run: func(cmd *cobra.Command, args []string) {
			var commandList Commands = getCommands()
			commandList.print()	
		},
	}
	initCmd = &cobra.Command{
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
	insertCmd = &cobra.Command{
		Use: "insert",
		Short: "Create a new command to insert into the map",
		Run: func(cmd *cobra.Command, args []string) {

			var command string = args[0]
			var description string = args[1]
			fmt.Println(command)
			fmt.Println(description)
			
			var c Command
			c.command = command
			c.description = description

			commandList := getCommands()
			
			commandList.Commands = append(commandList.Commands, c)
			writeCommands(commandList)

		},
			
	}
	searchCmd = &cobra.Command {
		Use: "search",
		Short: "Search for a command by the description",
		Run: func(cmd *cobra.Command, args []string){
//			commandList := getCommands()
		},
	}
	clearCmd = &cobra.Command {
		Use: "clear",
		Short: "Clear all entries from ~/.comandmap",
		Run: func(cmd *cobra.Command, args []string){
			if err := os.Truncate("/home/isahmed/.commandmap", 0); err != nil {
				log.Fatal("Failed to clear ~/.commandmap")
			}
		},
	}
	
)

func init() {

	// Persistent flags, apply to sub commands as well
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlag", "p", false, "a persistent root flag")

	// Local flags, only applies to root cmd
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "l", false, "a local root flag")

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(insertCmd)
	rootCmd.AddCommand(searchCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
