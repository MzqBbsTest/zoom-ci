package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/sftp"
	"github.com/zoom-ci/zoom-ci/server/module/server"
	"github.com/zoom-ci/zoom-ci/server/render"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
	"os"
	"path"
	"strconv"
)

type SftpQueryBind struct {
	Id      int    `form:"id"`
	Path    string `form:"path"`
	NewPath string `form:"new_path"`
	Mod     string `form:"mod"`
	ModUser string `form:"mod_user"`
}

func connect(query SftpQueryBind) (*ssh.Client, error) {
	ser := &server.Server{
		ID: query.Id,
	}

	if err := ser.Detail(); err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: ser.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(ser.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	// 建立 SSH 连接
	return ssh.Dial("tcp", fmt.Sprintf("%s:%d", ser.Ip, ser.SSHPort), config)
}

func SftpIndex(c *gin.Context) {
	var query SftpQueryBind
	if err := c.ShouldBindQuery(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	// 建立 SSH 连接
	conn, err := connect(query)
	if err != nil {
		//log.Fatalf("Failed to dial: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer conn.Close()

	// 建立 SFTP 连接
	client, err := sftp.NewClient(conn)
	if err != nil {
		//log.Fatalf("Failed to create client: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer client.Close()

	// 远程目录路径
	remoteDir := query.Path
	if len(remoteDir) == 0 {
		remoteDir = "."
	}

	if remoteDir[len(remoteDir)-1] == '.' {
		remoteDir, err = client.RealPath(remoteDir)
		if err != nil {
			render.ParamError(c, err.Error())
			return
		}
	}

	// 读取远程目录
	files, err := client.ReadDir(remoteDir)
	if err != nil {
		//log.Fatalf("Failed to read directory: %s", err)
		render.ParamError(c, err.Error())
		return
	}

	var _files []map[string]interface{}
	fileCount := 0
	dirCount := 0
	for _, file := range files {

		fileInfo := getFileInfo(file, remoteDir)

		if file.Mode()&os.ModeSymlink != 0 {
			// 这是一个符号链接
			target, err := client.ReadLink(remoteDir + file.Name())
			if err != nil {
				render.ParamError(c, err.Error())
				return
			}
			fileInfo["link_target"] = target
		}

		if fileInfo["type"] == 'd' {
			dirCount += 1
		} else {
			fileCount += 1
		}
		_files = append(_files, fileInfo)
	}

	render.JSON(c, map[string]interface{}{
		"files":       _files,
		"file_count":  fileCount,
		"dir_count":   dirCount,
		"current_dir": remoteDir,
	})
}

func SftpCreateDir(c *gin.Context) {
	var query SftpQueryBind
	if err := c.ShouldBindQuery(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	conn, err := connect(query)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer conn.Close()

	// 建立 SFTP 连接
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("Failed to create client: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer client.Close()

	err = client.Mkdir(query.Path)
	if err != nil {
		log.Fatalf("Failed to create directory: %s", err)
		render.ParamError(c, err.Error())
		return
	}

	render.JSON(c, map[string]interface{}{})
}

func SftpUploadFile(c *gin.Context) {
	var query SftpQueryBind
	if err := c.ShouldBind(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	form, err := c.MultipartForm()
	if err != nil {
		render.ParamError(c, err.Error())
		return
	}

	// 建立 SSH 连接
	conn, err := connect(query)
	if err != nil {
		render.ParamError(c, err.Error())
		return
	}
	defer conn.Close()

	// 建立 SFTP 连接
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("Failed to create client: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer client.Close()

	var ret []string
	files := form.File["files"]
	for _, file := range files {
		srcFile, err := file.Open()
		if err != nil {
			continue
		}
		fileName := file.Filename
		dstFile, err := client.Create(path.Join(query.Path, fileName))
		if err != nil {
			continue
		}
		_, err = io.Copy(dstFile, srcFile)
		if err != nil {
			continue
		}
		_ = srcFile.Close()
		_ = dstFile.Close()
		ret = append(ret, fileName)
	}
	msg := strconv.Itoa(len(ret)) + " 个文件上传成功"
	render.JSON(c, map[string]interface{}{
		"msg":  msg,
		"data": ret,
	})
}

func SftpDeleteFile(c *gin.Context) {
	var query SftpQueryBind
	if err := c.ShouldBindQuery(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	// 建立 SSH 连接
	conn, err := connect(query)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer conn.Close()

	// 建立 SFTP 连接
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("Failed to create client: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer client.Close()

	fileInfo, err := client.Stat(query.Path)
	if err != nil {
		render.ParamError(c, err.Error())
		return
	}

	if fileInfo.IsDir() == false {
		err = client.Remove(query.Path)
		if err != nil {
			render.ParamError(c, err.Error())
			return
		}

		render.JSON(c, map[string]interface{}{})
		return
	}

	err = SftpDeleteDir(client, query.Path)
	if err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.JSON(c, map[string]interface{}{})
}

func SftpDeleteDir(client *sftp.Client, rootPath string) error {
	files, err := client.ReadDir(rootPath)
	if err != nil {
		return err
	}

	for _, file := range files {
		// 获取当前文件或子目录的完整路径
		fullPath := path.Join(rootPath, file.Name())

		if file.IsDir() {
			// 如果是目录，递归删除子目录
			err := SftpDeleteDir(client, fullPath)
			if err != nil {
				return err
			}
		} else {
			// 如果是文件，删除文件
			err := client.Remove(fullPath)
			if err != nil {
				return err
			}
		}
	}
	err = client.Remove(rootPath)
	if err != nil {
		return err
	}
	return nil
}

func SftpCreateZip(c *gin.Context) {
	var query SftpQueryBind
	if err := c.ShouldBind(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	// 建立 SSH 连接
	conn, err := connect(query)
	if err != nil {
		render.ParamError(c, err.Error())
		return
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		//log.Fatalf("Failed to create session: %v", err)
		render.AppError(c, err.Error())
		return
	}
	defer session.Close()

	if _, err = session.CombinedOutput(fmt.Sprintf("tar -cf %s.tar %s", query.NewPath, query.Path)); err != nil {
		render.AppError(c, err.Error())
		return
	}

	render.JSON(c, map[string]interface{}{
		"zip": fmt.Sprintf("%s.tar", query.NewPath),
	})
}

func SftpUnZip(c *gin.Context) {
	var query SftpQueryBind
	if err := c.ShouldBind(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	// 建立 SSH 连接
	conn, err := connect(query)
	if err != nil {
		//log.Fatalf("Failed to dial: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		//log.Fatalf("Failed to create session: %v", err)
		render.AppError(c, err.Error())
		return
	}
	defer session.Close()

	fix := query.Path[len(query.Path)-3:]
	cmd := ""
	switch fix {
	case "zip":
		cmd = fmt.Sprintf("unzip  %s -d %s", query.Path, query.NewPath)
	case "tar":
		cmd = fmt.Sprintf("mkdir -p  %s && tar -xvf %s -C %s", query.NewPath, query.Path, query.NewPath)
	case ".7z":
		cmd = fmt.Sprintf("7z -x %s.7z -r -o%s", query.Path, query.NewPath)
	}

	if _, err = session.CombinedOutput(cmd); err != nil {
		render.AppError(c, err.Error())
		return
	}

	render.JSON(c, map[string]interface{}{
		"files": query.NewPath,
	})
}

func SftpRename(c *gin.Context) {
	var query SftpQueryBind
	if err := c.ShouldBindQuery(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	// 建立 SSH 连接
	conn, err := connect(query)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer conn.Close()

	// 建立 SFTP 连接
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("Failed to create client: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer client.Close()
	err = client.Rename(query.Path, query.NewPath)

	if err != nil {
		render.ParamError(c, err.Error())
		return
	}

	render.JSON(c, map[string]interface{}{})
}

func SftpDown(c *gin.Context) {
	var query SftpQueryBind
	if err := c.ShouldBindQuery(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	// 建立 SSH 连接
	conn, err := connect(query)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer conn.Close()

	// 建立 SFTP 连接
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatalf("Failed to create client: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer client.Close()

	stat, err := client.Stat(query.Path)
	if err != nil {
		log.Fatalf("Failed to create client: %s", err)
		render.ParamError(c, err.Error())
		return
	}

	if stat.IsDir() {
		render.ParamError(c, "down only file")
		return
	}

	remoteFile, err := client.Open(query.Path)
	if err != nil {
		log.Fatalf("Failed to open remote file: %s", err)
	}
	defer remoteFile.Close()

	localFilePath := "./" + stat.Name()
	localFile, err := os.Create(localFilePath)
	if err != nil {
		log.Fatalf("Failed to create local file: %s", err)
	}
	defer localFile.Close()

	// 复制文件内容到本地文件
	_, err = io.Copy(localFile, remoteFile)
	if err != nil {
		log.Fatalf("Failed to copy file: %s", err)
	}

	// 设置下载头并发送文件
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", stat.Name()))
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", fmt.Sprintf("%d", stat.Size()))
	c.File(localFile.Name())
}

func SftpChmod(c *gin.Context) {
	var query SftpQueryBind
	if err := c.ShouldBindQuery(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	// 建立 SSH 连接
	conn, err := connect(query)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		//log.Fatalf("Failed to create session: %v", err)
		render.AppError(c, err.Error())
		return
	}
	defer session.Close()

	if _, err = session.CombinedOutput(fmt.Sprintf("chmod -R %s %s", query.Mod, query.Path)); err != nil {
		render.AppError(c, err.Error())
		return
	}

	render.JSON(c, map[string]interface{}{})
}

func SftpChown(c *gin.Context) {
	var query SftpQueryBind
	if err := c.ShouldBindQuery(&query); err != nil {
		render.ParamError(c, err.Error())
		return
	}

	// 建立 SSH 连接
	conn, err := connect(query)
	if err != nil {
		log.Fatalf("Failed to dial: %s", err)
		render.ParamError(c, err.Error())
		return
	}
	defer conn.Close()

	session, err := conn.NewSession()
	if err != nil {
		//log.Fatalf("Failed to create session: %v", err)
		render.AppError(c, err.Error())
		return
	}
	defer session.Close()

	if _, err = session.CombinedOutput(fmt.Sprintf("chmod -R %s %s", query.ModUser, query.Path)); err != nil {
		render.AppError(c, err.Error())
		return
	}

	render.JSON(c, map[string]interface{}{})
}
