package main

import (
	"github.com/dustin/go-broadcast"
)

type Message struct {
	UserID string
	RoomID string
	Text   string
}

type Listener struct {
	RoomID string
	Chan   chan interface{}
}

type Manager struct {
	roomChannels map[string]broadcast.Broadcaster
	open         chan *Listener
	close        chan *Listener
	delete       chan string
	messages     chan *Message
}

func NewRoomManager() *Manager {
	m := &Manager{
		roomChannels: make(map[string]broadcast.Broadcaster),
		open:         make(chan *Listener, 100),
		close:        make(chan *Listener, 100),
		delete:       make(chan string),
		messages:     make(chan *Message, 100),
	}

	go m.run()
	return m
}

func (m *Manager) run() {
	for {
		select {
		case listener := <-m.open:
			m.register(listener)
		case listener := <-m.close:
			m.unregister(listener)
		case rid := <-m.delete:
			m.deleteBroadcase(rid)
		case message := <-m.messages:
			m.room(message.RoomID).Submit(message.UserID + " : " + message.Text)
		}
	}
}
func (m *Manager) register(listener *Listener) {
	m.room(listener.RoomID).Register(listener.Chan)
}

func (m *Manager) unregister(listener *Listener) {
	m.room(listener.RoomID).Unregister(listener.Chan)
	close(listener.Chan)
}

func (m *Manager) deleteBroadcase(roomId string) {
	b, ok := m.roomChannels[roomId]
	if ok {
		b.Close()
		delete(m.roomChannels, roomId)
	}
}

func (m *Manager) room(roomId string) broadcast.Broadcaster {
	b, ok := m.roomChannels[roomId]
	if !ok {
		b = broadcast.NewBroadcaster(10)
		m.roomChannels[roomId] = b
	}
	return b
}

func (m *Manager) OpenListener(roomId string) chan interface{} {
	listener := make(chan interface{})
	m.open <- &Listener{
		RoomID: roomId,
		Chan:   listener,
	}
	return listener
}

func (m *Manager) CloseListener(roomId string, channel chan interface{}) {
	m.close <- &Listener{
		RoomID: roomId,
		Chan:   channel,
	}
}

func (m *Manager) DeleteListener(roomId string) {
	m.delete <- roomId
}

func (m *Manager) Submit(user_id, room_id, text string) {
	m.messages <- &Message{
		RoomID: room_id,
		UserID: user_id,
		Text:   text,
	}
}
