package route

import (
	"github.com/gin-gonic/gin"
	reqApi "github.com/zoom-ci/zoom-ci/server/router/route/api"
	"github.com/zoom-ci/zoom-ci/server/router/server"
)

func xtermGroup(api *gin.RouterGroup) {

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

	api.GET(reqApi.CMD_LIST, server.CmdList)
	api.POST(reqApi.CMD_ADD, server.CmdSave)
	api.POST(reqApi.CMD_UPDATE, server.CmdSave)
	api.POST(reqApi.CMD_DELETE, server.CmdDelete)

}
