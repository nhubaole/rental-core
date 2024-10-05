package main

import (
	"fmt"
	"log"
	"time"

	socketio_client "github.com/zhouhui8915/go-socket.io-client"
)

func main() {
	// Cấu hình sử dụng giao thức polling thay vì websocket
	opts := &socketio_client.Options{
		Transport: "polling",  // Sử dụng polling
	}

	client, err := socketio_client.NewClient("http://localhost:3000", opts)
	if err != nil {
		log.Fatalf("Failed to connect to Node.js server: %v", err)
	}

	// Lắng nghe sự kiện "connection"
	client.On("connection", func(msg interface{}) {
		fmt.Printf("Received from Node.js: %v\n", msg)
	})

	// Gửi sự kiện "connection"
	client.Emit("connection", "Hello from Go")

	// Giữ client chạy để nhận sự kiện từ server Node.js
	for {
		time.Sleep(1 * time.Second)
	}
}
