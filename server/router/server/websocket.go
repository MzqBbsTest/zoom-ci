package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketQueryBind struct {
	Id        int `form:"id"`
	SessionId int `form:"session_id"`
	W         int `form:"w"`
	H         int `form:"h"`
}

func WebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade websocket:", err)
		return
	}
	defer conn.Close()

	var query WebSocketQueryBind
	if err := c.ShouldBindQuery(&query); err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}

	defer func() {
		if manage.serMap[query.Id].session != nil {
			manage.serMap[query.Id].session[query.SessionId] = nil
		}
	}()

	var w sync.WaitGroup
	serClient, err := manage.generateSerClient(query.Id, query.SessionId)
	if err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}
	serClient.WindowChange(query.W, query.H)
	serClient.setCoon(conn)
	connList = append(connList, &MessageConn{
		w:         &w,
		serClient: serClient,
		conn:      conn,
	})

	//// 从 WebSocket 读取数据并发送到 SSH
	//go func() {
	//	defer w.Done()
	//	for {
	//		_, message, err := conn.ReadMessage()
	//		if err != nil {
	//			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	//			break
	//		}
	//
	//		err = serClient.write(&message)
	//		if err != nil {
	//			conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
	//			break
	//		}
	//	}
	//}()

	w.Add(1)
	w.Wait()

}
