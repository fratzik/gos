package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

const (
	StatusDisconnected = 10
	StatusConnected    = 20
	StatusLoggedIn     = 30
)

type Bot struct {
	Server  string
	Channel string
	Nick    string
	Status  int
	Conn    net.Conn
}

func NewBot(server string, channel string, botname string) *Bot {
	return &Bot{
		Server:  server,
		Channel: channel,
		Nick:    botname,
		Status:  StatusDisconnected}
}

func (bot *Bot) Connect() {
	conn, err := net.Dial("tcp", bot.Server)
	if err != nil {
		log.Fatal("Unable to connect to IRC server ", err)
	}

	bot.Conn = conn
	log.Printf("Connected to IRC server %s (%s)\n", bot.Server, bot.Conn.RemoteAddr())
}

func (bot *Bot) Disconnect() {
	bot.Conn.Close()
}

// func (bot *Bot) GetCommandsSet() CommandsSet {
// 	return Commands[bot.Status]
// }

func (bot *Bot) Run() {
	//conect the bot
	bot.Connect()
	// bot.
	bot.ExecCommand(RegisterUser(bot.Nick, "A name here"))

	reader := bufio.NewReader(bot.Conn)
	for true {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v\n", err)
		}

		fmt.Println(line)

		command, err := ParseCommand(line)
		if err == nil {
			fmt.Printf("command here: %#v", command)
			// switch command.Command {
			// case replies.CmdPing:
			// 	bot.Write(commands.Pong(command.Params))
			// default:
			// 	bot.State.HandleCommand(command, bot)
			// }
		} else {
			log.Println(err)
		}
	}

	// cmds := GetCommands(bot.Status)

	// var wg sync.WaitGroup
	// wg.Add(1)

	// go func() {
	// 	reader := bufio.NewReader(bot.Conn)
	// 	for {
	// 		line, err := reader.ReadString('\n')
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		if err != nil {
	// 			log.Fatalf("%v\n", err)
	// 		}

	// 		fmt.Printf("%v\n", line)
	// 		if err == nil {
	// 			switch command.Command {
	// 			case replies.CmdPing:
	// 				bot.Write(commands.Pong(command.Params))
	// 			default:
	// 				bot.State.HandleCommand(command, bot)
	// 			}
	// 		}
	// 		bot.HandleServerResponse(line)
	// 	}
	// 	wg.Done()
	// }()

	// for {

	// }
	// fmt.Printf("%#v\n", cmds)
}

func (bot *Bot) HandleServerResponse(line string) {
	fmt.Printf("handle this line: ", line)
}

func (bot *Bot) ExecCommand(command string) {
	fmt.Fprint(bot.Conn, command)
}
