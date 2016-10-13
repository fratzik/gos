package main

import (
	"log"
	"net"
)

// NewBot - Create a new bot
func NewBot() *Bot {
	return &Bot{Server: server,
		Channel: "#" + channel,
		Nick:    botname}
}

// Connect the bot to the IRC
func (bot *Bot) Connect() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", bot.Server)
	if err != nil {
		log.Fatal("Unable to connect to IRC server ", err)
	}
	bot.Conn = conn
	log.Printf("Connected to IRC server %s (%s)\n", bot.Server, bot.Conn.RemoteAddr())
	return bot.Conn, nil
}
