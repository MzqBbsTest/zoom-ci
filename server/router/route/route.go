// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/router/system"
	"html/template"

	"github.com/zoom-ci/zoom-ci"
	"github.com/zoom-ci/zoom-ci/server/router/deploy"
	"github.com/zoom-ci/zoom-ci/server/router/middleware"
	"github.com/zoom-ci/zoom-ci/server/router/project"
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

		api.POST(reqApi.SERVER_GROUP_ADD, server.GroupAdd)
		api.GET(reqApi.SERVER_GROUP_LIST, server.GroupList)
		api.POST(reqApi.SERVER_GROUP_DELETE, server.GroupDelete)
		api.GET(reqApi.SERVER_GROUP_DETAIL, server.GroupDetail)
		api.POST(reqApi.SERVER_GROUP_UPDATE, server.GroupUpdate)
		api.GET(reqApi.SERVER_GROUP_PATH, server.GroupPath)
		api.POST(reqApi.SERVER_GROUP_PATH, server.GroupPathAdd)

		api.POST(reqApi.SERVER_ADD, server.ServerAdd)
		api.POST(reqApi.SERVER_UPDATE, server.ServerUpdate)
		api.GET(reqApi.SERVER_LIST, server.ServerList)
		api.POST(reqApi.SERVER_DELETE, server.ServerDelete)
		api.GET(reqApi.SERVER_DETAIL, server.ServerDetail)
		api.GET(reqApi.SERVER_SSH_TEST, server.ServerSshTest)
		api.GET(reqApi.SERVER_SESSION, server.CreateSession)

		api.GET(reqApi.USER_ROLE_PRIV_LIST, user.RolePrivList)
		api.POST(reqApi.USER_ROLE_ADD, user.RoleAdd)
		api.POST(reqApi.USER_ROLE_UPDATE, user.RoleUpdate)
		api.GET(reqApi.USER_ROLE_LIST, user.RoleList)
		api.GET(reqApi.USER_ROLE_DETAIL, user.RoleDetail)
		api.POST(reqApi.USER_ROLE_DELETE, user.RoleDelete)
		api.POST(reqApi.USER_ADD, user.UserAdd)
		api.POST(reqApi.USER_UPDATE, user.UserUpdate)
		api.GET(reqApi.USER_LIST, user.UserList)
		api.GET(reqApi.USER_EXISTS, user.UserExists)
		api.GET(reqApi.USER_DETAIL, user.UserDetail)
		api.POST(reqApi.USER_DELETE, user.UserDelete)

		api.POST(reqApi.PROJECT_SPACE_ADD, project.SpaceAdd)
		api.POST(reqApi.PROJECT_SPACE_UPDATE, project.SpaceUpdate)
		api.GET(reqApi.PROJECT_SPACE_LIST, project.SpaceList)
		api.GET(reqApi.PROJECT_SPACE_DETAIL, project.SpaceDetail)
		api.POST(reqApi.PROJECT_SPACE_DELETE, project.SpaceDelete)
		api.GET(reqApi.PROJECT_MEMBER_SEARCH, project.MemberSearch)
		api.POST(reqApi.PROJECT_MEMBER_ADD, project.MemberAdd)
		api.GET(reqApi.PROJECT_MEMBER_LIST, project.MemberList)
		api.POST(reqApi.PROJECT_MEMBER_REMOVE, project.MemberRemove)
		api.POST(reqApi.PROJECT_ADD, project.ProjectAdd)
		api.POST(reqApi.PROJECT_COPY, project.ProjectCopy)
		api.POST(reqApi.PROJECT_UPDATE, project.ProjectUpdate)
		api.GET(reqApi.PROJECT_LIST, project.ProjectList)
		api.POST(reqApi.PROJECT_SWITCHSTATUS, project.ProjectSwitchStatus)
		api.GET(reqApi.PROJECT_DETAIL, project.ProjectDetail)
		api.POST(reqApi.PROJECT_DELETE, project.ProjectDelete)
		api.POST(reqApi.PROJECT_BUILDSCRIPT, project.ProjectBuildScript)
		api.POST(reqApi.PROJECT_HOOKSCRIPT, project.ProjectHookScript)

		api.GET(reqApi.DEPLOY_APPLY_PROJECT_DETAIL, deploy.ApplyProjectDetail)
		api.POST(reqApi.DEPLOY_APPLY_SUBMIT, deploy.ApplySubmit)
		api.GET(reqApi.DEPLOY_APPLY_PROJECT_ALL, deploy.ApplyProjectAll)
		api.GET(reqApi.DEPLOY_APPLY_LIST, deploy.ApplyList)
		api.GET(reqApi.DEPLOY_APPLY_DETAIL, deploy.ApplyDetail)
		api.POST(reqApi.DEPLOY_APPLY_AUDIT, deploy.ApplyAudit)
		api.POST(reqApi.DEPLOY_APPLY_UPDATE, deploy.ApplyUpdate)
		api.POST(reqApi.DEPLOY_APPLY_DROP, deploy.ApplyDrop)
		api.POST(reqApi.DEPLOY_BUILD_START, deploy.BuildStart)
		api.GET(reqApi.DEPLOY_BUILD_STATUS, deploy.BuildStatus)
		api.POST(reqApi.DEPLOY_BUILD_STOP, deploy.BuildStop)
		api.POST(reqApi.DEPLOY_DEPLOY_START, deploy.DeployStart)
		api.GET(reqApi.DEPLOY_DEPLOY_STATUS, deploy.DeployStatus)
		api.POST(reqApi.DEPLOY_DEPLOY_STOP, deploy.DeployStop)
		api.GET(reqApi.DEPLOY_APPLY_ROLLBACK, deploy.ApplyRollbackList)
		api.POST(reqApi.DEPLOY_DEPLOY_ROLLBACK, deploy.DeployRollback)

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

	//zoom.App.Gin.GET("/api/session/create", server.CreateSession)
	zoom.App.Gin.GET("/api/ws", server.WebSocket)
}
