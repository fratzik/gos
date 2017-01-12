package main

import "fmt"

// Command object
type Command struct {
	Pattern string
	Name    string
	Bot     *Bot
	Params  map[string]interface{}
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
	case CmdURLTitle:
		retStr = fmt.Sprintf(com.Pattern, com.Bot.Channel, com.GetAdditionalParam("title"), crlf)
	}

	return retStr
}

func NewCommand(commandKey string, pattern string) *Command {
	command := &Command{Name: commandKey, Pattern: pattern}
	command.Params = make(map[string]interface{})

	return command
}

func isKnownCommand(commandKey string) bool {
	return commandKey == CmdPING || commandKey == CmdCONNECT || commandKey == Cmd376 || commandKey == CmdJOIN || commandKey == CmdPRIVMSG
}

func (com *Command) AddAdditionalParam(name string, value string) {
	com.Params[name] = value
}

func (com *Command) GetAdditionalParam(name string) interface{} {
	return com.Params[name]
}
