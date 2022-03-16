package server

import (
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	boxer "github.com/treilik/bubbleboxer"
)

type MainPage struct {
	tui boxer.Boxer
}

func InitialPage() MainPage {
	dat, _ := os.ReadFile("data/regular/yveltal.png.cow")
	items := []list.Item{
		item{title: "Pikachu", desc: "Index 001"},
		item{title: "Mikachu", desc: "Index 002"},
		item{title: "Likachu", desc: "Index 003"},
		item{title: "Dikachu", desc: "Index 004"},
	}
	menu := MenuModel{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	menu.list.Title = "Pokemon List"

	poke := PokeModel(string(dat))

	m := MainPage{tui: boxer.Boxer{}}
	m.tui.LayoutTree = boxer.Node{
		// orientation
		VerticalStacked: false,
		// spacing
		SizeFunc: func(_ boxer.Node, widthOrHeight int) []int {
			return []int{widthOrHeight / 2, widthOrHeight / 2}
		},
		Children: []boxer.Node{
			// make sure to encapsulate the models into a leaf with CreateLeaf:
			m.tui.CreateLeaf("left", &menu),
			m.tui.CreateLeaf("right", &poke),
		},
	}
	return m
}

func (m MainPage) Init() tea.Cmd {
	return nil
}

func (m MainPage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.tui.UpdateSize(msg)
	}
	m.tui.ModelMap["left"].Update(msg)
	m.tui.ModelMap["right"].Update(msg)
	return m, nil
}
func (m MainPage) View() string { return m.tui.View() }
