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
	file, _ := os.ReadFile("data/sprites.json")
	err = json.Unmarshal(file, &Pokedex)
	if err != nil {
		log.Fatal(err)
	}
	self := &CharmServ{}
	self.Serv, err = wish.NewServer(
		wish.WithAddress(conn),
		wish.WithHostKeyPath(".ssh/term_info_ed25519"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)
	return self, err
}

func (srv *CharmServ) Start() {
	fmt.Println("Starting SSH server on " + srv.Serv.Addr)
	if err := srv.Serv.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

func (srv *CharmServ) Stop() {
	fmt.Println("Stopping SSH server")
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer func() { cancel() }()
	if err := srv.Serv.Shutdown(ctx); err != nil {
		log.Fatalln(err)
	}
}

func teaHandler(s ssh.Session) (tea.Model, []tea.ProgramOption) {
	_, _, active := s.Pty()
	if !active {
		return nil, nil
	}

	return InitialPage(), []tea.ProgramOption{tea.WithAltScreen()}
}
