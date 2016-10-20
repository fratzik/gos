package main

// Store - this is created as struct for the future state and functionality
type Store struct {
}

// Command object
type Command struct {
	Pattern string
	Params  []interface{}
}

func (store Store) process(chunks []string) {

	commandKey := chunks[1]
	// commands := []Command

	// fmt.Println("-----------")
	// fmt.Println(commandKey)
	// fmt.Printf("%v", commands)
	// fmt.Println("-----------")

	// commands := []Command{}

	for k, patterns := range commandPatterns {
		if commandKey == k {
			// params = getParams(k)

			if k == "CONNECT" {

			}
		}
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
