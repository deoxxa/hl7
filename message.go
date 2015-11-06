package hl7terser

import (
	"github.com/kdar/health/hl7"
)

type Message []hl7.Segment

func (m Message) Segments(name string) []hl7.Segment {
	var r []hl7.Segment

	for _, s := range m {
		v, ok := s.Index(0)
		if !ok {
			continue
		}

		if f, ok := v.(hl7.Field); ok {
			if f.String() == name {
				r = append(r, s)
			}
		}
	}

	return r
}

func (m Message) Segment(name string, index int) (hl7.Segment, bool) {
	segs := m.Segments(name)

	if len(segs) <= index {
		return nil, false
	}

	return segs[index], true
}
