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

	// config
	api.GET(reqApi.SERVER_GROUP_CONFIG, server.GroupConfig)
	api.POST(reqApi.SERVER_GROUP_CONFIG, server.GroupConfigSave)
	api.PUT(reqApi.SERVER_GROUP_CONFIG, server.GroupConfigSave)
	api.GET(reqApi.SERVER_GROUP_CONFIG_ID, server.GroupConfigDetail)
	api.DELETE(reqApi.SERVER_GROUP_CONFIG, server.GroupConfigDelete)

	// ssh
	api.GET(reqApi.SERVER_SSH_TEST, server.ServerSshTest)
	api.GET(reqApi.SERVER_SESSION, server.CreateSession)
	api.GET(reqApi.SERVER_SESSION_RESIZE, server.SessionResize)

	api.GET(reqApi.SERVER_SFTP, server.SftpIndex)
	api.GET(reqApi.SERVER_SFTP_RENAME, server.SftpRename)
	api.GET(reqApi.SERVER_SFTP_MOD, server.SftpChmod)
	api.GET(reqApi.SERVER_SFTP_CREATE_DIR, server.SftpCreateDir)
	api.POST(reqApi.SERVER_SFTP_UPLOAD, server.SftpUploadFile)
	api.GET(reqApi.SERVER_SFTP_DELETE, server.SftpDeleteFile)
	api.POST(reqApi.SERVER_SFTP_ZIP, server.SftpCreateZip)
	api.POST(reqApi.SERVER_SFTP_UNZIP, server.SftpUnZip)
	api.GET(reqApi.SERVER_SFTP_DOWN, server.SftpDown)
}
