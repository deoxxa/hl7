package hl7

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	ErrInvalidQuery = fmt.Errorf("can't parse query")
)

var (
	terserRegexp = regexp.MustCompile(`^([A-Z][A-Z0-9]+)(?:\(([0-9]{1,3})\))?(?:-([0-9]{1,3})(?:\(([0-9]{1,3})\))?(?:-([0-9]{1,3})(?:-([0-9]{1,3}))?)?)?$`)
)

func ParseQuery(s string) (*Query, error) {
	m := terserRegexp.FindStringSubmatch(s)
	if m == nil {
		return nil, ErrInvalidQuery
	}

	var q Query

	q.Segment = m[1]

	if m[2] != "" {
		n, _ := strconv.ParseInt(m[2], 10, 32)
		q.SegmentOffset = max(int(n)-1, 0)
		q.HasSegmentOffset = true
	}

	if m[3] != "" {
		n, _ := strconv.ParseInt(m[3], 10, 32)
		q.Field = max(int(n)-1, 0)
		q.HasField = true
	}

	if m[4] != "" {
		n, _ := strconv.ParseInt(m[4], 10, 32)
		q.FieldOffset = max(int(n)-1, 0)
		q.HasFieldOffset = true
	}

	if m[5] != "" {
		n, _ := strconv.ParseInt(m[5], 10, 32)
		q.Component = max(int(n)-1, 0)
		q.HasComponent = true
	}

	if m[6] != "" {
		n, _ := strconv.ParseInt(m[6], 10, 32)
		q.SubComponent = max(int(n)-1, 0)
		q.HasSubComponent = true
	}

	return &q, nil
}
