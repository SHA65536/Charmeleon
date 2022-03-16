package server

import (
	"charmeleon/pokemon"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"
)

var Pokedex map[string]*pokemon.Pokemon

type CharmServ struct {
	Serv *ssh.Server
}

func MakeCharmServ(conn string) (*CharmServ, error) {
	var err error
	file, err := os.ReadFile("data/sprites.json")
	err = json.Unmarshal(file, &Pokedex)
	if err != nil {
		log.Fatal(err)
	}
	self := &CharmServ{}
	self.Serv, err = wish.NewServer(
		wish.WithAddress(conn),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)
	return self, err
}

func (self *CharmServ) Start() {
	fmt.Println("Starting SSH server on " + self.Serv.Addr)
	if err := self.Serv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func (self *CharmServ) Stop() {
	fmt.Println("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := self.Serv.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	_, _, active := s.Pty()
	if !active {
		fmt.Println("no active terminal, skipping")
		return nil, nil
	}

	return InitialPage(), []tea.ProgramOption{tea.WithAltScreen()}
}
