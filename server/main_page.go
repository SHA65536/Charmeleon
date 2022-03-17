package server

import (
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
		if msg.String() == "q" || msg.String() == "ctrl+c" {
			return m, tea.Quit
		} else {
			m.Menu.Update(msg)
			m.Image.Update(msg)
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
