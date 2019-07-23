package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"

	nats "github.com/nats-io/nats.go"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
		go reply()
	})
	runtime.Goexit()
}

func reply() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Publish("foo", []byte(text))
	fmt.Println(text)
}
