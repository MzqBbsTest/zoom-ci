package server

import (
	"errors"
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
	resp      *clientSessionResp
	sendChan  chan interface{}
}

type clientSessionResp struct {
	Action string
	Resp   map[string]interface{}
}

func (s *clientSession) setSessionStd(stdout, stderr io.Writer, stdin io.Reader) {
	s.session.Stdin = stdin
	s.session.Stdout = stdout
	s.session.Stderr = stderr
}

func (s *clientSession) error() error {
	return s.err
}

func (s *clientSession) result() interface{} {
	return s.resp
}

func (s *clientSession) Login(data any) error {
	s.err = nil
	if s.session != nil {
		s.resp = &clientSessionResp{
			Action: "LoginSuccess",
			Resp: map[string]interface{}{
				"server_id":  s.serverId,
				"session_id": s.sessionId,
			},
		}
		return nil
	}

	//id := data.(int)

	session, err := manage.serMap[s.serverId].client.NewSession()
	if err != nil {
		s.err = err
		return err
	}
	s.session = session

	// 获取会话的输入和输出
	sessionStdIn, err := session.StdinPipe()
	if err != nil {
		s.err = err
		return err
	}
	s.stdin = sessionStdIn

	sessionStdOut, err := session.StdoutPipe()
	if err != nil {
		s.err = err
		return err
	}
	s.stdout = sessionStdOut

	sessionStdErr, err := session.StderrPipe()
	if err != nil {
		s.err = err
		return err
	}
	s.stderr = sessionStdErr

	// 启动伪终端
	err = session.RequestPty("xterm", 80, 24, ssh.TerminalModes{
		ssh.ECHO: 0, // 关闭回显
	})
	if err != nil {
		s.err = err
		//log.Fatalf("Failed to request pty: %v", err)
		return err
	}

	// 启动 shell
	err = session.Shell()
	if err != nil {
		s.err = err
		//log.Fatalf("Failed to start shell: %v", err)
		return err
	}

	s.err = nil
	s.resp = &clientSessionResp{
		Action: "LoginSuccess",
		Resp: map[string]interface{}{
			"server_id":  s.serverId,
			"session_id": s.sessionId,
		},
	}

	//从 SSH 读取数据并发送到 WebSocket
	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := sessionStdOut.Read(buf)
			if err != nil {
				break
			}
			s.sendChan <- &clientSessionResp{
				Action: "CmdOut",
				Resp: map[string]interface{}{
					"data": buf[:n],
				},
			}
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
			s.sendChan <- &clientSessionResp{
				Action: "CmdErr",
				Resp: map[string]interface{}{
					"data": buf[:n],
				},
			}
			//conn.WriteMessage(websocket.TextMessage, buf[:n])
		}
	}()
	if manage.serMap[s.serverId].session == nil {
		manage.serMap[s.serverId].session = map[int]*clientSession{}
	}
	manage.serMap[s.serverId].session[s.sessionId] = s
	return nil
}

func (s *clientSession) Cmd(data any) error {
	if s.session == nil {
		s.err = errors.New("sessionId is nil")
		return s.err
	}
	message := data.(string)
	//output, err := s.session.CombinedOutput(message)
	//if err != nil {
	//	s.err = err
	//	return err
	//}
	//s.err = nil
	//s.resp = &clientSessionResp{
	//	Action: "CmdOut",
	//	Resp: map[string]interface{}{
	//		"data": output,
	//	},
	//}

	_, err := s.stdin.Write([]byte(message))
	if err != nil {
		return err
	}

	return nil
}
