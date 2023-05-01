/*
@Time : 2023/5/1 17:14
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : websocket_test
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package WebSocket

import (
	"github.com/gorilla/websocket"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWebSocket(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}

		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			t.Fatalf("upgrade error: %v", err)
		}

		defer conn.Close()

		// read message from websocket connection
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			t.Fatalf("read message error: %v", err)
		}

		if string(message) != "hello, world!" {
			t.Fatalf("received unexpected message: %s", string(message))
		}

		// send the message back to the client
		if err := conn.WriteMessage(messageType, message); err != nil {
			t.Fatalf("write message error: %v", err)
		}
	}))

	defer server.Close()

	u := "ws" + strings.TrimPrefix(server.URL, "http") + "/echo"
	conn, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		t.Fatalf("dial error: %v", err)
	}

	defer conn.Close()

	// write message to websocket connection
	message := []byte("hello, world!")
	if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
		t.Fatalf("write message error: %v", err)
	}

	// read message from websocket connection
	_, p, err := conn.ReadMessage()
	if err != nil {
		t.Fatalf("read message error: %v", err)
	}

	if string(p) != "hello, world!" {
		t.Fatalf("received unexpected message: %s", string(p))
	}

}

func TestWebscoketServe(t *testing.T) {
	WebscoketServe()
}

func TestWebsocketClient(t *testing.T) {
	WebsocketClient()
}
