package server

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/zoom-ci/zoom-ci/util/websocket"
	"golang.org/x/crypto/ssh"
	"sync"
)

//var upgrader = websocket.Upgrader{
//	ReadBufferSize:  1024,
//	WriteBufferSize: 1024,
//	CheckOrigin: func(r *http.Request) bool {
//		return true
//	},
//}

type requestXterm struct {
	Action  string
	Data    string
	Uid     string
	Id      int
	Session int
}

type respXterm struct {
	Code  int
	Error string
	Data  interface{}
}

func respError(code int, err error) []byte {
	marshal, _ := json.Marshal(respXterm{
		Code:  code,
		Error: err.Error(),
	})

	return marshal
}

func respOk(code int, data any) []byte {
	marshal, _ := json.Marshal(respXterm{
		Code: code,
		Data: data,
	})

	return marshal
}

var w sync.WaitGroup

type WebSocketQueryBind struct {
	Id        int `form:"id"`
	SessionId int `form:"session_id"`
}

func WebSocket(c *gin.Context) {
	websocket.Handler(func(ws *websocket.Conn) {
		//ctx := r.Context()

		var query WebSocketQueryBind
		if err := c.ShouldBindQuery(&query); err != nil {
			//render.ParamError(c, err.Error())
			return
		}

		serClient, err := manage.generateSerClient(query.Id, query.SessionId)
		if err != nil {
			//render.ParamError(c, err.Error())
			return
		}

		serClient.session.Stdin = ws
		serClient.session.Stderr = ws
		serClient.session.Stdout = ws

		// 启动伪终端
		err = serClient.session.RequestPty("xterm", 48, 95, ssh.TerminalModes{
			//ssh.ECHO: 0, // 关闭回显
		})
		if err != nil {
			//log.Fatalf("Failed to request pty: %v", err)
			return
		}

		// 启动 shell
		err = serClient.session.Shell()
		if err != nil {
			//log.Fatalf("Failed to start shell: %v", err)
			return
		}

		serClient.session.Wait()
		//sessionId := ws.Request().URL.Query().Get("session_id")
		//defer DeleteOnlineClient(sessionId)
		//w, err := strconv.Atoi(ws.Request().URL.Query().Get("w"))
		//if err != nil || (w < 40 || w > 8192) {
		//	_ = websocket.Message.Send(ws, "connect error window width !!!")
		//	DeleteOnlineClient(sessionId)
		//	return
		//}
		//h, err := strconv.Atoi(ws.Request().URL.Query().Get("h"))
		//if err != nil || (h < 2 || h > 4096) {
		//	_ = websocket.Message.Send(ws, "connect error window height !!!")
		//	DeleteOnlineClient(sessionId)
		//	return
		//}
		//
		//cli, ok := OnlineClients.Load(sessionId)
		//if !ok || cli == nil {
		//	_ = websocket.Message.Send(ws, "session_id not exists !!!")
		//	DeleteOnlineClient(sessionId)
		//	return
		//}
		//
		//conn, ok := cli.(*SshConn)
		//if !ok || conn == nil {
		//	_ = websocket.Message.Send(ws, "to ssh.Session error !!!")
		//	DeleteOnlineClient(sessionId)
		//	return
		//}
		//err = conn.RunTerminal(conn.Shell, ws, ws, ws, w, h, ws)
		//if err != nil {
		//	_ = websocket.Message.Send(ws, "connect error:"+err.Error())
		//	DeleteOnlineClient(sessionId)
		//	return
		//}
	}).ServeHTTP(c.Writer, c.Request)

	//conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	//if err != nil {
	//	log.Println("Failed to upgrade websocket:", err)
	//	return
	//}
	//defer conn.Close()
	//
	//serClient.session.Stdin = conn.NetConn()
	//serClient.session.Stderr = conn.NetConn()
	//serClient.session.Stdout = conn.NetConn()

	//select {

	//}
	//
	//
	//sendChan := make(chan interface{})
	//
	//
	//// 创建客户端
	//serClient.sendChan = sendChan
	//
	//go func() {
	//	defer w.Done()
	//	for {
	//		msg, ok := <-sendChan
	//		if !ok {
	//			break
	//		}
	//		// 返回结果
	//		conn.WriteMessage(websocket.TextMessage, respOk(0, msg))
	//	}
	//}()
	//
	//// 从 WebSocket 读取数据并发送到 SSH
	//go func() {
	//	defer w.Done()
	//	for {
	//		_, message, err := conn.ReadMessage()
	//		if err != nil {
	//			//conn.WriteMessage(websocket.TextMessage, respError(-1, err))
	//			break
	//		}
	//
	//		// 解包
	//		msg := requestXterm{}
	//		json.Unmarshal(message, &msg)
	//
	//		// 执行业务逻辑
	//		method := reflect.ValueOf(serClient).MethodByName(msg.Action)
	//		if !method.IsValid() {
	//			conn.WriteMessage(websocket.TextMessage, respError(-1, fmt.Errorf("method not found: %s", msg.Action)))
	//			return
	//		}
	//		inputs := make([]reflect.Value, 1)
	//		inputs[0] = reflect.ValueOf(msg.Data)
	//		method.Call(inputs)
	//		if serClient.error() != nil {
	//			conn.WriteMessage(websocket.TextMessage, respError(-1, serClient.error()))
	//			return
	//		}
	//
	//		// 写入
	//		result := serClient.result()
	//		if result != nil && !reflect.ValueOf(result).IsNil() {
	//			sendChan <- result
	//		}
	//	}
	//}()
	//w.Add(2)
	//w.Wait()

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
