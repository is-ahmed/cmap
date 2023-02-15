package types
import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"encoding/json"
)
type Commands struct {
	Commands []Command
}

type Command struct {
	Command string
	Description string
}

func GetCommands() Commands {
	var commandList Commands
	commandMap, err := os.Open("/home/isahmed/.commandmap")
	if err != nil {
		log.Fatal(err)
	}
	byteValue, _ := ioutil.ReadAll(commandMap)
	json.Unmarshal(byteValue, &commandList)
	return commandList
}

func WriteCommands(commandList Commands) {
    byteValue, _ := json.MarshalIndent(commandList, "", " ")
	_ = ioutil.WriteFile("/home/isahmed/.commandmap", byteValue, 0644)
}

func (c Command) Print() {
	fmt.Println(c.Command + " : " + c.Description)
}

func (c Commands) Print() {
	for i := 0; i < len(c.Commands); i++ {
		c.Commands[i].Print()
	}
}
