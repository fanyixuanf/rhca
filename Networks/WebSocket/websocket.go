/*
@Time : 2023/5/1 17:06
@Author : Yuxue.fan<fanyixuanf+go@gmail.com>
@File : websocket.go
@Software: GoLand
@Description:
@Version: 1.0.0
*/
package WebSocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"net/url"
	"time"
)

func WebsocketClient() {
	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/echo"}
	fmt.Println("connecting to", u.String())

	// set up websocket connection
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("dial error:", err)
	}

	defer conn.Close()

	// write message to websocket connection
	message := []byte("hello, world!")
	err = conn.WriteMessage(websocket.TextMessage, message)
	if err != nil {
		fmt.Println("write message error:", err)
	}

	// read message from websocket connection
	_, p, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("read message error:", err)
	}

	fmt.Printf("received message: %s\n", p)

	// wait for some time and then close the connection
	time.Sleep(2 * time.Second)
	fmt.Println("closing the connection...")
}

func WebscoketServe() {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("connection established")
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("upgrade error:", err)
			return
		}

		defer conn.Close()

		// read message from websocket connection
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("read message error:", err)
			return
		}

		fmt.Printf("received message: %s\n", message)

		// send the message back to the client
		if err := conn.WriteMessage(messageType, message); err != nil {
			fmt.Println("write message error:", err)
			return
		}

		fmt.Println("message sent back to client")

		// wait for some time and then close the connection
		time.Sleep(2 * time.Second)
		fmt.Println("closing the connection...")
	})

	fmt.Println("listening on :8080")
	http.ListenAndServe(":8080", nil)
}