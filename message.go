package hl7 // import "fknsrs.biz/p/hl7"

import (
	"github.com/facebookgo/stackerr"
)

type (
	Message      []Segment
	Segment      []Field
	Field        []FieldItem
	FieldItem    []Component
	Component    []Subcomponent
	Subcomponent string
)

func (m Message) Segments(name string) []Segment {
	var a []Segment

	for _, s := range m {
		if string(s[0][0][0][0]) == name {
			a = append(a, s)
		}
	}

	return a
}

func (m Message) Segment(name string, index int) Segment {
	i := 0
	for _, s := range m {
		if string(s[0][0][0][0]) == name {
			if i == index {
				return s
			}

			i++
		}
	}

	return nil
}

func (m Message) Query(s string) (res string, ok bool, err error) {
	q, err := ParseQuery(s)
	if err != nil {
		return "", false, stackerr.Wrap(err)
	}

	res, ok = q.Get(m)

	return res, ok, nil
}
