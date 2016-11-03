package main

import (
	"log"
	"strings"
)

// CreateIrcClient - Create and return a new instance of &IrcClient
func CreateIrcClient(bot *Bot) *IrcClient {
	return &IrcClient{Bot: bot}
}

// CreateBot - Create and return a new instance of &Bot
func CreateBot(server, channel, botname string) *Bot {
	return &Bot{Server: server, Channel: channel, Nick: botname}
}

func GetLineCommand(line string) (*Command, []string) {
	line = strings.TrimSuffix(line, crlf)
	chunks := strings.SplitN(line, " ", 5)

	if len(chunks) < 1 {
		return nil, nil
	}

	log.Println(line)
	commandKey := chunks[1]

	if isKnownCommand(strings.TrimSpace(chunks[0])) {
		commandKey = strings.TrimSpace(chunks[0])
	}

	pattern, patternExists := commandPatterns[commandKey]

	if !patternExists {
		return nil, nil
	}

	return &Command{Name: commandKey, Pattern: pattern}, chunks
}
