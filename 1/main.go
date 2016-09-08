package main

import (
	"letsgo/1/processors"
	"log"
	"os"
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
