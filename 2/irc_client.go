package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"

	"github.com/fratzik/gos/2/processors"
	"github.com/mvdan/xurls"
)

// IrcClient - the IRC Client
type IrcClient struct {
	Bot  *Bot
	Conn net.Conn
}

// HandleCommunication - handle the communication
func (ic *IrcClient) HandleCommunication() {

	ic.Connect()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		reader := bufio.NewReader(ic.Conn)
		for {
			line, err := reader.ReadString('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("%v\n", err)
			}

			ic.HandleServerResponse(line)
		}
		wg.Done()
	}()

	ic.connectBot()
	wg.Wait()
}

// Connect the bot to the IRC
func (ic *IrcClient) Connect() {
	conn, err := net.Dial("tcp", ic.Bot.Server)
	if err != nil {
		log.Fatal("Unable to connect to IRC server ", err)
	}

	ic.Conn = conn
	log.Printf("Connected to IRC server %s (%s)\n", ic.Bot.Server, ic.Conn.RemoteAddr())
}

func (ic *IrcClient) connectBot() {
	pattern, patternExists := commandPatterns[CmdCONNECT]
	if patternExists {
		command := &Command{Name: CmdCONNECT, Pattern: pattern, Bot: ic.Bot}
		ic.SendCommand(command)
	}
}

// HandleServerResponse - parse and handle text lines, received from IRC
func (ic *IrcClient) HandleServerResponse(line string) {
	command, chunks := GetLineCommand(line)

	if command != nil {

		//additional checks
		if command.Name == "JOIN" {
			//check not to sent greeting message to yourself
			if strings.Contains(chunks[0], ic.Bot.Nick) {
				log.Println("It is not a nice thing to salute yourself.")
				return
			}
		} else if command.Name == CmdPRIVMSG {
			urls := xurls.Strict.FindAllString(line, -1)
			if len(urls) > 0 {
				pageTitle, err := processors.GetUrlTitle(urls[0])
				if err != nil {
					log.Println(err)
				} else {
					command.Name = CmdURLTitle
					command.Pattern = "PRIVMSG %v :Recognized a title: %s %s"
					command.AddAdditionalParam("title", pageTitle)
				}
			}
		}

		command.Bot = ic.Bot
		ic.SendCommand(command)
	}
}

func (ic *IrcClient) disconnect() {
	ic.Conn.Close()
}

// SendCommand - send a message described in the c command
func (ic *IrcClient) SendCommand(c *Command) {
	fmt.Fprint(ic.Conn, c)
}
