package Core

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/jepsen-io/maelstrom/demo/go"
)

func (c *Core) Generate(msg maelstrom.Message) error {
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