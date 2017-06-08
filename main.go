package twitchy

import (
	"net",
	"net/textproto",
	"strings",
	"bufio"
)

func main() {

	conn, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
		if err != nil {
			panic(err);
		}
}
