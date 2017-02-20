// Package lib is common types and functions for cbaas go functions.
package lib

type Message struct {
	Protocol   string    `json:"protocol"`
	Sender     string    `json:"sender,omitempty"`
	To         string    `json:"to"`
	Body       []MsgFrag `json:"body"`
	BodyString string    `json:"bodyString"`
}

func (m *Message) IRC() string {
	if len(m.Body) == 0 {
		return m.BodyString
	}

	result := ""

	for _, frag := range m.Body {
		result += frag.IRC()
	}
	return result
}

type MsgFrag struct {
	Body    string `json:"body"`
	Mention string `json:"mention"`
	Emoji   *Emoji `json:"emoji"`
}

func (m MsgFrag) IRC() string {
	if m.Body != "" {
		return m.Body
	}

	if m.Mention != "" {
		return m.Mention
	}

	return ""
}

type Emoji struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}
