package server

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var PokeStyle = lipgloss.NewStyle().
	Width(68).Height(32).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63"))

type PokeModel struct {
	Content string //68 x 28(56)
}

func (m *PokeModel) Init() tea.Cmd { return nil }
func (m *PokeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		}
	case tea.WindowSizeMsg:
	}
	return m, cmd
}
func (m *PokeModel) View() string {
	return PokeStyle.Render(m.Content)
}

func InitialPoke() *PokeModel {
	logo, _ := os.ReadFile("logo.cow")
	return &PokeModel{string(logo)}
}
