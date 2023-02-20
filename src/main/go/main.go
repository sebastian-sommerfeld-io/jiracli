package main

import (
	"log"

	"github.com/sebastian-sommerfeld-io/jiracli/commands"
)

func init() {
	log.SetPrefix("[jiracli] ")
}

func main() {
	err := commands.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
