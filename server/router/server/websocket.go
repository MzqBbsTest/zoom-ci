package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/zoom-ci/zoom-ci/server/module/server"
	"github.com/zoom-ci/zoom-ci/server/render"
	"github.com/zoom-ci/zoom-ci/util/gostring"
	"golang.org/x/crypto/ssh"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade websocket:", err)
		return
	}
	defer conn.Close()
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

	// SSH 配置
	config := &ssh.ClientConfig{
		User: ser.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(ser.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 仅用于测试，生产环境应使用安全的 HostKeyCallback
	}

	// 连接到 SSH 服务器
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

	// 获取会话的输入和输出
	sessionStdIn, err := session.StdinPipe()
	if err != nil {
		//log.Fatalf("Failed to get session stdin: %v", err)
		render.AppError(c, err.Error())
		return
	}
	sessionStdOut, err := session.StdoutPipe()
	if err != nil {
		//log.Fatalf("Failed to get session stdout: %v", err)
		render.AppError(c, err.Error())
		return
	}
	sessionStdErr, err := session.StderrPipe()
	if err != nil {
		//log.Fatalf("Failed to get session stderr: %v", err)
		render.AppError(c, err.Error())
		return
	}

	// 启动伪终端
	err = session.RequestPty("xterm", 80, 24, ssh.TerminalModes{
		ssh.ECHO: 0, // 关闭回显
	})
	if err != nil {
		//log.Fatalf("Failed to request pty: %v", err)
		render.AppError(c, err.Error())
		return
	}

	// 启动 shell
	err = session.Shell()
	if err != nil {
		//log.Fatalf("Failed to start shell: %v", err)
		render.AppError(c, err.Error())
		return
	}

	// 从 SSH 读取数据并发送到 WebSocket
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := sessionStdOut.Read(buf)
			if err != nil {
				break
			}
			conn.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := sessionStdErr.Read(buf)
			if err != nil {
				break
			}
			conn.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()

	// 从 WebSocket 读取数据并发送到 SSH
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				sessionStdIn.Close()
				break
			}
			sessionStdIn.Write(message)
		}
	}()

	// 等待会话结束
	session.Wait()

	//cmd := exec.Command("ssh", "your_username@remote_server_address")
	//cmdStdIn, _ := cmd.StdinPipe()
	//cmdStdOut, _ := cmd.StdoutPipe()
	//cmdStdErr, _ := cmd.StderrPipe()
	//
	//go func() {
	//	buf := make([]byte, 1024)
	//	for {
	//		n, err := cmdStdOut.Read(buf)
	//		if err != nil {
	//			break
	//		}
	//		conn.WriteMessage(websocket.TextMessage, buf[:n])
	//	}
	//}()
	//
	//go func() {
	//	buf := make([]byte, 1024)
	//	for {
	//		n, err := cmdStdErr.Read(buf)
	//		if err != nil {
	//			break
	//		}
	//		conn.WriteMessage(websocket.TextMessage, buf[:n])
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		_, message, err := conn.ReadMessage()
	//		if err != nil {
	//			cmd.Process.Kill()
	//			break
	//		}
	//		cmdStdIn.Write(message)
	//	}
	//}()
	//
	//cmd.Run()
}
