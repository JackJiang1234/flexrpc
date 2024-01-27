package transport

import (
	"context"
	"xrpc/protocol"
)

type Sender interface {
	Send(ctx context.Context, m ...*protocol.Message) error
	SendAsync(m ...*protocol.Message) error
	Close() error
}
