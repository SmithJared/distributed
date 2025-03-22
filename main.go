package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {
	n := maelstrom.NewNode()
	c := NewCore(n)

	c.Handle("echo", func(msg maelstrom.Message) error {
		// Unmarshal the message body as an loosely-typed map.
		var body map[string]any
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		// Update the message type to return back.
		body["type"] = "echo_ok"

		// Echo the original message back with the updated message type.
		return n.Reply(msg, body)
	})

	c.Handle("generate", c.generate)


	if err := n.Run(); err != nil {
		log.Fatal(err)
	}
}

type Core struct {
	ID string
	counter uint64
	lock    *sync.Mutex
	*maelstrom.Node
}

func NewCore(node *maelstrom.Node) *Core {
	nanoNow := time.Now().UnixNano()

	return &Core{
		ID:      fmt.Sprint(nanoNow),
		Node:    node,
		counter: 0,
		lock:    &sync.Mutex{},
	}
}

func (c *Core) generate(msg maelstrom.Message) error {
	// Unmarshal the message body as an loosely-typed map.
	var body map[string]any
	if err := json.Unmarshal(msg.Body, &body); err != nil {
		return err
	}

	// Update the message type to return back.
	body["type"] = "generate_ok"
	body["id"] = c.generateUUID()

	// Echo the original message back with the updated message type.
	return c.Reply(msg, body)
}

func (c *Core) generateUUID() string {
	c.lock.Lock()
	defer c.lock.Unlock()
	now := time.Now().UnixNano()
	c.counter++
	return fmt.Sprintf("%s-%d-%d", c.ID, c.counter, now)
}
