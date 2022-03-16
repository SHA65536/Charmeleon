package server

import tea "github.com/charmbracelet/bubbletea"

type PokeModel string

func (s PokeModel) String() string {
	return string(s)
}

// satisfy the tea.Model interface
func (s PokeModel) Init() tea.Cmd                           { return nil }
func (s PokeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) { return s, nil }
func (s PokeModel) View() string                            { return s.String() }
