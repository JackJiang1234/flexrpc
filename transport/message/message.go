package message

// the frist byte is magic number
// the second byte is serialize type
type Header [2]byte


// Message is the generic type of Request and Response.
type Message struct {
	*Header

	Seq uint64 // sequence number chosen by client

	Error         string
	ServicePath   string
	ServiceMethod string
	Args          any
	Metadata      map[string]string
	Payload       []byte
}
