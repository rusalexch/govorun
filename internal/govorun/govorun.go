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

func (gov *Govorun) Start() {
	tick := time.Tick(1 * time.Second)
	go gov.startServer()
	log.Print("govorun was started")

	for range tick {
		for _, ch := range gov.ch {
			if ch != nil {
				ch <- gov.message
			}
		}
	}
}

func (gov *Govorun) Subscribe(name string, ch chan string) func() {
	gov.ch[name] = ch
	log.Printf("govorun: channel %s was subscribe", name)

	return func() {
		gov.Unsubscribe(name)
	}
}

func (gov *Govorun) Unsubscribe(name string) {
	log.Printf("govorun: channel %s was unsubscribe", name)
	delete(gov.ch, name)
}

func (gov *Govorun) UpdateMessage(msg string) {
	log.Printf("govorun: message was changed (%s)", msg)
	gov.message = msg
}
