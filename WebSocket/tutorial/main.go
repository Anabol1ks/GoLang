package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var hub = NewHub()

func wsHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Ошибка при обновлении соединения:", err)
		return
	}

	client := &Client{conn: conn, send: make(chan []byte)}
	hub.register <- client

	go client.readMessage()
	go client.writeMessage()
}
func (c *Client) readMessage() {
	defer func() {
		hub.unregister <- c
		c.conn.Close()
	}()
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		hub.broadcast <- message
	}
}

func (c *Client) writeMessage() {
	defer c.conn.Close()
	for message := range c.send {
		if err := c.conn.WriteMessage(websocket.TextMessage, message); err != nil {
			break
		}
	}
}

func main() {
	go hub.Run()

	r := gin.Default()

	r.GET("/ws", wsHandler)

	fmt.Println("Сервер запущен на порту 8080")
	r.Run(":8080")
}
