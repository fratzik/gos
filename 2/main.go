package main

import (
	"flag"
	"log"
)

var crlf = "\r\n"
var server, channel, botname string

func init() {

	log.SetFlags(0)
	log.SetPrefix("» ")

	flag.StringVar(&server, "server", "chat.freenode.net:6667", "The server to connect too")
	flag.StringVar(&channel, "channel", "#go-test-bot-2", "The channel connect too")
	flag.StringVar(&botname, "test-bot", "gobotnm", "The name of the boot")
	flag.Parse()

	initCommandsMap()

}

func main() {
	bot := CreateBot(server, channel, botname)
	ic := CreateIrcClient(bot)
	ic.HandleCommunication()
}
