package hl7terser

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	ErrInvalidParse = fmt.Errorf("can't parse query")
)

var (
	terserRegexp = regexp.MustCompile(`^([A-Z][A-Z0-9]+)(?:\(([0-9]{1,3})\))?(?:-([0-9]{1,3})(?:\(([0-9]{1,3})\))?(?:-([0-9]{1,3})(?:\(([0-9]{1,3})\))?(?:-([0-9]{1,3})(?:\(([0-9]{1,3})\))?)?)?)?$`)
)

func Parse(s string) (*Query, error) {
	m := terserRegexp.FindStringSubmatch(s)
	if m == nil {
		return nil, ErrInvalidParse
	}

	var q Query

	q.Segment = m[1]

	if m[2] != "" {
		n, _ := strconv.ParseInt(m[2], 10, 32)
		q.SegmentRepeat = int(n)
	}

	if m[3] != "" {
		n, _ := strconv.ParseInt(m[3], 10, 32)
		q.Field = int(n)
	}

	if m[4] != "" {
		n, _ := strconv.ParseInt(m[4], 10, 32)
		q.FieldRepeat = int(n)
	}

	if m[5] != "" {
		n, _ := strconv.ParseInt(m[5], 10, 32)
		q.Component = int(n)
	}
	if q.Component == 0 {
		q.Component = 1
	}

	if m[6] != "" {
		n, _ := strconv.ParseInt(m[6], 10, 32)
		q.ComponentRepeat = int(n)
	}

	if m[7] != "" {
		n, _ := strconv.ParseInt(m[7], 10, 32)
		q.SubComponent = int(n)
	}
	if q.SubComponent == 0 {
		q.SubComponent = 1
	}

	if m[8] != "" {
		n, _ := strconv.ParseInt(m[8], 10, 32)
		q.SubComponentRepeat = int(n)
	}

	return &q, nil
}
