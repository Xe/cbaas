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

		st, err := getMessage()
		if err != nil {
			return nil, err
		}

		resp := &lib.Message{
			To: m.To,
			Body: []lib.MsgFrag{
				{Body: fmt.Sprintf("status: %s, message: %s", st.Status, st.Body)},
			},
		}

		return []*lib.Message{resp}, nil
	})
}
