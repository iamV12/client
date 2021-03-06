// Auto-generated by avdl-compiler v1.3.7 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/chat1/gregor.avdl

package chat1

import (
	rpc "github.com/keybase/go-framed-msgpack-rpc"
)

type GenericPayload struct {
	Action string `codec:"Action" json:"Action"`
}

type NewMessagePayload struct {
	Action  string         `codec:"Action" json:"Action"`
	ConvID  ConversationID `codec:"convID" json:"convID"`
	Message MessageBoxed   `codec:"message" json:"message"`
}

type GregorInterface interface {
}

func GregorProtocol(i GregorInterface) rpc.Protocol {
	return rpc.Protocol{
		Name:    "chat.1.gregor",
		Methods: map[string]rpc.ServeHandlerDescription{},
	}
}

type GregorClient struct {
	Cli rpc.GenericClient
}
