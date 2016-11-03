package main

// Store - this is created as struct for the future state and functionality
type Store struct {
}

// Command object
// type Command struct {
// 	Pattern string
// 	Name    string
// 	Bot     *Bot
// }

// func (com Command) String() string {
// 	var retStr string
// 	switch com.Name {
// 	default: //376, JOIN, PRIVMSG
// 		retStr = fmt.Sprintf(com.Pattern, bot.Channel, crlf)
// 	case CmdPING:
// 		retStr = fmt.Sprintf(com.Pattern, com.Bot.Nick, crlf)
// 	case CmdCONNECT:
// 		retStr = fmt.Sprintf(com.Pattern, com.Bot.Nick, crlf, com.Bot.Nick, crlf)
// 	}

// 	return retStr
// }

// func isKnownCommand(commandKey string) bool {
// 	return commandKey == CmdPING || commandKey == CmdCONNECT || commandKey == Cmd376 || commandKey == CmdJOIN || commandKey == CmdPRIVMSG
// }

// func (store Store) process(chunks []string) {

// 	commandKey := chunks[1]
// 	if isKnownCommand(strings.TrimSpace(chunks[0])) {
// 		commandKey = strings.TrimSpace(chunks[0])
// 	}
// 	pattern, patternExists := commandPatterns[commandKey]
// 	if patternExists {
// 		command := Command{Name: commandKey, Pattern: pattern[0], Bot: bot}
// 		fmt.Fprint(bot.Conn, command)
// 	}
// }
