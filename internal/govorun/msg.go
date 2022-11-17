package govorun

import (
	"log"
	"time"
)

type Govorun struct {
	message string
	ch      map[string]chan string
}

func Init(msg string) *Govorun {
	var message string
	if msg == "" {
		message = "default message"
	}
	log.Println("govorun was init")

	return &Govorun{
		message: message,
		ch:      make(map[string]chan string),
	}
}

func (m *Govorun) Start() {
	tick := time.Tick(1 * time.Second)
	log.Print("govorun was started")

	for range tick {
		for _, ch := range m.ch {
			if ch != nil {
				ch <- m.message
			}
		}
	}
}

func (m *Govorun) Subscribe(name string, ch chan string) func() {
	m.ch[name] = ch
	log.Printf("govorun: channel %s was subscribe", name)

	return func() {
		m.Unsubscribe(name)
	}
}

func (m *Govorun) Unsubscribe(name string) {
	log.Printf("govorun: channel %s was unsubscribe", name)
	delete(m.ch, name)
}

func (m *Govorun) UpdateMessage(msg string) {
	log.Printf("govorun: message was changed (%s)", msg)
	m.message = msg
}
