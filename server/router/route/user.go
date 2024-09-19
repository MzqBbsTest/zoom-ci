package route

import (
	"github.com/gin-gonic/gin"
	reqApi "github.com/zoom-ci/zoom-ci/server/router/route/api"
	"github.com/zoom-ci/zoom-ci/server/router/user"
)

func userGroup(api *gin.RouterGroup) {
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
}
