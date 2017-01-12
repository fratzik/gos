package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/fratzik/gos/2/processors"
	"github.com/mvdan/xurls"
)

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
		retStr = fmt.Sprintf(com.Pattern, com.Bot.Channel, com.GetAdditionalParam("urlsLen"), com.GetAdditionalParam("title"), crlf)
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

func (com *Command) AddAdditionalParam(name string, value interface{}) {
	com.Params[name] = value
}

func (com *Command) GetAdditionalParam(name string) interface{} {
	return com.Params[name]
}

var titles []string

func ExecExtraProcess(command *Command, line string, chunks []string) bool {

	if command.Name == CmdJOIN {
		//check not to sent greeting message to yourself
		if strings.Contains(chunks[0], command.Bot.Nick) {
			// log.Println("It is not a nice thing to salute yourself.")
			return false
		}
	}

	if command.Name == CmdPRIVMSG {
		urls := xurls.Strict.FindAllString(line, -1)

		titles = []string{}

		for _, v := range urls {
			pageTitle, err := processors.GetUrlTitle(v)
			if err != nil {
				log.Printf("%v\n", err)
			} else {
				titles = append(titles, pageTitle)
				// fmt.Println("Title: ", pageTitle)
			}

		}
		// fmt.Printf("%v", urls)
		if len(titles) > 0 {
			command.Name = CmdURLTitle
			command.Pattern = "PRIVMSG %v :Recognized %d title(s): %s %s"
			command.AddAdditionalParam("urlsLen", len(titles))
			command.AddAdditionalParam("title", strings.Join(titles, " | "))
		}
	}

	return true
}
