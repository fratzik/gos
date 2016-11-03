package main

import "fmt"

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
		retStr = fmt.Sprintf(com.Pattern, com.Bot.Channel, crlf)
	case CmdPING:
		retStr = fmt.Sprintf(com.Pattern, com.Bot.Nick, crlf)
	case CmdCONNECT:
		retStr = fmt.Sprintf(com.Pattern, com.Bot.Nick, crlf, com.Bot.Nick, crlf)
	}

	return retStr
}

func isKnownCommand(commandKey string) bool {
	return commandKey == CmdPING || commandKey == CmdCONNECT || commandKey == Cmd376 || commandKey == CmdJOIN || commandKey == CmdPRIVMSG
}
