package cmd

import (
	"log"
	"github.com/spf13/cobra"
	tea "github.com/charmbracelet/bubbletea"
	tui "github.com/is-ahmed/command-map/internal/tui"
)
var (
	SearchCmd = &cobra.Command {
		Use: "search",
		Short: "Search for a command by the description",
		Run: func(cmd *cobra.Command, args []string){

			p := tea.NewProgram(tui.Model(5), tea.WithAltScreen())
			if _, err := p.Run(); err != nil {
				log.Fatal(err)
			}
		},
	}
)
