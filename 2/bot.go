package main

import (
	"fmt"
	"log"
	"net"
)

// NewBot - Create a new bot
func NewBot() *Bot {
	return &Bot{Server: server,
		Channel: "#" + channel,
		Nick:    botname,
		Conn:    nil}
}

// Connect the bot to the IRC
func (bot *Bot) Connect() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", bot.Server)
	if err != nil {
		log.Fatal("unable to connect to IRC server ", err)
	}
	bot.Conn = conn
	log.Printf("Connected to IRC server %s (%s)\n", bot.Server, bot.Conn.RemoteAddr())
	return bot.Conn, nil
}

func (c *Bot) register() {
	log.Println("Sending connection commands...")
	fmt.Fprintf(c.Conn, "NICK %v\n", c.Nick)
	fmt.Fprintf(c.Conn, "USER %v 8 * :Greeting Bot Written in GoLang\n", c.Nick)
}
