package hl7 // import "fknsrs.biz/p/hl7"

import (
	"github.com/facebookgo/stackerr"
)

type ErrInvalidQuery error

func ParseQuery(s string) (*Query, error) {
	var q Query

	var offset int

	if err := parseQueryHeader(s, &offset, &q.Segment); err != nil {
		return nil, err
	}
	if offset == len(s) {
		return &q, nil
	}

	if err := parseQueryParen(s, &offset, &q.HasSegmentOffset, &q.SegmentOffset); err != nil {
		return nil, err
	}
	if offset == len(s) {
		return &q, nil
	}

	if err := parseQueryNumber(s, &offset, &q.HasField, &q.Field); err != nil {
		return nil, err
	}
	if offset == len(s) {
		return &q, nil
	}

	if err := parseQueryParen(s, &offset, &q.HasFieldOffset, &q.FieldOffset); err != nil {
		return nil, err
	}
	if offset == len(s) {
		return &q, nil
	}

	if err := parseQueryNumber(s, &offset, &q.HasComponent, &q.Component); err != nil {
		return nil, err
	}
	if offset == len(s) {
		return &q, nil
	}

	if err := parseQueryNumber(s, &offset, &q.HasSubComponent, &q.SubComponent); err != nil {
		return nil, err
	}
	if offset == len(s) {
		return &q, nil
	}

	if offset != len(s) {
		return nil, stackerr.Newf("junk data found at position %d", offset)
	}

	return &q, nil
}

func parseQueryHeader(s string, o *int, v *string) error {
	b := make([]byte, 0, 3)

	var e int

	for e = *o; e < len(s); e++ {
		if !(s[e] >= 'A' && s[e] <= 'Z') && !(s[e] >= '0' && s[e] <= '9') {
			break
		}

		b = append(b, s[e])
	}

	*v = string(b)
	*o = e

	return nil
}

func parseQueryParen(s string, o *int, b *bool, v *int) error {
	if s[*o] != '(' {
		return nil
	}

	var e int
	var n int

loop:
	for e = *o + 1; e < len(s); e++ {
		switch {
		case s[e] >= '0' && s[e] <= '9':
			n = (n * 10) + int(s[e]-'0')
		case s[e] == ')':
			break loop
		default:
			return stackerr.Newf("invalid byte (%q) found at offset %d", s[e], e)
		}
	}

	*o = e + 1
	*b = true
	*v = max(n-1, 0)

	return nil
}

func parseQueryNumber(s string, o *int, b *bool, v *int) error {
	if s[*o] != '-' {
		return nil
	}

	var e int
	var n int

loop:
	for e = *o + 1; e < len(s); e++ {
		switch {
		case s[e] >= '0' && s[e] <= '9':
			n = (n * 10) + int(s[e]-'0')
		default:
			break loop
		}
	}

	*o = e
	*b = true
	*v = max(n-1, 0)

	return nil
}
