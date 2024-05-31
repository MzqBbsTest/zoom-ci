package server

import (
	"golang.org/x/crypto/ssh"
	"io"
)

type clientSession struct {
	serverId  int // 服务器id
	sessionId int // m.serMap[id].session下标 用来寻址 m.serMap[id].session
	session   *ssh.Session
	stdin     io.WriteCloser
	stdout    io.Reader
	stderr    io.Reader
	err       error
	resp      map[string]interface{}
	sendChan  chan interface{}
}

func (s *clientSession) error() error {
	return s.err
}

func (s *clientSession) result() interface{} {

	return nil
}

func (s *clientSession) login(id int) error {
	if s.session != nil {
		s.err = nil
		s.resp = map[string]interface{}{
			"server_id":  s.serverId,
			"session_id": s.sessionId,
		}
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
		ssh.ECHO: 0, // 关闭回显
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

	s.err = nil
	s.resp = map[string]interface{}{
		"server_id":  s.serverId,
		"session_id": s.sessionId,
	}

	//从 SSH 读取数据并发送到 WebSocket
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := sessionStdOut.Read(buf)
			if err != nil {
				break
			}
			s.sendChan <- buf[:n]
			//conn.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := sessionStdErr.Read(buf)
			if err != nil {
				break
			}
			s.sendChan <- buf[:n]
			//conn.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()

	return nil
}
