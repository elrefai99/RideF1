package notifications

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

type Events struct {
	Type string `json:"type"`
	Data any    `json:"Data"`
}

type Hub struct {
	mux        sync.Mutex
	subscriber map[chan Events]struct{}
}

func NewHub() *Hub {
	return &Hub{
		subscriber: make(map[chan Events]struct{}),
	}
}

func (h *Hub) Publish(event Events) {
	h.mux.Lock()
	defer h.mux.Unlock()

	for sub := range h.subscriber {
		select {
		case sub <- event:
		default:
		}
	}
}

func (h *Hub) Stream(c *gin.Context) {
	flusher, ok := c.Writer.(http.Flusher)
	if !ok {
		c.Status(http.StatusInternalServerError)
		return
	}

	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")
	c.Header("X-Accel-Buffering", "no")

	subscriber := make(chan Events, 16)

	h.mux.Lock()
	h.subscriber[subscriber] = struct{}{}
	h.mux.Unlock()

	defer func() {
		h.mux.Lock()
		delete(h.subscriber, subscriber)
		close(subscriber)
		h.mux.Unlock()
	}()

	heartbeat := time.NewTicker(25 * time.Second)
	defer heartbeat.Stop()

	for {
		select {
		case <-c.Request.Context().Done():
			return

		case event := <-subscriber:
			data, err := json.Marshal(event.Data)
			if err != nil {
				continue
			}

			if _, err := fmt.Fprintf(
				c.Writer,
				"event: %s\ndata: %s\n\n",
				event.Type,
				data,
			); err != nil {
				return
			}
			flusher.Flush()

		case <-heartbeat.C:
			if _, err := fmt.Fprint(c.Writer, ": ping\n\n"); err != nil {
				return
			}
			flusher.Flush()
		}
	}
}
