package server

import (
	"context"
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/wish"
	bm "github.com/charmbracelet/wish/bubbletea"
	lm "github.com/charmbracelet/wish/logging"
	"github.com/gliderlabs/ssh"
)

type CharmServ struct {
	Serv *ssh.Server
}

func MakeCharmServ() (*CharmServ, error) {
	var err error
	self := &CharmServ{}
	self.Serv, err = wish.NewServer(
		wish.WithAddress("0.0.0.0:23234"),
		wish.WithMiddleware(
			bm.Middleware(teaHandler),
			lm.Middleware(),
		),
	)
	return self, err
}

func (self *CharmServ) Start() {
	fmt.Println("Starting SSH server on 0.0.0.0:23234")
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
