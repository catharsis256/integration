package server

import (
	"context"
	"fmt"
	"gitlab.smartblocklab.com/integration/library/proto"
)

type Greeter struct {
	Exclaim bool
}

func (g *Greeter) Greet(ctx context.Context, r *proto.GreeterRequest) (*proto.GreeterResponse, error) {
	msg := fmt.Sprintf("%s %s", r.GetGreeting(), r.GetName())

	if g.Exclaim {
		msg += "!"
	} else {
		msg += "."
	}

	return &proto.GreeterResponse{Response: msg}, nil
}
