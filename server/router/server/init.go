package server

import (
	"github.com/gorilla/websocket"
	"sync"
	"time"
)

var manage Manage

type MessageConn struct {
	w         *sync.WaitGroup
	serClient *ClientSession
	conn      *websocket.Conn
}

type MessageSsh struct {
	msg       []byte
	serClient *ClientSession
}

var connList []*MessageConn

var sshChan chan MessageSsh

func init() {
	manage = Manage{
		serMap: map[int]*client{},
	}
	sshChan = make(chan MessageSsh, 1000)

	go func() {
		for {
			select {
			case message := <-sshChan:
				err := message.serClient.conn.WriteMessage(websocket.BinaryMessage, message.msg)
				if err != nil {
					println("error", message.msg)
				}
			default:
				//println("write")
			}
			time.Sleep(100 * time.Microsecond)
		}
	}()

}
