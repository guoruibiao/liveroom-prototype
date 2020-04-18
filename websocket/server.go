package main

import (
	"fmt"
	"net/http"
	
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	// 跨域
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
var clients = make(map[*websocket.Conn]string)

func main() {
	http.HandleFunc("/ws", WebSocketHandler)
	err := http.ListenAndServe("0.0.0.0:9998", nil)
	if err != nil {
		fmt.Println("websocket listen and serve is failed")
	}
}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("websocket connect is failed,err:", err)
	}
	defer c.Close()
	clients[c] = r.RemoteAddr
	
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("websocket read is failed")
			break
		}
		
		fmt.Println("recv websocket msg: ", string(message))
		
		for client, _ := range clients {
			if c == client {
				err = client.WriteMessage(mt, []byte("[我]:" + string(message)))
			}else{
				name := clients[c]
				err = client.WriteMessage(mt, []byte(name + ":" + string(message)))
			}
			if err != nil {
				fmt.Println("websocket write is failed")
				break
			}
		}
	}
}
