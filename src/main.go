package main

import (
	"fmt"
	"time"

	"github.com/lucasmenendez/gop2p"
)

func main() {

	main := gop2p.InitNode(5001, false)

	main.OnConnection(func(_ gop2p.Peer) {
		fmt.Printf("[main handler] -> Connected\n")
	})

	main.OnMessage(func(msg []byte, _ gop2p.Peer) {
		fmt.Printf("[main handler] -> Message: %s\n", string(msg))
	})

	main.OnDisconnection(func(_ gop2p.Peer) {
		fmt.Printf("[main handler] -> Disconnected\n")
	})

	go func() {
		time.Sleep(time.Second)

		entry := main.Self
		node := gop2p.InitNode(5002, true)

		node.Connect(entry)
		time.Sleep(time.Second)
		node.Broadcast([]byte("Hello peers!"))
		time.Sleep(2 * time.Second)
		node.Disconnect()
	}()

	go func() {
		time.Sleep(time.Second)
		entry := main.Self

		node := gop2p.InitNode(5003, false)

		node.Connect(entry)
		time.Sleep(2 * time.Second)
		node.Disconnect()
	}()

	time.Sleep(6 * time.Second)
	main.Broadcast([]byte("Hello peers!"))
	time.Sleep(2 * time.Second)
	main.Disconnect()

}
