package hl7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type parserTestPair struct {
	s string
	q Query
}

var parserTestCases = []parserTestPair{
	parserTestPair{"MSH", Query{
		Segment:       "MSH",
		SegmentOffset: 0,
		Field:         0,
		FieldOffset:   0,
		Component:     0,
		SubComponent:  0,
	}},
	parserTestPair{"MSH(1)", Query{
		Segment:          "MSH",
		SegmentOffset:    0,
		HasSegmentOffset: true,
		Field:            0,
		FieldOffset:      0,
		Component:        0,
		SubComponent:     0,
	}},
	parserTestPair{"MSH(1)-2", Query{
		Segment:          "MSH",
		SegmentOffset:    0,
		HasSegmentOffset: true,
		Field:            1,
		HasField:         true,
		FieldOffset:      0,
		Component:        0,
		SubComponent:     0,
	}},
	parserTestPair{"MSH(1)-2(3)", Query{
		Segment:          "MSH",
		SegmentOffset:    0,
		HasSegmentOffset: true,
		Field:            1,
		HasField:         true,
		FieldOffset:      2,
		HasFieldOffset:   true,
		Component:        0,
		SubComponent:     0,
	}},
	parserTestPair{"MSH(1)-2(3)-4", Query{
		Segment:          "MSH",
		SegmentOffset:    0,
		HasSegmentOffset: true,
		Field:            1,
		HasField:         true,
		FieldOffset:      2,
		HasFieldOffset:   true,
		Component:        3,
		HasComponent:     true,
		SubComponent:     0,
	}},
	parserTestPair{"MSH-2", Query{
		Segment:       "MSH",
		SegmentOffset: 0,
		Field:         1,
		HasField:      true,
		FieldOffset:   0,
		Component:     0,
		SubComponent:  0,
	}},
	parserTestPair{"MSH-2(3)", Query{
		Segment:        "MSH",
		SegmentOffset:  0,
		Field:          1,
		HasField:       true,
		FieldOffset:    2,
		HasFieldOffset: true,
		Component:      0,
		SubComponent:   0,
	}},
	parserTestPair{"MSH-2(3)-4", Query{
		Segment:        "MSH",
		SegmentOffset:  0,
		Field:          1,
		HasField:       true,
		FieldOffset:    2,
		HasFieldOffset: true,
		Component:      3,
		HasComponent:   true,
		SubComponent:   0,
	}},
	parserTestPair{"MSH-2-4", Query{
		Segment:       "MSH",
		SegmentOffset: 0,
		Field:         1,
		HasField:      true,
		FieldOffset:   0,
		Component:     3,
		HasComponent:  true,
		SubComponent:  0,
	}},
	parserTestPair{"MSH-2-4-6", Query{
		Segment:         "MSH",
		SegmentOffset:   0,
		Field:           1,
		HasField:        true,
		FieldOffset:     0,
		Component:       3,
		HasComponent:    true,
		SubComponent:    5,
		HasSubComponent: true,
	}},
}

func TestParseQuery(t *testing.T) {
	a := assert.New(t)

	for _, c := range parserTestCases {
		q, err := ParseQuery(c.s)
		a.NoError(err)

		if a.NotNil(q) {
			a.Equal(*q, c.q, c.s)
		}
	}
}
