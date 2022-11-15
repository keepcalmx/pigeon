package msg

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	const_ "github.com/keepcalmx/go-pigeon/common/constant"
	"github.com/keepcalmx/go-pigeon/common/utils"
	"github.com/keepcalmx/go-pigeon/model/ws"
)

const (
	CHAT  = "chat"
	USER  = "user"
	GROUP = "group"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// uuid of user who use this client
	uuid string

	// websocket connection.
	conn *websocket.Conn

	// buffer to write response
	buffer chan *ws.Response
}

func (c *Client) read() {
	defer func() {
		c.conn.Close()
	}()

	// basic settings
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(
		func(string) error {
			c.conn.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		},
	)

	for {
		msg := &ws.Msg{}
		err := c.conn.ReadJSON(msg)
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("unexpected read error: ", err)
			}
			c.hub.unregister <- c
			break
		}

		if msg.ToType == USER {
			c.hub.unicast <- msg
		}
		if msg.ToType == GROUP {
			c.hub.broadcast <- msg
		}
	}
}

func (c *Client) write() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case r, ok := <-c.buffer:
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			c.conn.WriteJSON(r)
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// Server handles websocket requests from the peer.
func HandleRequest(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("token")
	if token == "" {
		// w.Write([]byte("a valid token is required"))
		return
	}

	claims, err := utils.ParseToken(token)
	if err != nil {
		// w.Write([]byte("a valid token is required"))
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("upgrade http to websocket with error ", err)
		return
	}

	client := &Client{
		hub:    GetHub(),
		uuid:   claims.UUID,
		conn:   conn,
		buffer: make(chan *ws.Response, const_.SMALL_BUFFER_SIZE),
	}

	go client.read()
	go client.write()

	client.hub.register <- client
}
