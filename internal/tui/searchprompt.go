package tui


import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	input string

}

type Model int

type TickMsg time.Time

func (m Model) Init() tea.Cmd {
	return tea.Batch(tick(), tea.EnterAltScreen)
}

func (m Model) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		}

	case TickMsg:
		m--
		if m <= 0 {
			return m, tea.Quit
		}
		return m, tick()
	}

	return m, nil
}

func (m Model) View() string {
	return fmt.Sprintf("\n\n     Hi. This program will exit in %d seconds...", m)
}

func tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}
