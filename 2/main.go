package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	_ "net/textproto"
	"strings"
	"sync"
)

var crlf = "\r\n"
var server, channel, botname string

// Bot struct
type Bot struct {
	Server  string
	Channel string
	Nick    string
	Conn    net.Conn
}

func init() {

	log.SetFlags(0)
	log.SetPrefix("» ")

	flag.StringVar(&server, "server", "chat.freenode.net:6667", "The server to connect too")
	flag.StringVar(&channel, "channel", "go-test-bot", "The channel connect too")
	flag.StringVar(&botname, "test-bot", "gobotnm", "The name of the boot")
	flag.Parse()

	if server == "" || channel == "" {
		log.Fatalln("Please set the server and channel params")
	}
}

func isPongLine(line string) bool {
	return strings.Contains(line, "PING")
}

func main() {
	log.Println("Incepe aplicatia de chat")
	bot := NewBot()
	conn, _ := bot.Connect()
	fmt.Fprintf(conn, "JOIN %s"+crlf, bot.Channel)
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	//tp := textproto.NewReader(reader)
	go func() {
		reader := bufio.NewReader(conn)
		for {
			line, err := reader.ReadString('\n')

			if err == io.EOF {
				break // break loop on errors
			}

			if err != nil {
				//could be this really fatal?
				log.Fatalf("%v\n", err)
			}

			handleResponseLine(line, conn)
		}

		wg.Done()
	}()

	bot.register()

	wg.Wait()
}

func handleResponseLine(line string, conn net.Conn) {
	line = strings.TrimSuffix(line, crlf)
	log.Println(line)
	params := strings.Split(line, ":")
	if len(params) > 2 {
		//log.Println(params[2])
	}

	if strings.Contains(line, "PING") {
		log.Println("Send PONG to the server")
		fmt.Fprintf(conn, "PONG :%s%s", params[1], crlf)
	}

}
