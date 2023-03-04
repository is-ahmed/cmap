package tui

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/is-ahmed/command-map/types"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type (
	errMsg error
)
var (
	titleStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Right = "├"
		return lipgloss.NewStyle().BorderStyle(b).Padding(0, 1)
	}()

	infoStyle = func() lipgloss.Style {
		b := lipgloss.RoundedBorder()
		b.Left = "┤"
		return titleStyle.Copy().BorderStyle(b)
	}()
)
type Model struct {
	textInput textinput.Model
	viewport viewport.Model
	ready bool
	err error	
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()

	return Model{
		textInput: ti,
		err:       nil,
		ready: false,
	}	
}


type TickMsg time.Time

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	
	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(m.headerView())
		footerHeight := lipgloss.Height(m.footerView())
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			// Since this program is using the full size of the viewport we
			// need to wait until we've received the window dimensions before
			// we can initialize the viewport. The initial dimensions come in
			// quickly, though asynchronously, which is why we wait for them
			// here.
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}

	}

	var cmds []tea.Cmd
	var searchResults string = getSearchResults(m.textInput.Value())
	m.viewport.SetContent(searchResults)
	// Handle keyboard and mouse events in the viewport
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)
	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)

}
func (m Model) headerView() string {
	title := titleStyle.Render("Results")
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(title)))
	return lipgloss.JoinHorizontal(lipgloss.Center, title, line)
}

func (m Model) footerView() string {
	info := infoStyle.Render(fmt.Sprintf("%3.f%%", m.viewport.ScrollPercent()*100))
	line := strings.Repeat("─", max(0, m.viewport.Width-lipgloss.Width(info)))
	return lipgloss.JoinHorizontal(lipgloss.Center, line, info)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}


func getSearchResults(command string) string {
	var commandList []types.Command = types.GetCommands().Commands

	sort.Slice(commandList, func(i, j int) bool {
		return -1 * fuzzy.RankMatch(command, commandList[i].Command) <  -1 * fuzzy.RankMatch(command, commandList[j].Command)
	})

	return formatSearchResults(commandList)

}

func formatSearchResults(results []types.Command) string {
	var output string = ""
	for i:=0; i < len(results); i++ {
		output += results[i].Command + " : " + results[i].Description + "\n"
	}
	return output
}



func (m Model) View() string {
	return fmt.Sprintf("%s\n%s\n%s", m.textInput.View(), m.headerView(), m.viewport.View())
}

