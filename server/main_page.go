package server

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	boxer "github.com/treilik/bubbleboxer"
)

type MainPage struct {
	tui boxer.Boxer
}

func InitialPage() MainPage {
	dat, _ := os.ReadFile("logo.cow")
	items := make([]list.Item, 0)
	for i := 1; i <= len(Pokedex); i++ {
		k := fmt.Sprintf("%03d", i)
		v := Pokedex[k]
		items = append(items, item{title: v.Name, desc: "#" + k + ": " + v.Name, idx: k})
	}

	poke := PokeModel{index: "001", cur_form: 0, cow: string(dat)}

	menu := MenuModel{pokemodel: &poke, list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	menu.list.SetShowPagination(false)

	menu.list.Title = "Pokemon List"

	m := MainPage{tui: boxer.Boxer{}}
	m.tui.LayoutTree = boxer.Node{
		// orientation
		VerticalStacked: false,
		// spacing
		SizeFunc: func(_ boxer.Node, widthOrHeight int) []int {
			menuWidth := widthOrHeight / 2
			pokeWidth := widthOrHeight - menuWidth
			return []int{menuWidth, pokeWidth}
		},
		Children: []boxer.Node{
			m.tui.CreateLeaf("menu", &menu),
			m.tui.CreateLeaf("poke", &poke),
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
	m.tui.ModelMap["menu"].Update(msg)
	m.tui.ModelMap["poke"].Update(msg)
	return m, nil
}
func (m MainPage) View() string { return m.tui.View() }
