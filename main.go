package main

import (
	"log"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	c := NewCore(n)

	c.Handle("echo", c.Echo)

	c.Handle("generate", c.Generate)


	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}
