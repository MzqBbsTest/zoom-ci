// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package api

const (
	SYSTEM_INSTALL        = "/system/install"
	SYSTEM_INSTALL_STATUS = "/system/install_status"

	LOGIN        = "/login"
	LOGOUT       = "/logout"
	LOGIN_STATUS = "/login/status"

	MY_USER_SETTING  = "/user/my/setting"
	MY_USER_PASSWORD = "/user/my/password"

	SERVER_GROUP_ADD    = "/server/group/add"
	SERVER_GROUP_LIST   = "/server/group/list"
	SERVER_GROUP_DELETE = "/server/group/delete"
	SERVER_GROUP_DETAIL = "/server/group/detail"
	SERVER_GROUP_UPDATE = "/server/group/update"
	SERVER_ADD          = "/server/add"
	SERVER_UPDATE       = "/server/update"
	SERVER_LIST         = "/server/list"
	SERVER_DELETE       = "/server/delete"
	SERVER_DETAIL       = "/server/detail"

	USER_ROLE_PRIV_LIST = "/user/role/privlist"
	USER_ROLE_ADD       = "/user/role/add"
	USER_ROLE_UPDATE    = "/user/role/update"
	USER_ROLE_LIST      = "/user/role/list"
	USER_ROLE_DETAIL    = "/user/role/detail"
	USER_ROLE_DELETE    = "/user/role/delete"
	USER_ADD            = "/user/add"
	USER_UPDATE         = "/user/update"
	USER_LIST           = "/user/list"
	USER_EXISTS         = "/user/exists"
	USER_DETAIL         = "/user/detail"
	USER_DELETE         = "/user/delete"

	PROJECT_SPACE_ADD     = "/project/space/add"
	PROJECT_SPACE_UPDATE  = "/project/space/update"
	PROJECT_SPACE_LIST    = "/project/space/list"
	PROJECT_SPACE_DETAIL  = "/project/space/detail"
	PROJECT_SPACE_DELETE  = "/project/space/delete"
	PROJECT_MEMBER_SEARCH = "/project/member/search"
	PROJECT_MEMBER_ADD    = "/project/member/add"
	PROJECT_MEMBER_LIST   = "/project/member/list"
	PROJECT_MEMBER_REMOVE = "/project/member/remove"
	PROJECT_ADD           = "/project/add"
	PROJECT_COPY           = "/project/copy"
	PROJECT_UPDATE        = "/project/update"
	PROJECT_LIST          = "/project/list"
	PROJECT_SWITCHSTATUS  = "/project/switchstatus"
	PROJECT_DETAIL        = "/project/detail"
	PROJECT_DELETE        = "/project/delete"
	PROJECT_BUILDSCRIPT   = "/project/buildscript"
	PROJECT_HOOKSCRIPT    = "/project/hookscript"

	DEPLOY_APPLY_PROJECT_DETAIL = "/deploy/apply/project/detail"
	DEPLOY_APPLY_SUBMIT         = "/deploy/apply/submit"
	DEPLOY_APPLY_PROJECT_ALL    = "/deploy/apply/project/all"
	DEPLOY_APPLY_LIST           = "/deploy/apply/list"
	DEPLOY_APPLY_DETAIL         = "/deploy/apply/detail"
	DEPLOY_APPLY_AUDIT          = "/deploy/apply/audit"
	DEPLOY_APPLY_UPDATE         = "/deploy/apply/update"
	DEPLOY_APPLY_DROP           = "/deploy/apply/drop"
	DEPLOY_APPLY_ROLLBACK       = "/deploy/apply/rollbacklist"
	DEPLOY_BUILD_START          = "/deploy/build/start"
	DEPLOY_BUILD_STATUS         = "/deploy/build/status"
	DEPLOY_BUILD_STOP           = "/deploy/build/stop"
	DEPLOY_DEPLOY_START         = "/deploy/deploy/start"
	DEPLOY_DEPLOY_STATUS        = "/deploy/deploy/status"
	DEPLOY_DEPLOY_STOP          = "/deploy/deploy/stop"
	DEPLOY_DEPLOY_ROLLBACK      = "/deploy/deploy/rollback"
)
