package hl7terser

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"fknsrs.biz/p/hl7parser"
)

type countTestCase struct {
	q string
	c int
	m []byte
}

var countTestCases = []countTestCase{
	countTestCase{"MSH", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX", 47, []byte(longTestMessageContent)},
	countTestCase{"WWW", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(1)", 16, []byte(longTestMessageContent)},
	countTestCase{"OBX(1)", 7, []byte(longTestMessageContent)},
	countTestCase{"WWW(1)", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(2)", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(2)", 6, []byte(longTestMessageContent)},
	countTestCase{"WWW(2)", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(30)", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(30)", 6, []byte(longTestMessageContent)},
	countTestCase{"WWW(30)", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH-1", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX-1", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(1)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"OBX(1)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW(1)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(2)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(2)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW(2)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(30)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(30)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"WWW(30)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH-100", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX-100", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-100", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(1)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(1)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW(1)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"MSH(2)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"OBX(2)-100", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW(2)-100", 0, []byte(longTestMessageContent)},
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
	countTestCase{"PID-3(1)", 5, []byte(longTestMessageContent)},
	countTestCase{"PID-3(2)", 5, []byte(longTestMessageContent)},
	countTestCase{"PID-3(3)", 5, []byte(longTestMessageContent)},
	countTestCase{"PID-3(4)", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-3(1)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-3(2)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-3(3)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-3(4)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5(1)", 7, []byte(longTestMessageContent)},
	countTestCase{"PID-5(2)", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5(3)", 0, []byte(longTestMessageContent)},
	countTestCase{"PID-5-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-5(1)-1", 1, []byte(longTestMessageContent)},
	countTestCase{"PID-5(2)-1", 0, []byte(longTestMessageContent)},
	countTestCase{"NK1(1)", 4, []byte(longTestMessageContent)},
	countTestCase{"NK1(2)", 7, []byte(longTestMessageContent)},
	countTestCase{"NK1(3)", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-1", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-1-1", 0, []byte(longTestMessageContent)},
	countTestCase{"WWW-1-2-3", 0, []byte(longTestMessageContent)},
}

func TestCount(t *testing.T) {
	a := assert.New(t)

	for _, c := range countTestCases {
		q, err := Parse(c.q)
		a.NoError(err)

		m, _, err := hl7parser.Parse(c.m)
		a.NoError(err)

		if a.NotNil(q) && a.NotNil(m) {
			l := q.Count(m)
			a.Equal(c.c, l, q.String())
		}
	}
}
