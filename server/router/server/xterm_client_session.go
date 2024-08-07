package server

import (
	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
	"io"
	"log"
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

	// 启动伪终端
	err = session.RequestPty("xterm", 80, 24, ssh.TerminalModes{
		//ssh.ECHO: 0, // 关闭回显
	})
	if err != nil {
		//log.Fatalf("Failed to request pty: %v", err)
		return err
	}

	return nil
}

func (s *ClientSession) write(message *[]byte) error {
	_, err := s.stdin.Write(*message)
	return err
}

func (s *ClientSession) close() error {
	err := s.session.Close()
	if err != nil && err != io.EOF {
		log.Println("close err: %v", err)
	}
	s.conn = nil
	s.stdin = nil
	s.stdout = nil
	s.stderr = nil
	manage.serMap[s.serverId].session[s.sessionId] = nil
	return err
}

func (s *ClientSession) run(conn *websocket.Conn) {
	s.conn = conn
	s.session.Stdin = &TransferStation{conn: conn}
	s.session.Stdout = &TransferStation{conn: conn}
	s.session.Stderr = &TransferStation{conn: conn}
	//启动 shell
	err := s.session.Shell()
	if err != nil {
		log.Println("start shell: %v", err)
		s.close()
		return
	}

	err = s.session.Wait()
	if err != nil {
		if err2, ok := err.(*ssh.ExitError); !ok || err2.Waitmsg.ExitStatus() != ssh.TTY_OP_OSPEED {
			log.Println("wait shell: %v", err)
		}
		s.close()
	}
}

func (s *ClientSession) WindowChange(w, h int) error {
	return s.session.WindowChange(h, w)
}
