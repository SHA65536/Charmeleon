package server

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var AppStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.ThickBorder()).
	BorderForeground(lipgloss.Color("63"))

type AppModel struct {
	Menu  *MenuModel
	Image *PokeModel
}

func (m *AppModel) Init() tea.Cmd { return nil }
func (m *AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case " ", "enter":
			idx := (m.Menu.List.Paginator.Page * m.Menu.List.Paginator.PerPage) + m.Menu.List.Cursor() + 1
			m.Image.UpdatePokemon(fmt.Sprintf("%03d", idx))
		case "left", "a":
			m.Image.UpdateForm(-1)
		case "right", "d":
			m.Image.UpdateForm(1)
		case "up", "down", "w", "s":
			m.Menu.Update(msg)
		}
	case tea.WindowSizeMsg:
	}
	return m, cmd
}
func (m *AppModel) View() string {
	return AppStyle.Render(lipgloss.JoinHorizontal(lipgloss.Left, m.Menu.View(), m.Image.View()))
}

func InitialPage() *AppModel {
	self := &AppModel{InitialMenu(), InitialPoke()}
	return self
}
