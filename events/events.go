package events

import (
	"context"
	"sync"
)

type Events[T any] struct {
	handlers map[string]*EventHandlersGroup[T]
	mutex    sync.RWMutex
}

func New[T any]() *Events[T] {
	return &Events[T]{
		handlers: make(map[string]*EventHandlersGroup[T]),
	}
}

func (h *Events[T]) Add(event string, handler func(T)) context.CancelFunc {
	h.mutex.Lock()
	defer h.mutex.Unlock()
	group, ex := h.handlers[event]
	if !ex {
		group = NewEventHandlersGroup[T]()
		h.handlers[event] = group
	}
	return group.Add(handler)
}

func (h *Events[T]) Emit(event string, value T) {
	h.mutex.RLock()
	defer h.mutex.RUnlock()
	group, ex := h.handlers[event]
	if !ex {
		return
	}
	group.Emit(value)
}
