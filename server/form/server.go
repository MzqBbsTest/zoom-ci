package form

type ServerForm struct {
	Id       int    `form:"id"`
	GroupId  int    `form:"group_id"`
	Name     string `form:"name" binding:"required"`
	Ip       string `form:"ip" binding:"required"`
	SSHPort  int    `form:"ssh_port,default=22" binding:"gte=1,lte=65535"`
	User     string `form:"user" `
	Password string `form:"password" `
	SshKeyId int    `form:"sshkey_id"`
}
