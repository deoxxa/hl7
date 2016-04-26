package hl7

import (
	"fmt"
)

type Query struct {
	Segment          string
	HasSegmentOffset bool
	SegmentOffset    int
	HasField         bool
	Field            int
	HasFieldOffset   bool
	FieldOffset      int
	HasComponent     bool
	Component        int
	HasSubComponent  bool
	SubComponent     int
}

func New(segment string, segmentOffset, field, fieldOffset, component, subComponent int) Query {
	return Query{
		Segment:       segment,
		SegmentOffset: max(segmentOffset-1, 0),
		Field:         max(field-1, 0),
		FieldOffset:   max(fieldOffset-1, 0),
		Component:     max(component-1, 0),
		SubComponent:  max(subComponent-1, 0),
	}
}

func (q Query) String() string {
	s := q.Segment

	if q.HasSegmentOffset {
		s += fmt.Sprintf("(%d)", q.SegmentOffset+1)
	}

	if !q.HasField {
		return s
	}

	s += fmt.Sprintf("-%d", q.Field+1)

	if q.HasFieldOffset {
		s += fmt.Sprintf("(%d)", q.FieldOffset+1)
	}

	if !q.HasComponent {
		return s
	}

	s += fmt.Sprintf("-%d", q.Component+1)

	if !q.HasSubComponent {
		return s
	}

	s += fmt.Sprintf("-%d", q.SubComponent+1)

	return s
}

func (q Query) GetString(m Message) string {
	s, _ := q.Get(m)
	return s
}

func (q Query) Get(m Message) (string, bool) {
	s := m.Segment(q.Segment, q.SegmentOffset)

	if len(s) <= q.Field+1 {
		return "", false
	}
	f := s[q.Field+1]

	if len(f) <= q.FieldOffset {
		return "", false
	}
	fi := f[q.FieldOffset]

	if len(fi) <= q.Component {
		return "", false
	}
	c := fi[q.Component]

	if len(c) <= q.SubComponent {
		return "", false
	}

	return string(c[q.SubComponent]), true
}

func (q Query) Count(m Message) int {
	if !q.HasSegmentOffset && !q.HasField {
		return len(m.Segments(q.Segment))
	}

	s := m.Segment(q.Segment, q.SegmentOffset)
	if !q.HasField {
		return len(s)
	}

	if len(s) <= q.Field+1 {
		return 0
	}
	f := s[q.Field+1]
	if !q.HasFieldOffset && !q.HasComponent {
		return len(f)
	}

	if len(f) <= q.FieldOffset {
		return 0
	}
	fi := f[q.FieldOffset]
	if !q.HasComponent {
		return len(fi)
	}

	if len(fi) <= q.Component {
		return 0
	}
	c := fi[q.Component]
	if !q.HasSubComponent {
		return len(c)
	}

	if len(c) <= q.SubComponent {
		return 0
	}

	return 1
}
