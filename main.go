package twitchy

import (
	"net"
	"net/textproto"
	"bufio"
	"strings"
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

	tp := textproto.NewReader(bufio.NewReader(conn))
	for {
		msg, err := tp.ReadLine()
		if err != nil {
			panic(err)
		}
		msgParts := strings.Split(msg, " ")

		if msgParts[0] == "PING" {
			conn.Write([]byte("PONG " + msgParts[1]))
			continue
		}
		if msgParts[1] == "PRIVMSG" {
			conn.Write([]byte("PRIVMSG " + msgParts[2] + " " + msgParts[3] + "\r\n"))
		}
	}
}