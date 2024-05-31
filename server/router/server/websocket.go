package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"reflect"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type requestXterm struct {
	action  string
	data    string
	uid     string
	id      int
	session int
}

type respXterm struct {
	code  int
	error error
	data  interface{}
}

func respError(code int, err error) []byte {
	marshal, _ := json.Marshal(respXterm{
		code:  code,
		error: err,
	})

	return marshal
}

func respOk(code int, data any) []byte {
	marshal, _ := json.Marshal(respXterm{
		code: code,
		data: data,
	})

	return marshal
}

var w sync.WaitGroup

func WebSocket(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Failed to upgrade websocket:", err)
		return
	}
	defer conn.Close()

	sendChan := make(chan interface{})

	go func() {
		defer w.Done()
		for {
			msg, ok := <-sendChan
			if !ok {
				break
			}
			// 返回结果
			conn.WriteMessage(websocket.TextMessage, respOk(0, msg))
		}
	}()

	// 从 WebSocket 读取数据并发送到 SSH
	go func() {
		defer w.Done()
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				//conn.WriteMessage(websocket.TextMessage, respError(-1, err))
				break
			}

			// 解包
			msg := requestXterm{}
			json.Unmarshal(message, &msg)

			// 创建客户端
			serClient, err := manage.generateSerClient(msg.id, msg.session)
			if err != nil {
				conn.WriteMessage(websocket.TextMessage, respError(-1, err))
				return
			}
			serClient.sendChan = sendChan

			// 执行业务逻辑
			method := reflect.ValueOf(serClient).MethodByName(msg.action)
			if !method.IsValid() {
				conn.WriteMessage(websocket.TextMessage, respError(-1, fmt.Errorf("method not found: %s", msg.action)))
				return
			}
			inputs := make([]reflect.Value, 1)
			inputs[0] = reflect.ValueOf(msg.data)
			method.Call(inputs)
			if serClient.error() != nil {
				conn.WriteMessage(websocket.TextMessage, respError(-1, serClient.error()))
				return
			}

			// 写入
			result := serClient.result()
			if result != nil {
				sendChan <- result
			}
		}
	}()
	w.Add(2)
	w.Wait()

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
