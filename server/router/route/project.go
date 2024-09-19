package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/router/project"
	reqApi "github.com/zoom-ci/zoom-ci/server/router/route/api"
)

func projectGroup(api *gin.RouterGroup) {
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
}
