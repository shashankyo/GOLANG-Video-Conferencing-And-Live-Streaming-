package handlers

import (
	"videochat/pkg/chat"
	w "videochat/pkg/webrtc"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/net/websocket"
)

func RoomChat(c *fiber.Ctx) error {
	return c.Render("chat", fiber.Map{}, "layout/main")

}
func RoomChatWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}
	w.RoomLock.Lock()
	room := w.Rooms[uuid]
	w.RoomsLock.Unlock()
	if room == nil {
		return
	}
	if room.Hub == nil {
		return
	}
	chat.PeerChatConn(c.Conn, room.Hub)
}

func StreamChatWebSocket(c *websocket.Conn) {
	uuid := c.Params("suuid")
	if suuid == "" {
		return
	}
	w.RoomsLock.Lock()

	if stream, ok := w.Streams[suuid]; ok {
		w.RoomsLock.Unlock()

		if stream.Hub == nil {
			hub := chat.NewHub()
			stream.Hub = hub
			go hub.Run()
		}
		chat.PeerChatConn(c.Conn, stream.Hub)
		return
	}
	w.RoomsLock.Unlock()
}
