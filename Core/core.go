package Core

import (
	"fmt"
	"sync"
	"time"

	"github.com/jepsen-io/maelstrom/demo/go"
)

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