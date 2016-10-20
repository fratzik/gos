package main

import (
	"fmt"
	"log"
)

// Store - this is created as struct for the future state and functionality
type Store struct {
}

// Command object
type Command struct {
	Pattern string
	Params  []interface{}
}

func (store Store) dispatch(command string, bot *Bot, chunks []string) {

	// params := []interface{}{}
	// pattern := ""
	commands := []Command{}

	switch command {
	case "PING":
		log.Println("Send PONG to the server")

		commands = append(commands, Command{"PONG :%s%s", []interface{}{chunks[0]}})
	case "376":
		log.Println("Sending join command")
		commands = append(commands, Command{"JOIN %v%s", []interface{}{bot.Channel}})
	case "JOIN":
		log.Println("Sending welcome command")
		commands = append(commands, Command{"PRIVMSG %v :Hi!%s", []interface{}{bot.Channel}})
	case "PRIVMSG":
		log.Printf("Private message received: %v", chunks[2])
		commands = append(commands, Command{"PRIVMSG %v :Response to message.%s", []interface{}{bot.Channel}})
	case "CONNECT":

		commands = append(commands, Command{"NICK %v%s", []interface{}{bot.Nick}})
		commands = append(commands, Command{"USER %v 8 * :This would be a description%s", []interface{}{bot.Nick}})
	}

	if len(commands) > 0 {

		for _, command := range commands {
			// log.Printf("value=%v", command)
			params := append(command.Params, crlf)
			// log.Println(params)
			commandString := fmt.Sprintf(command.Pattern, params...)
			// log.Print(commandString)
			fmt.Fprint(bot.Conn, commandString)
		}
		// params = append(params, crlf)
		// commandString := fmt.Sprintf(pattern, params...)
		// log.Print(commandString)
		// fmt.Fprint(bot.Conn, commandString)
	}

}
