package transport

import (
	"context"
	"xrpc/protocol"
)

type Receiver interface {
	Receive(ctx context.Context) (*protocol.Message, error);
	Close() error;
}