package common

import (
	"fmt"
	"net/http"

	"github.com/pnutmath/gappa-tappa/backend/pkg/websocket"
)

// ServeWs serves the websocket
func ServeWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")

	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
