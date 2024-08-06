package server

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zoom-ci/zoom-ci/server/render"
	"log"
	"net/http"
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
		render.RepeatError(c, err.Error())
		return
	}

	var query WebSocketQueryBind
	if err := c.ShouldBindQuery(&query); err != nil {
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}

	//var w sync.WaitGroup
	serClient, err := manage.generateSerClient(query.Id, query.SessionId)
	if err != nil {
		log.Println("generateSerClient:", err)
		conn.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		return
	}

	serClient.run(conn)

}
