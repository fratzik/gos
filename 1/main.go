package main

import (
	"log"
	"os"

	"github.com/fratzik/gos/1/processors"
)

func init() {
	log.SetFlags(0)
	log.SetPrefix("» ")
}

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		log.Fatal("Please provide the input param")
	}

	processors.ProcessArchive(args[0])
}
