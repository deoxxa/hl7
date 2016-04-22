package hl7parser

import (
	"bytes"
	"fmt"
	"strings"

	"fknsrs.biz/p/supersplit"
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

	fs := string(buf[3])
	cs := string(buf[4])
	rs := string(buf[5])
	ec := string(buf[6])
	ss := string(buf[7])

	d := Delimiters{fs, cs, rs, ec, ss}

	lines := bytes.Split(buf, []byte("\r"))

	m := make(Message, len(lines))

	for lIndex, l := range lines {
		// cut off the header (two fields) for the first segment
		if lIndex == 0 {
			l = l[9:]
		}

		lBits := supersplit.Escaped(string(l), fs, ec)

		// put two empty fields at the start for the first segment
		if lIndex == 0 {
			lBits = append([]string{"", ""}, lBits...)
		}

		segment := make(Segment, len(lBits))

		for findex, fString := range lBits {
			fBits := supersplit.Escaped(fString, rs, ec)

			field := make(Field, len(fBits))

			for rIndex, rString := range fBits {
				rBits := supersplit.Escaped(rString, cs, ec)

				fieldItem := make(FieldItem, len(rBits))

				for cIndex, cString := range rBits {
					cBits := supersplit.Escaped(cString, ss, ec)

					component := make(Component, len(cBits))

					for sIndex, sString := range cBits {
						component[sIndex] = Subcomponent(escaped(sString, &d))
					}

					fieldItem[cIndex] = component
				}

				field[rIndex] = fieldItem
			}

			segment[findex] = field
		}

		// manually fill in the header for the first segment
		if lIndex == 0 {
			segment[0] = Field{FieldItem{Component{Subcomponent("MSH")}}}
			segment[1] = Field{FieldItem{Component{Subcomponent(string(buf[3:8]))}}}
		}

		m[lIndex] = segment
	}

	return m, &d, nil
}
