package route

import (
	"github.com/gin-gonic/gin"
	reqApi "github.com/zoom-ci/zoom-ci/server/router/route/api"
	"github.com/zoom-ci/zoom-ci/server/router/server"
)

func serverGroup(api *gin.RouterGroup) {

	api.POST(reqApi.SERVER_ADD, server.ServerSave)
	api.POST(reqApi.SERVER_UPDATE, server.ServerSave)
	api.GET(reqApi.SERVER_LIST, server.ServerList)
	api.POST(reqApi.SERVER_DELETE, server.ServerDelete)
	api.GET(reqApi.SERVER_DETAIL, server.ServerDetail)

	api.POST(reqApi.SERVER_GROUP_ADD, server.GroupAdd)
	api.GET(reqApi.SERVER_GROUP_LIST, server.GroupList)
	api.POST(reqApi.SERVER_GROUP_DELETE, server.GroupDelete)
	api.GET(reqApi.SERVER_GROUP_DETAIL, server.GroupDetail)
	api.POST(reqApi.SERVER_GROUP_UPDATE, server.GroupUpdate)

	// path
	api.GET(reqApi.SERVER_GROUP_PATH, server.GroupPath)
	api.POST(reqApi.SERVER_GROUP_PATH, server.GroupPathSave)
	api.PUT(reqApi.SERVER_GROUP_PATH, server.GroupPathSave)
	api.GET(reqApi.SERVER_GROUP_PATH_ID, server.GroupPathDetail)
	api.DELETE(reqApi.SERVER_GROUP_PATH, server.GroupPathDelete)
	api.GET(reqApi.SERVER_GROUP_PATH_ALIAS, server.GroupPathAlias)

	// config
	api.GET(reqApi.SERVER_GROUP_CONFIG, server.GroupConfig)
	api.POST(reqApi.SERVER_GROUP_CONFIG, server.GroupConfigSave)
	api.PUT(reqApi.SERVER_GROUP_CONFIG, server.GroupConfigSave)
	api.GET(reqApi.SERVER_GROUP_CONFIG_ID, server.GroupConfigDetail)
	api.DELETE(reqApi.SERVER_GROUP_CONFIG, server.GroupConfigDelete)
	api.GET(reqApi.SERVER_GROUP_CONFIG_ALIAS, server.GroupConfigAlias)

	// config
	api.GET(reqApi.SERVER_GROUP_CMD, server.GroupCmd)
	api.POST(reqApi.SERVER_GROUP_CMD, server.GroupCmdSave)
	api.PUT(reqApi.SERVER_GROUP_CMD, server.GroupCmdSave)
	api.GET(reqApi.SERVER_GROUP_CMD_ID, server.GroupCmdDetail)
	api.DELETE(reqApi.SERVER_GROUP_CMD, server.GroupCmdDelete)
}
