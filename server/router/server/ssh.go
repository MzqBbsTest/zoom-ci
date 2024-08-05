package server

import (
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/render"
	"strconv"
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

func SessionResize(c *gin.Context) {
	var query WebSocketQueryBind
	if err := c.ShouldBindQuery(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	id := query.Id
	sessionId := query.SessionId
	if manage.serMap[id] == nil {
		render.NoDataError(c, "id not found")
		return
	}
	if manage.serMap[id].session[sessionId] == nil {
		render.NoDataError(c, "session id not found")
		return
	}

	err := manage.serMap[id].session[sessionId].WindowChange(query.W, query.H)
	if err != nil {
		render.NoDataError(c, err.Error())
		return
	}
	render.JSON(c, map[string]interface{}{})
}
