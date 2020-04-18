package steam

import (
	"github.com/benpye/go-steam/netutil"
	"github.com/benpye/go-steam/protocol"
)

// When this event is emitted by the Client, the connection is automatically closed.
// This may be caused by a network error, for example.
type FatalErrorEvent error

type ConnectedEvent struct{}

func (e *ConnectedEvent) GetJobID() protocol.JobID {
	return 0
}

type DisconnectedEvent struct {
	UserRequested bool
}

func (e *DisconnectedEvent) GetJobID() protocol.JobID {
	return 0
}

// A list of connection manager addresses to connect to in the future.
// You should always save them and then select one of these
// instead of the builtin ones for the next connection.
type ClientCMListEvent struct {
	Addresses []*netutil.PortAddr
}

func (e *ClientCMListEvent) GetJobID() protocol.JobID {
	return 0
}
