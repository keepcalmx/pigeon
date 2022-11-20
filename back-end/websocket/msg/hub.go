package msg

import (
	"log"
	"sync"
	"time"

	"github.com/keepcalmx/go-pigeon/cache"
	C "github.com/keepcalmx/go-pigeon/common/constant"
	"github.com/keepcalmx/go-pigeon/ent"
	"github.com/keepcalmx/go-pigeon/model/ws"
	"github.com/keepcalmx/go-pigeon/storage"
)

var (
	once sync.Once
	hub  *Hub
)

type Hub struct {

	// user client map(thread safe)
	clients *sync.Map

	// group message broadcast channel
	broadcast chan *ws.Msg

	// private message unicast channel
	unicast chan *ws.Msg

	// status push channel
	status chan *ws.Status

	// client register channel
	register chan *Client

	// client unregister channel
	unregister chan *Client
}

// GetHub returns a singleton hub
func GetHub() *Hub {
	once.Do(func() {
		hub = &Hub{
			clients:    &sync.Map{},
			broadcast:  make(chan *ws.Msg, C.LARGE_BUFFER_SIZE),
			unicast:    make(chan *ws.Msg, C.MID_BUFFER_SIZE),
			status:     make(chan *ws.Status, C.MID_BUFFER_SIZE),
			register:   make(chan *Client, C.MID_BUFFER_SIZE),
			unregister: make(chan *Client, C.MID_BUFFER_SIZE),
		}
		log.Println("hub is initialized...")
	})
	return hub
}

func (h *Hub) Run() {
	log.Println("hub is now running...")
	// use buffered channels to write private and group messages
	groupMsgBuffer := make(chan *ent.GroupMsg, C.LARGE_BUFFER_SIZE)
	privateMsgBuffer := make(chan *ent.PrivateMsg, C.MID_BUFFER_SIZE)
	go func() {
		for {
			select {
			case msg := <-groupMsgBuffer:
				msg_, err := storage.CreateGroupMsg(msg)
				cache.AddGroupMsg(msg_)
				if err != nil {
					log.Println("create group message with error ", err)
				}
			case msg := <-privateMsgBuffer:
				msg_, err := storage.CreatePrivateMsg(msg)
				cache.AddPrivateMsg(msg_)
				if err != nil {
					log.Println("create private message with error ", err)
				}
			}
		}
	}()

	for {
		select {
		case c := <-h.register:
			h.clients.Store(c.uuid, c)
			h.status <- &ws.Status{
				UUID:   c.uuid,
				Target: "online",
				Value:  true,
			}
		case c := <-h.unregister:
			if _, ok := h.clients.Load(c.uuid); ok {
				close(c.buffer)
				h.clients.Delete(c.uuid)
			}
			h.status <- &ws.Status{
				UUID:   c.uuid,
				Target: "online",
				Value:  false,
			}
		case status := <-h.status:
			h.clients.Range(func(key, value interface{}) bool {
				if key == status.UUID {
					return true
				}
				value.(*Client).buffer <- &ws.Response{
					Type: C.STATUS_TYPE,
					Data: status,
				}
				return true
			})
		case msg := <-h.unicast:
			privateMsgBuffer <- &ent.PrivateMsg{
				From:      msg.From,
				To:        msg.To,
				Type:      msg.Type,
				Content:   msg.Content,
				Read:      false,
				CreatedAt: time.Now(),
			}

			client, ok := h.clients.Load(msg.To)
			if !ok {
				// message receiver is offline
				continue
			}
			response := &ws.Response{
				Type: C.MESSAGE_TYPE,
				Data: msg,
			}
			select {
			case client.(*Client).buffer <- response:
			default:
				h.unregister <- client.(*Client)
			}
		case msg := <-h.broadcast:
			groupMsgBuffer <- &ent.GroupMsg{
				From:      msg.From,
				To:        msg.To,
				Type:      msg.Type,
				Content:   msg.Content,
				CreatedAt: time.Now(),
			}

			members, _ := storage.ListGroupUsers(msg.To)
			for _, member := range members {
				if member.UUID == msg.From {
					continue
				}
				client, ok := h.clients.Load(member.UUID)
				if !ok {
					// member is offline
					continue
				}
				response := &ws.Response{
					Type: C.MESSAGE_TYPE,
					Data: msg,
				}
				select {
				case client.(*Client).buffer <- response:
				default:
					h.unregister <- client.(*Client)
				}
			}
		}
	}
}

func (h *Hub) IsOnline(uuid string) bool {
	_, ok := h.clients.Load(uuid)
	return ok
}
