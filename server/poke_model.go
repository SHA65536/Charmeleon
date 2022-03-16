package server

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type PokeModel struct {
	index    string
	cur_form int
	cow      string
}

func (s *PokeModel) UpdateIndex(idx string) {
	s.index = idx
	s.cur_form = 0
	dat, _ := os.ReadFile(Pokedex[idx].Forms[0].Cow)
	s.cow = string(dat)
}

func (s *PokeModel) UpdateForm(add int) {
	if (s.cur_form + add) < 0 {
		s.cur_form = len(Pokedex[s.index].Forms) - 1
	} else {
		s.cur_form = (s.cur_form + add) % len(Pokedex[s.index].Forms)
	}
	dat, _ := os.ReadFile(Pokedex[s.index].Forms[s.cur_form].Cow)
	s.cow = string(dat)
}

// satisfy the tea.Model interface
func (s *PokeModel) Init() tea.Cmd { return nil }
func (s *PokeModel) View() string  { return s.cow }
func (s *PokeModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "a", "left":
			s.UpdateForm(-1)
		case "d", "right":
			s.UpdateForm(1)
		}
	}
	return s, nil
}
