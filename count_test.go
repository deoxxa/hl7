package hl7terser

import (
	"testing"

	"github.com/kdar/health/hl7"
)

type countTestCase struct {
	q string
	c int
	m []byte
}

var countTestCases = []countTestCase{
	countTestCase{"MSH", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX", 30, []byte(longTestMessageContent)},
	countTestCase{"WWW", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(0)", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX(0)", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW(0)", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(1)", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(1)", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW(1)", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(30)", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(30)", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW(30)", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH-1", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX-1", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(0)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX(0)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW(0)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(1)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(1)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW(1)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(30)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(30)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW(30)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH-100", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX-100", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-100", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(0)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(0)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW(0)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(1)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(1)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW(1)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(30)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(30)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW(30)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH-1", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-1", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-2", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-3", 3, []byte(longTestMessageContent)},
	countTestCase{"PID-4", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-3(0)", 5, []byte(longTestMessageContent)},
	countTestCase{"PID-3(1)", 5, []byte(longTestMessageContent)},
	countTestCase{"PID-3(2)", 5, []byte(longTestMessageContent)},
	countTestCase{"PID-3(3)", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-3(0)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-3(1)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-3(2)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-3(3)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5(0)", 7, []byte(longTestMessageContent)},
	countTestCase{"PID-5(1)", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5(2)", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-5(0)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-5(1)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5-1(0)", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-5-1(1)", 0, []byte(longTestMessageContent)},
	countTestCase{"NK1(1)", 1, []byte(longTestMessageContent)},
	countTestCase{"NK1(2)", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-1", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-1-1", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-1-2-3", 0, []byte(longTestMessageContent)},
}

func TestCount(t *testing.T) {
	for _, c := range countTestCases {
		q, err := Parse(c.q)
		if err != nil {
			t.Error(err)
		}

		s, err := hl7.Unmarshal(c.m)
		if err != nil {
			t.Error(err)
		}

		l := q.Count(Message(s))

		if l != c.c {
			t.Errorf("expected count to be %d; was %d (%q)", c.c, l, c.q)
		}
	}
}
