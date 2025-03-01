package event

import (
	"github.com/lazygophers/utils/routine"
	"github.com/lazygophers/utils/runtime"
	"sync"
)

var defaultManager = NewManager()

type EventHandler func(args any)

type eventItem struct {
	handler EventHandler

	async bool
}

type Manager struct {
	eventMux sync.RWMutex
	events   map[string][]*eventItem
	c        chan *emitItem
}

func (p *Manager) register(eventName string, item *eventItem) {
	p.eventMux.Lock()
	defer p.eventMux.Unlock()

	p.events[eventName] = append(p.events[eventName], item)
}

func Register(eventName string, handler EventHandler) {
	defaultManager.Register(eventName, handler)
}

func (p *Manager) Register(eventName string, handler EventHandler) {
	p.register(eventName, &eventItem{
		handler: handler,
	})
}

func RegisterAsync(eventName string, handler EventHandler) {
	defaultManager.RegisterAsync(eventName, handler)
}

func (p *Manager) RegisterAsync(eventName string, handler EventHandler) {
	p.register(eventName, &eventItem{
		handler: handler,
		async:   true,
	})
}

func (p *Manager) getItems(eventName string) []*eventItem {
	p.eventMux.RLock()
	defer p.eventMux.RUnlock()

	return p.events[eventName]
}

func Emit(eventName string, args any) {
	defaultManager.Emit(eventName, args)
}

type emitItem struct {
	handler EventHandler
	args    any
}

func (p *emitItem) do() {
	defer runtime.CachePanic()

	p.handler(p.args)
}

func (p *Manager) Emit(eventName string, args any) {
	for _, event := range p.getItems(eventName) {
		if event.async {
			p.c <- &emitItem{
				handler: event.handler,
				args:    args,
			}
			continue
		}

		event.handler(args)
	}
}

func NewManager() *Manager {
	p := &Manager{
		events: make(map[string][]*eventItem),

		c: make(chan *emitItem, 10),
	}

	routine.GoWithRecover(func() (err error) {
		for item := range p.c {
			item.do()
		}
		return nil
	})

	return p
}
