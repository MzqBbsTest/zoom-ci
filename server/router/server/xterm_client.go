package server

import (
	"fmt"
	"github.com/zoom-ci/zoom-ci/server/module/server"
	"golang.org/x/crypto/ssh"
	"sync"
)

type Manage struct {
	serMap map[int]*client
}

func (m *Manage) generateSerClient(id, session int) (*clientSession, error) {
	if m.serMap[id] != nil {
		return m.generateSession(id, session), nil
	}

	ser := &server.Server{
		ID: id,
	}
	if err := ser.Detail(); err != nil {
		return nil, err
	}
	m.serMap[id] = &client{
		ser: ser,
	}
	m.serMap[id].session = map[int]*clientSession{}

	if err := ser.Detail(); err != nil {
		return nil, err
	}

	if m.serMap[id].isLogin() == false {
		if err := m.serMap[id].dial(); err != nil {
			return nil, err
		}
	}

	return m.generateSession(id, session), nil
}

// 創建一個
func (m *Manage) createNewSessionId(id int) int {
	// 加锁
	mutex := sync.Mutex{}
	mutex.Lock()
	index := len(m.serMap[id].session) + 1
	mutex.Unlock()
	return index
}

func (m *Manage) generateSession(id int, sessionId int) *clientSession {
	if sessionId == 0 {
		c := &clientSession{
			serverId:  id,
			sessionId: m.createNewSessionId(id),
		}
		m.serMap[id].session[c.sessionId] = c
		return c
	}
	c := m.serMap[id].session[sessionId]
	c.login(id)
	return c
}

type client struct {
	ser     *server.Server
	client  *ssh.Client
	session map[int]*clientSession
	err     error
}

func (s *client) isLogin() bool {
	return s.client != nil
}

func (s *client) dial() error {

	config := &ssh.ClientConfig{
		User: s.ser.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(s.ser.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(), // 仅用于测试，生产环境应使用安全的 HostKeyCallback
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", s.ser.Ip, s.ser.SSHPort), config)
	if err != nil {
		return err
	}
	s.client = client

	return nil
}

func (s *client) error() error {
	return s.err
}

func (s *client) result() interface{} {

	return nil
}
