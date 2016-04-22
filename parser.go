package hl7parser

import (
	"bytes"
	"errors"
)

var (
	// ErrTooShort is returned if a message isn't long enough to contain a valid
	// header
	ErrTooShort = errors.New("message must be at least eight bytes long")
	// ErrInvalidHeader is returned if a message doesn't start with "MSH"
	ErrInvalidHeader = errors.New("expected message to begin with MSH")
)

type (
	Message      []Segment
	Segment      []Field
	Field        []FieldItem
	FieldItem    []Component
	Component    []Subcomponent
	Subcomponent string
)

type Delimiters struct {
	Field, Component, Repeat, Escape, Subcomponent byte
}

func Parse(buf []byte) (Message, *Delimiters, error) {
	if len(buf) < 8 {
		return nil, nil, ErrTooShort
	}

	if !bytes.HasPrefix(buf, []byte("MSH")) {
		return nil, nil, ErrInvalidHeader
	}

	fs := buf[3]
	cs := buf[4]
	rs := buf[5]
	ec := buf[6]
	ss := buf[7]

	d := Delimiters{fs, cs, rs, ec, ss}

	var (
		message   Message
		segment   Segment
		field     Field
		fieldItem FieldItem
		component Component
		s         []byte
	)

	segment = Segment{
		Field{FieldItem{Component{Subcomponent("MSH")}}},
		Field{FieldItem{Component{Subcomponent(string(buf[3:8]))}}},
	}

	commitBuffer := func(force bool) {
		if s != nil || force {
			component = append(component, Subcomponent(unescape(s, &d)))
			s = nil
		}
	}

	commitComponent := func(force bool) {
		commitBuffer(false)

		if component != nil || force {
			fieldItem = append(fieldItem, component)
			component = nil
		}
	}

	commitFieldItem := func(force bool) {
		commitComponent(false)

		if fieldItem != nil || force {
			field = append(field, fieldItem)
			fieldItem = nil
		}
	}

	commitField := func(force bool) {
		commitFieldItem(false)

		if field != nil || force {
			segment = append(segment, field)
			field = nil
		}
	}

	commitSegment := func(force bool) {
		commitField(false)

		if segment != nil || force {
			message = append(message, segment)
			segment = nil
		}
	}

	for _, c := range buf[9:] {
		switch c {
		case '\r':
			commitSegment(true)
		case fs:
			commitField(true)
		case rs:
			commitFieldItem(true)
		case cs:
			commitComponent(true)
		case ss:
			commitBuffer(true)
		default:
			s = append(s, c)
		}
	}

	commitSegment(false)

	return message, &d, nil
}

func unescape(b []byte, d *Delimiters) []byte {
	r := make([]byte, len(b))

	j, e := 0, false
	for i := 0; i < len(b); i++ {
		c := b[i]

		switch e {
		case true:
			switch c {
			case 'F':
				r[j] = d.Field
				i++
			case 'S':
				r[j] = d.Component
				i++
			case 'T':
				r[j] = d.Subcomponent
				i++
			case 'R':
				r[j] = d.Repeat
				i++
			case 'E':
				r[j] = d.Escape
				i++
			default:
				r[j] = d.Escape
				j++
				r[j] = c
			}

			j++

			e = false
		case false:
			switch c {
			case d.Escape:
				e = true
			default:
				r[j] = c
				j++
			}
		}
	}

	return r[:j]
}
