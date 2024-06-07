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
	sshChan := make(chan MessageSsh, 1000)

	go func() {
		for {
			for i, connMessage := range connList {
				if connMessage == nil {
					continue
				}

				// 讀取消息
				_, message, err := connMessage.conn.ReadMessage()
				if err != nil {
					connList[i] = nil
					connMessage.w.Done()
					continue
				}

				// 写消息
				err = connMessage.serClient.write(&message)
				if err != nil {
					connMessage.conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
					break
				}

			}
			time.Sleep(100)
		}
	}()

	go func() {
		for {
			select {
			case message := <-sshChan:
				println(message)
			default:
				println("write")
			}
		}
	}()

}
