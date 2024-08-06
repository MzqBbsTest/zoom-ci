package server

import (
	"github.com/gorilla/websocket"
	"log"
)

type TransferStation struct {
	conn *websocket.Conn
}

func (t *TransferStation) Read(buf []byte) (n int, err error) {
	n, buf2, err := t.conn.ReadMessage()
	if err != nil {
		log.Println("read error :", err.Error())
		return 0, err
	}
	copy(buf, buf2)
	return n, nil
}

func (t *TransferStation) Write(buf []byte) (n int, err error) {
	err = t.conn.WriteMessage(websocket.BinaryMessage, buf)
	if err != nil {
		log.Println("write error :", err.Error())
		return 0, err
	}
	return len(buf), nil
}
