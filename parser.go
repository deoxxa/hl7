package hl7parser

import (
	"bytes"
	"fmt"
	"strings"
)

type (
	Message      []Segment
	Segment      []Field
	Field        []FieldItem
	FieldItem    []Component
	Component    []Subcomponent
	Subcomponent string
)

func escaped(s string, d *Delimiters) string {
	return strings.NewReplacer(
		"\\F", d.Field,
		"\\S", d.Component,
		"\\T", d.Subcomponent,
		"\\R", d.Repeat,
		"\\E", d.Escape,
	).Replace(s)
}

type Delimiters struct {
	Field, Component, Repeat, Escape, Subcomponent string
}

func Parse(buf []byte) (Message, *Delimiters, error) {
	if len(buf) < 8 {
		return nil, nil, fmt.Errorf("message must be at least eight bytes long")
	}

	if !bytes.HasPrefix(buf, []byte("MSH")) {
		return nil, nil, fmt.Errorf("expected message to begin with MSH")
	}

	fs := buf[3]
	cs := buf[4]
	rs := buf[5]
	ec := buf[6]
	ss := buf[7]

	d := Delimiters{string(fs), string(cs), string(rs), string(ec), string(ss)}

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
			component = append(component, Subcomponent(s))
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

	escaping := false
	for _, c := range buf[9:] {
		switch escaping {
		case true:
			switch c {
			case fs, cs, rs, ec, ss:
				s = append(s, c)
			default:
				s = append(s, ec, c)
			}
			escaping = false
		case false:
			switch c {
			case ec:
				escaping = true
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
	}

	commitSegment(false)

	return message, &d, nil
}
