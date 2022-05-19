package server

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var MenuStyle = lipgloss.NewStyle().
	Height(32).Width(20).
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("63"))

type MenuModel struct {
	List list.Model
}

func (m *MenuModel) Init() tea.Cmd { return nil }
func (m *MenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "w":
			m.List.CursorUp()
		case "s":
			m.List.CursorDown()
		}
		m.List, cmd = m.List.Update(msg)
		return m, cmd
	case tea.WindowSizeMsg:
	}
	return m, cmd
}
func (m *MenuModel) View() string {
	return MenuStyle.Render(m.List.View())
}

func InitialMenu() *MenuModel {
	items := make([]list.Item, 0)
	for i := 1; i <= len(Pokedex); i++ {
		k := fmt.Sprintf("%03d", i)
		v := Pokedex[k]
		items = append(items, item{title: v.Name, desc: "#" + k + ": " + v.Name, idx: k})
	}
	pokeList := list.New(items, list.NewDefaultDelegate(), 0, 32)
	pokeList.Title = "Pokedex"
	pokeList.SetShowPagination(false)
	pokeList.SetShowHelp(false)

	return &MenuModel{pokeList}
}

type item struct {
	title, desc, idx string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }
