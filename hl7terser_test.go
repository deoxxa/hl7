package hl7terser

import (
	"testing"
)

type pair struct {
	s string
	q Query
}

var cases = []pair{
	pair{"MSH", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x0,
		FieldRepeat:        0x0,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH(1)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		Field:              0x0,
		FieldRepeat:        0x0,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH(1)-2", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		Field:              0x2,
		FieldRepeat:        0x0,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH(1)-2(3)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		Field:              0x2,
		FieldRepeat:        0x3,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH(1)-2(3)-4", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		Field:              0x2,
		FieldRepeat:        0x3,
		Component:          0x4,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH(1)-2(3)-4(5)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		Field:              0x2,
		FieldRepeat:        0x3,
		Component:          0x4,
		ComponentRepeat:    0x5,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH(1)-2(3)-4(5)-6", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		Field:              0x2,
		FieldRepeat:        0x3,
		Component:          0x4,
		ComponentRepeat:    0x5,
		SubComponent:       0x6,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH(1)-2(3)-4(5)-6(7)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x1,
		Field:              0x2,
		FieldRepeat:        0x3,
		Component:          0x4,
		ComponentRepeat:    0x5,
		SubComponent:       0x6,
		SubComponentRepeat: 0x7,
	}},
	pair{"MSH-2", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		FieldRepeat:        0x0,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH-2(3)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		FieldRepeat:        0x3,
		Component:          0x1,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH-2(3)-4", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		FieldRepeat:        0x3,
		Component:          0x4,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH-2(3)-4(5)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		FieldRepeat:        0x3,
		Component:          0x4,
		ComponentRepeat:    0x5,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH-2(3)-4(5)-6", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		FieldRepeat:        0x3,
		Component:          0x4,
		ComponentRepeat:    0x5,
		SubComponent:       0x6,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH-2(3)-4(5)-6(7)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		FieldRepeat:        0x3,
		Component:          0x4,
		ComponentRepeat:    0x5,
		SubComponent:       0x6,
		SubComponentRepeat: 0x7,
	}},
	pair{"MSH-2-4", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		FieldRepeat:        0x0,
		Component:          0x4,
		ComponentRepeat:    0x0,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH-2-4(5)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		FieldRepeat:        0x0,
		Component:          0x4,
		ComponentRepeat:    0x5,
		SubComponent:       0x1,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH-2-4(5)-6", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		FieldRepeat:        0x0,
		Component:          0x4,
		ComponentRepeat:    0x5,
		SubComponent:       0x6,
		SubComponentRepeat: 0x0,
	}},
	pair{"MSH-2-4(5)-6(7)", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		FieldRepeat:        0x0,
		Component:          0x4,
		ComponentRepeat:    0x5,
		SubComponent:       0x6,
		SubComponentRepeat: 0x7,
	}},
	pair{"MSH-2-4-6", Query{
		Segment:            "MSH",
		SegmentRepeat:      0x0,
		Field:              0x2,
		FieldRepeat:        0x0,
		Component:          0x4,
		ComponentRepeat:    0x0,
		SubComponent:       0x6,
		SubComponentRepeat: 0x0,
	}},
}

func TestParse(t *testing.T) {
	for _, c := range cases {
		q, err := Parse(c.s)

		if err != nil {
			t.Error(err)
		}

		if *q != c.q {
			t.Errorf("[%q] expected %#v to be %#v", c.s, q, c.q)
		}
	}
}

func TestNew(t *testing.T) {
	q1 := New("MSH", 0, 0, 0, 0, 0, 0, 0)
	if q1.Segment != "MSH" {
		t.Errorf("segment should be MSH")
	}
	if q1.Component != 1 {
		t.Errorf("component should be 1")
	}
	if q1.SubComponent != 1 {
		t.Errorf("subcomponent should be 1")
	}
}
