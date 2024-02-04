package transport

type ServerTransport interface {
	Open() error
	Close() error
}