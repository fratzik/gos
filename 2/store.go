package main

import (
	"fmt"
	"log"
)

// Store - this is created as struct for the future state and functionality
type Store struct {
}

func (store Store) dispatch(command string, bot *Bot, chunks []string) {
	params := []interface{}{}
	pattern := ""

	switch command {
	case "PING":
		log.Println("Send PONG to the server")
		pattern = "PONG :%s%s"
		params = append(params, chunks[0])
	case "376":
		log.Println("Sending join command")
		pattern = "JOIN %v%s"
		params = append(params, bot.Channel)
	case "JOIN":
		log.Println("Sending welcome command")
		pattern = "PRIVMSG %v :Hi!%s"
		params = append(params, bot.Channel)
	case "PRIVMSG":
		log.Printf("Private message received: %v", chunks[2])
		pattern = "PRIVMSG %v :Response to message.%s"
		params = append(params, bot.Channel)
	case "CONNECT":

		//this should be removed in a way
		fmt.Fprintf(bot.Conn, "NICK %v\n", bot.Nick)
		pattern = "USER %v 8 * :This would be a description%s"
		params = append(params, bot.Nick)
	}

	if pattern != "" {
		params = append(params, crlf)
		commandString := fmt.Sprintf(pattern, params...)
		log.Print(commandString)
		fmt.Fprint(bot.Conn, commandString)
	}

}
