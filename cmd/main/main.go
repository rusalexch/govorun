package main

import (
	"fmt"
	"govorun/internal/govorun"
	"time"
)

func main() {
	gov := govorun.Init("")

	go client1(gov)

	go update(gov)

	gov.Start()
}

func client1(gov *govorun.Govorun) {
	time.Sleep(5 * time.Second)
	ch := make(chan string)
	//defer close(ch)
	name := "1 channel"

	unsubscribe := gov.Subscribe(name, ch)
	go listen(ch, name)

	go closeSubs(10*time.Second, func() {
		unsubscribe()
		close(ch)
	})
}

func closeSubs(t time.Duration, un func()) {
	time.Sleep(t)
	un()
}

func update(gov *govorun.Govorun) {
	time.Sleep(7 * time.Second)

	gov.UpdateMessage("new message")
}

func listen(ch chan string, name string) {
	for msg := range ch {
		fmt.Println(name, msg)
	}
}
