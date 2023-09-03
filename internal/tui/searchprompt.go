package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/atotto/clipboard"
	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/is-ahmed/cmap/types"
	"github.com/sahilm/fuzzy"
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

	baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))
)
type Model struct {
	textInput textinput.Model
	viewport viewport.Model
	table table.Model
	ready bool
	err error	
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = ""
	ti.Focus()

	columns := []table.Column{
		{Title: "Command", Width: 20},
		{Title: "Description", Width: 20},
	}

	rows := []table.Row{}

	tb := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(20),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	tb.SetStyles(s)

	
	return Model{
		textInput: ti,
		table: tb,
		err:       nil,
		ready: false,
	}	
}


type TickMsg time.Time

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			if m.table.Focused() {
				m.table.Blur()
			} else {
				m.table.Focus()
			}
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			var command string = m.table.SelectedRow()[0]
			clipboard.WriteAll(command)
		}

	// We handle errors just like any other message
	case errMsg:
		m.err = msg
		return m, nil
	
//	case tea.WindowSizeMsg:
//		headerHeight := lipgloss.Height(m.headerView())
//		footerHeight := lipgloss.Height(m.footerView())
//		verticalMarginHeight := headerHeight + footerHeight
//
//		if !m.ready {
//			// Since this program is using the full size of the viewport we
//			// need to wait until we've received the window dimensions before
//			// we can initialize the viewport. The initial dimensions come in
//			// quickly, though asynchronously, which is why we wait for them
//			// here.
//			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
//			m.ready = true
//		} else {
//			m.viewport.Width = msg.Width
//			m.viewport.Height = msg.Height - verticalMarginHeight
//		}

	}

	var cmds []tea.Cmd

	var searchResults []table.Row = getSearchResults(m.textInput.Value());
	m.table.SetRows(searchResults)

	m.table, cmd = m.table.Update(msg)
	cmds = append(cmds, cmd)

	m.textInput, cmd = m.textInput.Update(msg)
	cmds = append(cmds, cmd)

	m.viewport.SetContent("sample")
	m.viewport, cmd = m.viewport.Update(msg)
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

func contains(needle int, haystack []int) bool {
	for _, i := range haystack {
		if needle == i {
			return true
		}
	}
	return false
}


func getSearchResults(command string) []table.Row {
	/*
		Give a new list of results corresponding to the rows of the table
		Each member is of type table.Row
	*/
	var results []table.Row


	var commandList []types.Command = types.GetCommands().Commands

	var data []string

	for i := 0; i < len(commandList); i++ {
		data = append(data, commandList[i].Command + " : " + commandList[i].Description)
	}

	matches := fuzzy.Find(command, data)
	for _, match := range matches {
		command, desc := strings.Split(match.Str, " : ")[0], strings.Split(match.Str, " : ")[1]
		results = append(results, table.Row{command, desc})
	}

	return results

}




func (m Model) View() string {
	return fmt.Sprintf("\n%s\n\n%s", m.textInput.View(), baseStyle.Render(m.table.View()))
}

