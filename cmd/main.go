package main

import (
	"charmeleon/pokemon"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const data_export = `https://github.com/msikma/pokesprite/trunk/data/pokemon.json`
const sprite_export = `https://github.com/msikma/pokesprite/trunk/pokemon-gen8`
const download_cmd = "svn export " + data_export + " && svn export " + sprite_export
const clean_cmd = "sudo rm -r data/ pokemon-gen8/*.png"
const rename_cmd = `mv pokemon-gen8/ data/ && mv pokemon.json data/sprites.json`

func main() {
	redownload()
	parse()
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

func parse() {
	var pokedex map[string]*pokemon.Pokemon
	var formatted []byte
	var err error
	jsonFile, _ := os.Open("data/sprites.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &pokedex)
	for _, v := range pokedex {
		if formatted, err = json.MarshalIndent(v, "", "   "); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(string(formatted))
		}
		fmt.Scanln()
	}
}
