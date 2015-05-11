package Pigeon

type IEvent interface {
    Fire()
}

type Event struct {
    IEvent
    observer *Observer
}

func (e *Event) Fire() {
    cb := e.observer.callback[e]
    cb(e)
}

type Observer struct {
    callback map[*Event]func (*Event)
}

func (o *Observer) Observe (e *Event, cb func(*Event)) {
    if o.callback == nil {
        o.callback = make(map[*Event]func (*Event), 1)
    }
    o.callback[e] = cb
    e.observer = o
}