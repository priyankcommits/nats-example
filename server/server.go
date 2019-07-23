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
	nc.Publish("foo", []byte("Hello World"))
	send()
	runtime.Goexit()
}

func send() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)
	nc, _ := nats.Connect(nats.DefaultURL)
	nc.Publish("foo", []byte(text))
	send()
}
