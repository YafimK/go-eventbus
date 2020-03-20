package eventbus

type Header struct {
	Topic string

	done chan interface{}
}

type Messsage struct {
	MsgID           uint64
	TransportHeader Header
	Message         interface{}
}
