package main

var commandPatterns map[string][]string

func initCommandPatterns() {
	commandPatterns = map[string][]string{
		"PING":    {"PONG :%s%s"},
		"376":     {"JOIN %v%s"},
		"JOIN":    {"PRIVMSG %v :Hi!%s"},
		"PRIVMSG": {"PRIVMSG %v :Response to message.%s"},
		"CONNECT": {"NICK %v%s", "USER %v 8 * :Here is a text string%s"},
	}
}
