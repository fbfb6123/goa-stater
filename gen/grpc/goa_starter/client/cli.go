// Code generated by goa v3.4.3, DO NOT EDIT.
//
// goa_starter gRPC client CLI support package
//
// Command:
// $ goa gen goa_starter/design

package client

import (
	"encoding/json"
	"fmt"
	goastarter "goa_starter/gen/goa_starter"
	goa_starterpb "goa_starter/gen/grpc/goa_starter/pb"
)

// BuildAddPayload builds the payload for the goa_starter add endpoint from CLI
// flags.
func BuildAddPayload(goaStarterAddMessage string) (*goastarter.AddPayload, error) {
	var err error
	var message goa_starterpb.AddRequest
	{
		if goaStarterAddMessage != "" {
			err = json.Unmarshal([]byte(goaStarterAddMessage), &message)
			if err != nil {
				return nil, fmt.Errorf("invalid JSON for message, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"a\": 5121140462866214315,\n      \"b\": 2783468530862518908\n   }'")
			}
		}
	}
	v := &goastarter.AddPayload{
		A: int(message.A),
		B: int(message.B),
	}

	return v, nil
}
