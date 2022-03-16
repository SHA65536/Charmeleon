package main

import (
	"charmeleon/pokemon"
	"charmeleon/server"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
)

const data_export = `https://github.com/msikma/pokesprite/trunk/data/pokemon.json`
const sprite_export = `https://github.com/msikma/pokesprite/trunk/pokemon-gen8`
const download_cmd = "svn export " + data_export + " && svn export " + sprite_export
const clean_cmd = "sudo rm -r data/ pokemon-gen8/*.png"
const rename_cmd = `mv pokemon-gen8/ data/ && mv pokemon.json data/sprites.json`

func main() {
	startServer()
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
	fmt.Println("Complete! You should rebuild the server.")
}

func parse() (map[string]*pokemon.Pokemon, error) {
	var pokedex map[string]*pokemon.Pokemon
	jsonFile, _ := os.Open("data/sprites.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	err := json.Unmarshal(byteValue, &pokedex)
	return pokedex, err
}

func convertToXterm(pokes map[string]*pokemon.Pokemon) {
	for _, poke := range pokes {
		for _, path := range poke.Forms.Entries {
			png := strings.Replace(path.Path, "{name}", poke.Slug.Str, 1) + ".png"
			cow := png + ".cow"
			exec.Command("img2xterm", png, cow).Run()
		}
	}
}

func startServer() {
	serv, err := server.MakeCharmServ()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go serv.Start()

	<-done
	serv.Stop()
}
