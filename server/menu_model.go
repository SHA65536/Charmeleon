package server

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type MenuModel struct {
	pokemodel *PokeModel
	list      list.Model
}

func (m *MenuModel) Init() tea.Cmd { return nil }
func (m *MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case " ", "enter":
			idx := (m.list.Paginator.Page * m.list.Paginator.PerPage) + m.list.Cursor() + 1
			m.pokemodel.UpdateIndex(fmt.Sprintf("%03d", idx))
		case "left", "right", "a", "d":
			return m, nil
		}
	case tea.WindowSizeMsg:
		top, right, bottom, left := docStyle.GetMargin()
		m.list.SetSize(msg.Width-left-right, msg.Height-top-bottom)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
func (m *MenuModel) View() string {
	return docStyle.Render(m.list.View())
}

type item struct {
	title, desc, idx string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }
