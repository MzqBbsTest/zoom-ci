package server

import (
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
)

type ClientSession struct {
	serverId  int // 服务器id
	sessionId int // m.serMap[id].session下标 用来寻址 m.serMap[id].session
	session   *ssh.Session
	stdin     io.WriteCloser
	stdout    io.Reader
	stderr    io.Reader
	conn      *websocket.Conn
}

func (s *ClientSession) login(id int) error {
	if s.session != nil {
		return nil
	}

	session, err := manage.serMap[id].client.NewSession()
	if err != nil {
		return err
	}
	s.session = session

	// 获取会话的输入和输出
	sessionStdIn, err := session.StdinPipe()
	if err != nil {
		return err
	}
	s.stdin = sessionStdIn

	sessionStdOut, err := session.StdoutPipe()
	if err != nil {
		return err
	}
	s.stdout = sessionStdOut

	sessionStdErr, err := session.StderrPipe()
	if err != nil {
		return err
	}
	s.stdout = sessionStdErr

	// 启动伪终端
	err = session.RequestPty("xterm", 80, 24, ssh.TerminalModes{
		//ssh.ECHO: 0, // 关闭回显
	})
	if err != nil {
		//log.Fatalf("Failed to request pty: %v", err)
		return err
	}

	// 启动 shell
	err = session.Shell()
	if err != nil {
		//log.Fatalf("Failed to start shell: %v", err)
		return err
	}

	//从 SSH 读取数据并发送到 WebSocket
	go func() {
		buf := make([]byte, 1024)
		reader := io.MultiReader(sessionStdOut, sessionStdErr)
		for {
			n, err := reader.Read(buf)
			if err != nil {
				if err == io.EOF {
					break
				}
				continue
			}
			sshChan <- MessageSsh{
				msg:       buf[:n],
				serClient: s,
			}
		}
	}()

	return nil
}

func (s *ClientSession) write(message *[]byte) error {
	_, err := s.stdin.Write(*message)
	return err
}

func (s *ClientSession) setCoon(conn *websocket.Conn) {
	s.conn = conn
}

func (s *ClientSession) WindowChange(w, h int) error {
	return s.session.WindowChange(h, w)
}
