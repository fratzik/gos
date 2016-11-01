package main

import "fmt"

// Store - this is created as struct for the future state and functionality
type Store struct {
}

// Command object
type Command struct {
	Pattern string
	Name    string
	Bot     *Bot
}

func (com Command) String() string {
	var retStr string
	switch com.Name {
	default: //376, JOIN, PRIVMSG
		retStr = fmt.Sprintf(com.Pattern, bot.Channel, crlf)
	case "PING":
		retStr = fmt.Sprintf(com.Pattern, com.Bot.Nick, crlf)
	case "CONNECT":
		retStr = fmt.Sprintf(com.Pattern, com.Bot.Nick, crlf, com.Bot.Nick, crlf)
	}

	return retStr
}

func (store Store) process(chunks []string) {

	commandKey := chunks[1]
	pattern, ok := commandPatterns[commandKey]
	if ok {
		command := Command{Name: commandKey, Pattern: pattern[0], Bot: bot}
		fmt.Fprint(bot.Conn, command)
	}

	// switch commandKey {
	// case "PING":
	// 	commands = append(commands, commandPatterns[commandKey], []interface{}{chunks[0]}})
	// case "376":
	// 	commands = append(commands, commandPatterns[commandKey], []interface{}{bot.Channel}})
	// case "JOIN":
	// 	commands = append(commands, Command{commandPatterns[commandKey], []interface{}{bot.Channel}})
	// case "PRIVMSG":
	// 	commands = append(commands, Command{commandPatterns[commandKey], []interface{}{bot.Channel}})
	// case "CONNECT":
	// 	commands = append(commands, Command{commandPatterns[commandKey], []interface{}{bot.Nick}})
	// 	commands = append(commands, Command{commandPatterns[commandKey], []interface{}{bot.Nick}})
	// }

	// if len(commands) > 0 {

	// 	for _, command := range commands {
	// 		params := append(command.Params, crlf)
	// 		commandString := fmt.Sprintf(command.Pattern, params...)
	// 		fmt.Fprint(bot.Conn, commandString)
	// 	}
	// }

}
