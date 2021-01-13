package streamer

import (
	"sync"
	"time"

	"github.com/pkg/errors"
)

var (
	ErrTimeout = errors.New("send timeout")
)

type EventStreamer struct {
	mu          *sync.Mutex
	listeners   map[string]chan interface{}
	sendTimeout int
}

func NewEventStreamer(timeout int) *EventStreamer {
	return &EventStreamer{
		mu:          &sync.Mutex{},
		listeners:   map[string]chan interface{}{},
		sendTimeout: timeout,
	}
}

func (streamer *EventStreamer) Register(wsuuid string) chan interface{} {
	streamer.mu.Lock()
	defer streamer.mu.Unlock()
	updates := make(chan interface{})
	streamer.listeners[wsuuid] = updates
	return updates
}

func (streamer *EventStreamer) Unregister(wsuuid string) {
	streamer.mu.Lock()
	defer streamer.mu.Unlock()
	delete(streamer.listeners, wsuuid)
}

func (streamer *EventStreamer) SendTo(wsuuid string, data interface{}) error {
	streamer.mu.Lock()
	defer streamer.mu.Unlock()
	select {
	case streamer.listeners[wsuuid] <- data:
	case <-time.After(time.Second * time.Duration(streamer.sendTimeout)):
		return ErrTimeout
	}
	return nil
}

func (streamer *EventStreamer) SendToAll(data interface{}) error {
	streamer.mu.Lock()
	defer streamer.mu.Unlock()

	var err error
	for wsuuid, ch := range streamer.listeners {
		select {
		case ch <- data:
		case <-time.After(time.Second * time.Duration(streamer.sendTimeout)):
			err = errors.Wrap(err, wsuuid)
		}
	}
	return err
}
