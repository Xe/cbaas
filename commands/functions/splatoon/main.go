package main

import (
	"encoding/json"
	"fmt"

	"github.com/Xe/cbaas/lib"
	"github.com/apex/go-apex"
)

func main() {
	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		m := &lib.Message{}

		if err := json.Unmarshal(event, m); err != nil {
			return nil, err
		}

		sd, err := Lookup()
		if err != nil {
			return nil, err
		}

		data := sd[0]
		stage1 := data.Stages[0]
		stage2 := data.Stages[1]

		resp := &lib.Message{
			To: m.To,
			Body: []lib.MsgFrag{
				{Body: fmt.Sprintf("Splatoon stages: %s and %s", englishIfy(stage1), englishIfy(stage2))},
			},
		}

		return []*lib.Message{resp}, nil
	})
}
