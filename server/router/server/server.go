// Copyright 2021 Zoom Author. All Rights Reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/server/form"
	"github.com/zoom-ci/zoom-ci/server/module/server"
	query2 "github.com/zoom-ci/zoom-ci/server/query"
	"github.com/zoom-ci/zoom-ci/server/render"
	"github.com/zoom-ci/zoom-ci/util/gostring"
	"golang.org/x/crypto/ssh"
)

//type QueryBind struct {
//	Keyword string `form:"keyword"`
//	GroupId string `form:"group_id"`
//	Offset  int    `form:"offset"`
//	Limit   int    `form:"limit,default=20" binding:"gte=1,lte=999" `
//}

//type ServerForm struct {
//	GroupId  int    `form:"group_id"`
//	Name     string `form:"name" binding:"required"`
//	Ip       string `form:"ip" binding:"required"`
//	SSHPort  int    `form:"ssh_port,default=22" binding:"gte=1,lte=65535"`
//	User     string `form:"user" `
//	Password string `form:"password" `
//	SshKeyId int    `form:"sshkey_id"`
//}

func ServerAdd(c *gin.Context) {
	serverCreateOrUpdate(c, 0)
}

func ServerUpdate(c *gin.Context) {
	id := gostring.Str2Int(c.PostForm("id"))
	if id == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}
	serverCreateOrUpdate(c, id)
}

func serverCreateOrUpdate(c *gin.Context, id int) {
	var serverForm form.ServerForm
	if err := c.ShouldBind(&serverForm); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	server := &server.Server{
		ID:       id,
		Name:     serverForm.Name,
		Ip:       serverForm.Ip,
		SSHPort:  serverForm.SSHPort,
		User:     serverForm.User,
		Password: serverForm.Password,
		SshKeyId: serverForm.SshKeyId,
	}
	if err := server.CreateOrUpdate(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.Success(c)
}

func ServerList(c *gin.Context) {
	var query query2.BindServer
	if err := c.ShouldBind(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}
	ser := &server.Server{}
	list, err := ser.List(&query)
	if err != nil {
		render.AppError(c, err.Error())
		return
	}

	total, err := ser.Total(query.Keyword)
	if err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.JSON(c, gin.H{
		"list":  list,
		"total": total,
	})
}

func ServerDelete(c *gin.Context) {
	id := gostring.Str2Int(c.PostForm("id"))
	if id == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}
	server := &server.Server{
		ID: id,
	}
	if err := server.Delete(); err != nil {
		render.AppError(c, err.Error())
		return
	}
	render.JSON(c, nil)
}

func ServerDetail(c *gin.Context) {
	id := gostring.Str2Int(c.Query("id"))
	if id == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}
	ser := &server.Server{
		ID: id,
	}
	if err := ser.Detail(); err != nil {
		render.AppError(c, err.Error())
		return
	}

	//group := &server.Group{
	//	ID: ser.GroupId,
	//}
	//if err := group.Detail(); err == nil {
	//	ser.GroupName = group.Name
	//}

	render.JSON(c, ser)
}

func ServerSshTest(c *gin.Context) {
	id := gostring.Str2Int(c.Query("id"))
	if id == 0 {
		render.ParamError(c, "id cannot be empty")
		return
	}
	ser := &server.Server{
		ID: id,
	}
	if err := ser.Detail(); err != nil {
		render.AppError(c, err.Error())
		return
	}

	config := &ssh.ClientConfig{
		User: ser.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(ser.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 连接到服务器
	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", ser.Ip, ser.SSHPort), config)
	if err != nil {
		//log.Fatalf("Failed to dial: %v", err)
		render.AppError(c, err.Error())
		return
	}
	defer client.Close()

	// 创建一个新的会话
	session, err := client.NewSession()
	if err != nil {
		//log.Fatalf("Failed to create session: %v", err)
		render.AppError(c, err.Error())
		return
	}
	defer session.Close()

	// 运行命令
	var b []byte
	if b, err = session.CombinedOutput("uname -a"); err != nil {
		//log.Fatalf("Failed to run: %v", err)
		render.AppError(c, err.Error())
		return
	}
	fmt.Println(string(b))
	render.JSON(c, nil)
}
