package main

import (
	"log"
	"net"
	"strings"

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

	lm := &message{
		Protocol:   "irc",
		Sender:     m.Prefix.Name,
		To:         to,
		BodyString: m.Params[1],
	}
	lo := []*message{}

	if m.Params[1][0] != ';' {
		return
	}

	command := strings.ToLower(m.Params[1][1:])

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

type message struct {
	Protocol   string    `json:"protocol"`
	Sender     string    `json:"sender,omitempty"`
	To         string    `json:"to"`
	Body       []msgFrag `json:"body"`
	BodyString string    `json:"bodyString"`
}

func (m *message) IRC() string {
	if len(m.Body) == 0 {
		return m.BodyString
	}

	result := ""

	for _, frag := range m.Body {
		result += frag.IRC()
	}
	return result
}

type msgFrag struct {
	Body    string `json:"body"`
	Mention string `json:"mention"`
	Emoji   *emoji `json:"emoji"`
}

func (m msgFrag) IRC() string {
	if m.Body != "" {
		return m.Body
	}

	if m.Mention != "" {
		return m.Mention
	}

	return ""
}

type emoji struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
