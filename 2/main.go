package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/textproto"
)

// Bot struct
type Bot struct {
	Server  string
	Port    string
	Nick    string
	User    string
	Channel string
	Conn    net.Conn
}

// NewBot - Create a new bot
func NewBot() *Bot {
	return &Bot{Server: "bucharest.ro.eu.undernet.org",
		Port:    "6667",
		Nick:    "golangbot2217",
		Channel: "#go-test-bot",
		Conn:    nil}
}

// Connect the bot to the IRC
func (bot *Bot) Connect() (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", bot.Server+":"+bot.Port)
	if err != nil {
		log.Fatal("unable to connect to IRC server ", err)
	}
	bot.Conn = conn
	log.Printf("Connected to IRC server %s (%s)\n", bot.Server, bot.Conn.RemoteAddr())
	return bot.Conn, nil
}

func main() {
	bot := NewBot()
	conn, _ := bot.Connect()
	fmt.Fprintf(conn, "USER %s 8 * :%s\r\n", bot.Nick, bot.Nick)
	fmt.Fprintf(conn, "NICK %s\r\n", bot.Nick)
	fmt.Fprintf(conn, "JOIN %s\r\n", bot.Channel)
	defer conn.Close()

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)
	for {
		line, err := tp.ReadLine()
		if err != nil {
			break // break loop on errors
		}
		fmt.Printf("%s\n", line)
	}
}
