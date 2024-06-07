package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/render"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
	"strconv"
	"sync"
)

func CreateSession(c *gin.Context) {
	query, ok := c.GetQuery("id")
	if !ok {
		render.ParamError(c, "id cannot be empty")
		return
	}
	id, err := strconv.Atoi(query)
	if err != nil {
		render.ParamError(c, err.Error())
		return
	}

	sessionId := 0
	query, ok = c.GetQuery("session_id")
	if ok {
		sessionId, err = strconv.Atoi(query)
		if err != nil {
			render.ParamError(c, err.Error())
			return
		}
	}

	session, err := manage.generateSerClient(id, sessionId)
	if err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.JSON(c, map[string]interface{}{
		"server_id":  id,
		"session_id": session.sessionId,
	})
}

type WebSocketReader struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func (r *WebSocketReader) Read(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	var msg string
	err = wsjson.Read(context.Background(), r.conn, &msg)
	if err != nil {
		return 0, err
	}

	copy(p, msg)
	return len(msg), nil
}

// WebSocketWriter 实现 io.Writer 接口
type WebSocketWriter struct {
	conn *websocket.Conn
	mu   sync.Mutex
}

func (w *WebSocketWriter) Write(p []byte) (n int, err error) {
	w.mu.Lock()
	defer w.mu.Unlock()

	err = wsjson.Write(context.Background(), w.conn, string(p))
	if err != nil {
		return 0, err
	}

	return len(p), nil
}
