package main

import (
	"log"
	"net"
	"strings"

	"github.com/Xe/cbaas/lib"
	"github.com/apex/invoke"
	irc "gopkg.in/irc.v1"
)

func main() {
	conn, err := net.Dial("tcp", "irc.ponychat.net:6667")
	if err != nil {
		log.Fatal(err)
	}

	c := irc.NewClient(conn, irc.ClientConfig{
		Nick:    "cbaas",
		User:    "cbaas",
		Name:    "Chatbot as a service",
		Handler: irc.HandlerFunc(handleLine),
	})

	c.Run()
}

func handleLine(c *irc.Client, m *irc.Message) {
	if m.Command != "PRIVMSG" {
		switch m.Command {
		case "001":
			c.Writef("JOIN #niichan")
		}
		return
	}

	to := ""

	if c.FromChannel(m) {
		to = m.Params[0]
	} else {
		to = m.Prefix.Name
	}

	lm := &lib.Message{
		Protocol:   "irc",
		Sender:     m.Prefix.Name,
		To:         to,
		BodyString: m.Params[1],
	}
	lo := []*lib.Message{}

	if m.Params[1][0] != ';' {
		return
	}

	splitLine := strings.Split(m.Params[1], " ")
	command := strings.ToLower(splitLine[0][1:])

	err := invoke.Sync("cbaas_"+command, lm, &lo)
	if err != nil {
		c.Writef("PRIVMSG %s :error: %v", to, err)
		log.Println(err)
		return
	}

	for _, line := range lo {
		log.Printf("%s: %s", to, line.IRC())
		c.Writef("PRIVMSG %s :%s", to, line.IRC())
	}
}
