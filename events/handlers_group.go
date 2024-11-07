package events

import (
	"context"
	"sync"

	"github.com/google/uuid"
)

type EventHandlersGroup[T any] struct {
	handlers map[string]func(T)
	mutex    sync.Mutex
}

func NewEventHandlersGroup[T any]() *EventHandlersGroup[T] {
	return &EventHandlersGroup[T]{
		handlers: make(map[string]func(T)),
	}
}

func (h *EventHandlersGroup[T]) Add(handler func(T)) context.CancelFunc {
	id := uuid.NewString()
	h.mutex.Lock()
	h.handlers[id] = handler
	h.mutex.Unlock()

	return func() {
		h.mutex.Lock()
		delete(h.handlers, id)
		h.mutex.Unlock()
	}
}
func (h *EventHandlersGroup[T]) Clear() {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	for k, _ := range h.handlers {
		delete(h.handlers, k)
	}
}

func (h *EventHandlersGroup[T]) Emit(value T) {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	for _, v := range h.handlers {
		v(value)
	}
}
