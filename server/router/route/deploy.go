package route

import (
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/router/deploy"
	reqApi "github.com/zoom-ci/zoom-ci/server/router/route/api"
)

func deployGroup(api *gin.RouterGroup) {

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

}
