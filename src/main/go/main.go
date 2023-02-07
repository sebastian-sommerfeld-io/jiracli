package main

import (
	"log"

	"github.com/sebastian-sommerfeld-io/jiracli/commands/jiracli"
)

func init() {
	log.SetPrefix("[jiracli] ")
}

func main() {
	err := jiracli.Execute()

	if err != nil {
		log.Fatal(err)
	}
}
