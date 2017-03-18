package main

import (
	"fmt"
	"os"

	"github.com/funny/gobuf/parser"

	"encoding/json"
	"flag"
	"log"
)

func main() {
	flag.Parse()

	file, isGoGen := goGenInfo()
	if isGoGen {
		genFile(file)
	} else {
		for _, file := range flag.Args() {
			genFile(file)
		}
	}
}

func genFile(name string) {
	doc, err := parser.Parse(name)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.MarshalIndent(doc, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", data)
}

func goGenInfo() (string, bool) {
	_, hasGOARCH := os.LookupEnv("GOOS")
	_, hasGOOS := os.LookupEnv("GOARCH")
	file, hasGOFILE := os.LookupEnv("GOFILE")
	_, hasGOLINE := os.LookupEnv("GOLINE")
	_, hasGOPACKAGE := os.LookupEnv("GOPACKAGE")
	dollar, _ := os.LookupEnv("DOLLAR")
	return file, hasGOARCH &&
		hasGOOS &&
		hasGOFILE &&
		hasGOLINE &&
		hasGOPACKAGE &&
		dollar == "$"
}
