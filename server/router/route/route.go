// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/router/system"
	"html/template"

	"github.com/zoom-ci/zoom-ci"
	"github.com/zoom-ci/zoom-ci/server/router/middleware"
	reqApi "github.com/zoom-ci/zoom-ci/server/router/route/api"
	"github.com/zoom-ci/zoom-ci/server/router/server"
	"github.com/zoom-ci/zoom-ci/server/router/ssh_key"
	"github.com/zoom-ci/zoom-ci/server/router/user"
	"net/http"
)

func RegisterRoute() {
	api := zoom.App.Gin.Group("/api", middleware.ApiPriv())
	{
		api.POST(reqApi.SYSTEM_INSTALL, system.Install)
		api.GET(reqApi.SYSTEM_INSTALL_STATUS, system.InstallStatus)
		api.POST(reqApi.SYSTEM_STATUS, system.Status)

		api.POST(reqApi.LOGIN, user.Login)
		api.POST(reqApi.LOGOUT, user.Logout)
		api.GET(reqApi.LOGIN_STATUS, user.LoginStatus)

		api.POST(reqApi.MY_USER_SETTING, user.MyUserSetting)
		api.POST(reqApi.MY_USER_PASSWORD, user.MyUserPassword)

		serverGroup(api)
		projectGroup(api)
		userGroup(api)
		deployGroup(api)
		xtermGroup(api)

		api.GET(reqApi.SSHKEY_LIST, ssh_key.SshKeyList)
		api.POST(reqApi.SSHKEY_ADD, ssh_key.SshKeyAdd)
		api.POST(reqApi.SSHKEY_UPDATE, ssh_key.SshKeyUpdate)
		api.POST(reqApi.SSHKEY_DELETE, ssh_key.SshKeyDelete)
		api.GET(reqApi.SSHKEY_DETAIL, ssh_key.SshKeyDetail)

	}
	RegisterWebRouter()
}

func RegisterWebRouter() {
	tpl := template.Must(template.New("").ParseFS(zoom.WebPath, "web/dist/*.html"))
	zoom.App.Gin.SetHTMLTemplate(tpl)
	zoom.App.Gin.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"appName": zoom.App.Config.Zoom.AppName,
		})
	})
	zoom.App.Gin.GET("/web/dist/*filepath", gin.WrapH(http.FileServer(http.FS(zoom.WebPath))))

	zoom.App.Gin.GET("/api/ws", server.WebSocket)
}
