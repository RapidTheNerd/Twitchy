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
	conn.Write([]byte("PASS " + "oauth:token" + "\r\n"))
	conn.Write([]byte("NICK " + "username" + "\r\n"))
	conn.Write([]byte("JOIN " + "#channel" + "\r\n"))
	defer conn.Close()
	}
