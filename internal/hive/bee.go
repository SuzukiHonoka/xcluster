package hive

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"xcluster/internal/server"
)

type Bee struct {
	ServerID        server.ID
	Conn            *websocket.Conn
	MessageReceiver chan Message
	working         bool
}

func NewBee(serverID server.ID, conn *websocket.Conn) *Bee {
	bee := &Bee{
		ServerID:        serverID,
		Conn:            conn,
		MessageReceiver: make(chan Message),
		working:         true,
	}
	bee.Communicate()
	return bee
}

func (b *Bee) String() string {
	return fmt.Sprintf("bee (serverID=%s, remote=%s)", b.ServerID, b.Conn.RemoteAddr())
}

func (b *Bee) Communicate() {
	go func() {
		err := b.communicate()
		if err != nil {
			// filter normal closure
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				log.Printf("%s: websocket disconnted", b)
				return
			}
			log.Printf("%s: communication failed, err=%s", b, err)
		}
	}()
}

func (b *Bee) communicate() error {
	for b.working {
		var msg Message
		if err := b.Conn.ReadJSON(&msg); err != nil {
			// assume the bee has been attacked, left or dead
			b.working = false
			Central.Refrain <- b
			return err
		}
		// pass msg to processor
	}
	return nil
}

func (b *Bee) GoAway() error {
	if b.Conn == nil {
		return nil
	}
	return b.Conn.Close()
}
