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
var bot *Bot
var store Store

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

	store = Store{}
}

func main() {
	bot = NewBot()
	conn, _ := bot.Connect()
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		reader := bufio.NewReader(conn)
		for {
			line, err := reader.ReadString('\n')

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("%v\n", err)
			}

			handleResponseLine(line, conn)
		}

		wg.Done()
	}()

	fakeChunks := make([]string, 2)
	store.dispatch("CONNECT", bot, fakeChunks)

	wg.Wait()
}

func handleResponseLine(line string, conn net.Conn) {
	line = strings.TrimSuffix(line, crlf)
	chunks := strings.SplitN(line, " ", 5)

	if len(chunks) < 1 {
		return
	}

	fmt.Println(line)
	fmt.Println(chunks)
	store.dispatch(chunks[1], bot, chunks)
}
