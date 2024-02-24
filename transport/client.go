package transport

import "xrpc/transport/message"

type Client interface {
	sync(m message.Message) error
	async(m message.Message) (<-chan *message.Message, error)
	oneway(m message.Message) error
}