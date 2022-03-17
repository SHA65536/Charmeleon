package server

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var PokeStyle = lipgloss.NewStyle().
	Width(68).Height(28).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63"))
var DetailStyle = lipgloss.NewStyle().Bold(true).Width(36).Height(4)
var Intructions = DetailStyle.Width(24).Render(`
Use ^v to browse%s
Use <> to change form%s
Use space to select%s`)

type PokeModel struct {
	Content string //68 x 28(56)
	Index   string
	Form    int
	Details string
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
	return PokeStyle.Render(m.Content + m.Details)
}

func InitialPoke() *PokeModel {
	logo, _ := os.ReadFile("logo.cow")
	intro := fmt.Sprintf(Intructions, "           Pokedex Over SSH!",
		"      Made with charmbracelet/charm toolkit", "        Source at github.com/SHA65536/Charmeleon")
	return &PokeModel{string(logo), "001", 0, intro}
}

func (m *PokeModel) UpdateDetails() {
	newdetails := fmt.Sprintf(Intructions, "              Name: "+Pokedex[m.Index].Name,
		"         Index: "+m.Index, "           Form: "+Pokedex[m.Index].Forms[m.Form].Name)
	m.Details = newdetails
}

func (m *PokeModel) UpdatePokemon(idx string) {
	m.Index = idx
	m.Form = 0
	dat, _ := os.ReadFile(Pokedex[idx].Forms[0].Cow)
	m.Content = string(dat)
	m.UpdateDetails()
}

func (m *PokeModel) UpdateForm(direction int) {
	if (m.Form + direction) < 0 {
		m.Form = len(Pokedex[m.Index].Forms) - 1
	} else {
		m.Form = (m.Form + direction) % len(Pokedex[m.Index].Forms)
	}
	dat, _ := os.ReadFile(Pokedex[m.Index].Forms[m.Form].Cow)
	m.Content = string(dat)
	m.UpdateDetails()
}
