package chain

import (
	"github.com/skeyic/reverseproxy/router"
	"log/slog"
	"net/http"
)

var (
	TheChain = &Chain{}
)

type filterNode struct {
	next    *filterNode
	handler Filter
}

type Chain struct {
	head   *filterNode
	cursor *filterNode
}

func (c *Chain) Add(handler Filter) {
	var (
		current *filterNode
	)

	current = &filterNode{
		next:    nil,
		handler: handler,
	}

	if c.cursor == nil {
		c.head = current
	} else {
		c.cursor.next = current
	}
	c.cursor = current
}

func (c *Chain) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	current := c.head
	for current != nil {
		err := current.handler.Handle(w, req)
		if err != nil {
			slog.Error("ServeHTTP failed", "filter", current.handler, "err", err)
			return
		}
		current = current.next
	}

	router.TheRouter.Do(w, req)
}
