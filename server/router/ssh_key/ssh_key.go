package ssh_key

import (
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/model"
	"github.com/zoom-ci/zoom-ci/server/module/ssh_key"
	"github.com/zoom-ci/zoom-ci/server/render"
	"github.com/zoom-ci/zoom-ci/util/gostring"
	"github.com/zoom-ci/zoom-ci/util/table"
)

type QueryBind struct {
	//SpaceId int    `form:"space_id"`
	//Keyword string `form:"keyword"`
	Offset int `form:"offset"`
	Limit  int `form:"limit" binding:"required,gte=1,lte=999"`
}

type SshKeyFormBind struct {
	ID         int    `form:"id"`
	Name       string `form:"name" binding:"required"`
	AuthKey    bool   `form:"auth_key" binding:"required"`
	Password   string `form:"password"`
	PublicKey  string `form:"public_key"`
	PrivateKey string `form:"private_key"`
}

func SshKeyList(c *gin.Context) {
	var query QueryBind
	if err := c.ShouldBind(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	userId, exists := c.Get("user_id")
	if !exists {
		render.ParamError(c, "user_id cannot be empty")
		return
	}

	page := &table.Page{
		Where: []model.WhereParam{
			model.WhereParam{
				Field:   "user_id",
				Tag:     "=",
				Prepare: userId,
			},
		},
		Offset: query.Offset,
		Limit:  query.Limit,
	}
	data, err := page.Table(&ssh_key.SSHKey{})

	if err != nil {
		render.AppError(c, err.Error())
		return
	}

	render.JSON(c, gin.H{
		"list":  data["list"],
		"total": data["total"],
	})
}

func SshKeyAdd(c *gin.Context) {
	var form SshKeyFormBind
	if err := c.ShouldBind(&form); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	userId, _ := c.Get("user_id")

	// 自动生成签名
	if form.AuthKey {

	}

	info := &ssh_key.SSHKey{
		UserID:     userId.(int),
		Name:       form.Name,
		PublicKey:  form.PublicKey,
		PrivateKey: form.PrivateKey,
	}
	err := info.CreateOrUpdate()
	if err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.JSON(c, gin.H{
		"id": info.ID,
	})
}

func SshKeyUpdate(c *gin.Context) {
	var form SshKeyFormBind
	if err := c.ShouldBind(&form); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	userId, _ := c.Get("user_id")
	info := &ssh_key.SSHKey{
		ID:         form.ID,
		UserID:     userId.(int),
		Name:       form.Name,
		Password:   form.Password,
		PublicKey:  form.PublicKey,
		PrivateKey: form.PrivateKey,
	}
	err := info.CreateOrUpdate()
	if err != nil {
		render.AppError(c, err.Error())
		return
	}

	render.JSON(c, nil)
}

func SshKeyDelete(c *gin.Context) {
	id := gostring.Str2Int(c.PostForm("id"))
	if id == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}

	member := &ssh_key.SSHKey{
		ID: id,
	}
	if err := member.Delete(); err != nil {
		render.AppError(c, err.Error())
		return
	}

	render.JSON(c, nil)
}

func SshKeyDetail(c *gin.Context) {
	id, ok := c.GetQuery("id")
	if !ok || len(id) == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}

	member := &ssh_key.SSHKey{
		ID: gostring.Str2Int(id),
	}
	if err := member.Detail(); err != nil {
		render.AppError(c, err.Error())
		return
	}

	render.JSON(c, member)
}
