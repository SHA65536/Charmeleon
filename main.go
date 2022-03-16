package main

import (
	"charmeleon/pokemon"
	"charmeleon/server"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/akamensky/argparse"
)

const data_export = `https://github.com/msikma/pokesprite/trunk/data/pokemon.json`
const sprite_export = `https://github.com/msikma/pokesprite/trunk/pokemon-gen8`
const download_cmd = "svn export " + data_export + " && svn export " + sprite_export
const clean_cmd = "sudo rm -r data/ pokemon-gen8/*.png"
const rename_cmd = `mv pokemon-gen8/ data/ && mv pokemon.json data/sprites.json`

func main() {
	parser := argparse.NewParser("charmeleon", "Pokedex over SSH!")

	startCmd := parser.NewCommand("start", "Start the SSH server")
	host := startCmd.String("a", "address", &argparse.Options{Required: false, Help: "Address to listen on", Default: "0.0.0.0"})
	port := startCmd.Int("p", "port", &argparse.Options{Required: false, Help: "Port to listen on", Default: 23234})

	buildCmd := parser.NewCommand("build", "Download images and remake cow files")

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		return
	}
	if startCmd.Happened() {
		startServer(*host, *port)
	} else if buildCmd.Happened() {
		redownload()
		convertToXterm()
	} else {
		err := fmt.Errorf("bad arguments, please check usage")
		fmt.Print(parser.Usage(err))
	}
}

//Warning: SVN needs to installed for this to work
//Only tested on Ubuntu
func redownload() {
	fmt.Println("Downloading... ", download_cmd)
	cmd := exec.Command("sh", "-c", download_cmd)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Cleaning... ", clean_cmd)
	cmd = exec.Command("sudo", "sh", "-c", clean_cmd)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Renaming... ", rename_cmd)
	cmd = exec.Command("sh", "-c", rename_cmd)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

//Warning: img2xterm is required for this to work
//Only tested on Ubuntu
func convertToXterm() {
	var pokedex map[string]*pokemon.Pokemon
	file, err := os.ReadFile("data/sprites.json")
	err = json.Unmarshal(file, &pokedex)
	if err != nil {
		log.Fatal(err)
	}
	for _, poke := range pokedex {
		for _, form := range poke.Forms {
			exec.Command("img2xterm", form.Png, form.Cow).Run()
		}
	}
}

func startServer(host string, port int) {
	serv, err := server.MakeCharmServ(fmt.Sprintf("%s:%v", host, port))
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go serv.Start()

	<-done
	serv.Stop()
}
