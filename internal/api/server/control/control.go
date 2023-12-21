package control

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"xcluster/internal/api/server"
	"xcluster/internal/hive"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServeServerControl(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logger.LogError(err)
		return
	}
	logger.Log(fmt.Sprintf("websocket established, remote=%s", conn.RemoteAddr()))
	var msg hive.Message
	if err = conn.ReadJSON(&msg); err != nil {
		if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
			logger.Log(fmt.Sprintf("websocket disconnected (abnormal), remote=%s", conn.RemoteAddr()))
		} else {
			logger.Log(fmt.Sprintf("enforcing disconnect from %s, websocket err=%s", conn.RemoteAddr(), err))
		}
		// conn maybe already closed, continue anyway
		_ = conn.Close()
		return
	}
	if msg.Operation != hive.OperationEngage {
		logger.Log(fmt.Sprintf("enforcing disconnect from %s, op=%s mismatch", conn.RemoteAddr(), msg.Operation))
		_ = conn.Close()
		return
	}
	data, ok := msg.Data.(map[string]interface{})
	if !ok {
		_ = conn.Close()
		logger.Log("data type error, expect to be a map")
		return
	}
	var boundInfo server.BoundInfo
	dataRaw, _ := json.Marshal(data)
	if err = json.Unmarshal(dataRaw, &boundInfo); err != nil {
		logger.Log(fmt.Sprintf("convert to boundInfo failed, err=%s", err))
		return
	}
	// auth
	s, err := boundInfo.ServerID.GetServer()
	if err != nil {
		logger.LogError(err)
		err = conn.WriteJSON(hive.MessageWithStatus{
			Message: hive.Message{
				Data: "server not found",
			},
			Status: false,
		})
		logger.LogIfError(err)
		_ = conn.Close()
		return
	}
	if !s.Secret.Compare(boundInfo.SecretRaw) {
		err = conn.WriteJSON(hive.MessageWithStatus{
			Message: hive.Message{
				Data: "auth failed",
			},
			Status: false,
		})
		logger.LogIfError(err)
		_ = conn.Close()
		return
	}
	err = conn.WriteJSON(hive.MessageWithStatus{Status: true})
	// add to hive
	hive.Central.Engage <- hive.NewBee(boundInfo.ServerID, conn)
	// keepalive by hive
	// do not close the connection
}
