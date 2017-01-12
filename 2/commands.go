package main

var commandPatterns map[string]string

//command idents
const (
	CmdPING     = "PING"
	Cmd376      = "376"
	CmdJOIN     = "JOIN"
	CmdPRIVMSG  = "PRIVMSG"
	CmdCONNECT  = "CONNECT"
	CmdURLTitle = "SEND_URL_TITLE"
)

//messages
const (
	MsgJOIN        = "Hello there!"
	MsgMESSAGERESP = "This is a response to what you wrote."
	MsgTextCONNECT = "Here is a text string"
)

func initCommandsMap() {
	commandPatterns = map[string]string{
		CmdPING:    "PONG :%s%s",
		Cmd376:     "JOIN %v%s",
		CmdJOIN:    "PRIVMSG %v :" + MsgJOIN + "%s",
		CmdPRIVMSG: "PRIVMSG %v :" + MsgMESSAGERESP + "%s",
		CmdCONNECT: "NICK %v%sUSER %v 8 * :" + MsgTextCONNECT + "%s"}
}
