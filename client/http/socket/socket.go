package socket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/vsayfb/e-commerce-scrapper/client/channel"
	"net/http"
	"strconv"
	"strings"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ConcurrentSearchHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(conn *websocket.Conn) {
		connErr := conn.Close()
		if connErr != nil {
			fmt.Println(connErr)
		}
	}(conn)

	for {
		_, msg, readErr := conn.ReadMessage()

		if readErr != nil {
			fmt.Println(readErr)
			return
		}

		query := string(msg)

		query = strings.TrimSpace(query)

		split := strings.Split(query, "-")

		page, uintErr := strconv.ParseUint(split[1], 10, 8)

		if uintErr != nil {
			fmt.Println(uintErr)
			return
		}

		channel.ReceiveProducts(split[0], uint8(page), conn)
	}
}
