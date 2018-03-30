package web

import (
	"clarissa/kernel"

	"github.com/googollee/go-socket.io"
)

func (e *kernel.Event) processEvent(so socketio.Socket, room string) {
	// Send to the specified room the details of the event
	so.Emit(e.Type, e.Details)
}
