package stracer

import (
	"log"
	"time"
)

type Tracer struct {
	enabled bool
	tag     string
	id      string
	start   time.Time
	last    time.Time
}

func New(tag string, enable bool) *Tracer {
	now := time.Now()
	t := &Tracer{
		enabled: enable,
		tag:     tag,
		start:   now,
		last:    now,
	}

	return t
}

func (t *Tracer) Step(event string) {
	if t.enabled {
		now := time.Now()
		log.Printf("%s : %s : %s [%s]", t.tag, t.id, event, now.Sub(t.last))
		t.last = now
	}
}

func (t *Tracer) Done() {
	if t.enabled {
		now := time.Now()
		log.Printf("%s : %s : done [%s]", t.tag, t.id, now.Sub(t.start))
		t.last = now
	}
}

