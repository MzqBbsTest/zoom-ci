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

func init() {
	manage = Manage{
		serMap: map[int]*client{},
	}

	go func() {

		// 自动清理掉线的session
		for {
			for serverId, session := range manage.serMap {
				_session := map[int]*ClientSession{}
				for sessionId, v := range session.session {
					if v == nil {
						continue
					}
					_session[sessionId] = v
				}
				manage.serMap[serverId].session = _session
			}
			time.Sleep(10 * time.Minute)
		}
	}()

}
