package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Xe/cbaas/lib"
	"github.com/apex/go-apex"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		m := &lib.Message{}

		if err := json.Unmarshal(event, m); err != nil {
			return nil, err
		}

		hipster, err := getHipsterText()
		if err != nil {
			return nil, err
		}

		resp := &lib.Message{
			To: m.To,
			Body: []lib.MsgFrag{
				{Body: hipster},
			},
		}

		return []*lib.Message{resp}, nil
	})
}

func getHipsterText() (string, error) {
	resp, err := http.Get("http://hipsterjesus.com/api/?type=hipster-centric&html=false&paras=1")
	if err != nil {
		return "", err
	}

	textStruct := &struct {
		Text string `json:"text"`
	}{}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	json.Unmarshal(body, textStruct)

	text := strings.Split(textStruct.Text, ". ")[0]

	return text, nil
}
