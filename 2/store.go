package main

import (
	"fmt"
	"log"
)

type Store struct {
}

func (store Store) dispatch(command string, bot *Bot, chunks []string) {
	// params := make([]string, 2)
	params := []interface{}{}
	var pattern string

	switch command {
	case "PING":
		log.Println("Send PONG to the server")
		// fmt.Fprintf(bot.Conn, "PONG :%s%s", chunks[0], crlf)
		pattern = "PONG :%s%s"
		params = append(params, chunks[0])
	case "376":
		log.Println("Sending join command")
		// fmt.Fprintf(bot.Conn, "JOIN %v%s", bot.Channel, crlf)
		pattern = "JOIN %v%s"
		params = append(params, bot.Channel)
	case "JOIN":
		log.Println("Sending welcome command")
		// log.Printf("PRIVMSG %v :Hi!%s", bot.Channel, crlf)
		// fmt.Fprintf(bot.Conn, "PRIVMSG %v :Hi!%s", bot.Channel, crlf)
		pattern = "PRIVMSG %v :Hi!%s"
		params = append(params, bot.Channel)
	case "PRIVMSG":
		log.Printf("Private message received: %v", chunks[2])
		// fmt.Fprintf(bot.Conn, "PRIVMSG %v :Response to message.%s", bot.Channel, crlf)
		pattern = "PRIVMSG %v :Response to message.%s"
		params = append(params, bot.Channel)
	case "CONNECT":

		//this should be removed in a way
		fmt.Fprintf(bot.Conn, "NICK %v\n", bot.Nick)
		// fmt.Fprintf(bot.Conn, "USER %v 8 * :This would be a description%s", bot.Nick, crlf)
		pattern = "USER %v 8 * :This would be a description%s"
		params = append(params, bot.Nick)
	}

	// params = append(params, crlf)
	// values := []interface{}{}
	params = append(params, crlf)
	commandString := fmt.Sprintf(pattern, params...)
	log.Print(commandString)
	fmt.Fprint(bot.Conn, commandString)

}
