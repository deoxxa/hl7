package hl7

import (
	"bytes"

	"github.com/facebookgo/stackerr"
)

type (
	// ErrTooShort is returned if a message isn't long enough to contain a valid
	// header
	ErrTooShort error
	// ErrInvalidHeader is returned if a message doesn't start with "MSH"
	ErrInvalidHeader error
)

type Delimiters struct {
	Field, Component, Repeat, Escape, Subcomponent byte
}

// ParseMessage takes input as a `[]byte`, and returns the whole message, the
// control characters (as `*Delimiters`), and maybe an error.
func ParseMessage(buf []byte) (Message, *Delimiters, error) {
	// This is a sanity check, to make sure the message is long enough to
	// contain a valid header. If it's less than eight bytes long, it can't
	// possibly contain the required information.

	if len(buf) < 8 {
		return nil, nil, ErrTooShort(stackerr.Newf("message must be at least eight bytes long; instead was %d", len(buf)))
	}

	// Every valid HL7 message will begin with `MSH`. This isn't specifically
	// mandated in the specification, but by combining a few constraints, we can
	// safely come to this conclusion. This allows us to reject junk data pretty
	// quickly.

	if !bytes.HasPrefix(buf, []byte("MSH")) {
		return nil, nil, ErrInvalidHeader(stackerr.Newf("expected message to begin with MSH; instead found %q", buf[0:3]))
	}

	// These are the control characters. `fs` is the field separator, `cs` the
	// component separator, `rs` the field repeat separator, `ec` the escape
	// character, and `ss` the sub-component separator.

	fs := buf[3]
	cs := buf[4]
	rs := buf[5]
	ec := buf[6]
	ss := buf[7]

	d := Delimiters{fs, cs, rs, ec, ss}

	// These are the variables we'll be working with. We reuse these variables a
	// lot in the parsing loop below. A `FieldItem` is one instance of a field
	// value - the HL7 standard calls this a "repetition," but I found that
	// `FieldItem` was easier to think about.

	var (
		message   Message
		segment   Segment
		field     Field
		fieldItem FieldItem
		component Component
		s         []byte
	)

	// We manually construct the first few fields of the message, as we know
	// that it has to be structured this way. It's easier than having special
	// code to parse these weird fields out.

	segment = Segment{
		Field{FieldItem{Component{Subcomponent("MSH")}}},
		Field{FieldItem{Component{Subcomponent(buf[3])}}},
		Field{FieldItem{Component{Subcomponent(string(buf[4:8]))}}},
	}

	// These functions are used when we encounter control characters. When we
	// see a control character, it signals the end of a certain kind of element.
	// `|` means the end of a field, `~` a repetition, `^` a component, and `&`
	// a subcomponent. Another property of these separators is that each one not
	// only ends that element itself, but also any elements it contains. For
	// example, hitting `|` not only means that you've found the end of the
	// current field, but also the end of the current repetition, component, and
	// sub-component. This is expressed below as nested calls in the different
	// `commitX` functions.

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

	// This is the main parse loop. We go through the input byte-by-byte,
	// accumulating data until we hit any of the control characters. When we do,
	// we commit whatever we have "buffered" for that level. Carriage returns
	// and line breaks count as control characters, as they delimit segments
	// themselves.

	sawNewline := false
	for _, c := range buf[9:] {
		switch c {
		case '\r', '\n':
			if !sawNewline {
				commitSegment(true)
			}
			sawNewline = true
		case fs:
			sawNewline = false
			commitField(true)
		case rs:
			sawNewline = false
			commitFieldItem(true)
		case cs:
			sawNewline = false
			commitComponent(true)
		case ss:
			sawNewline = false
			commitBuffer(true)
		default:
			sawNewline = false
			s = append(s, c)
		}
	}

	// After we've gotten to the end of the input, we might still have some data
	// buffered up, so we make sure that gets committed.

	commitSegment(false)

	// That's it - we're done! Return the message, the `Delimiters` object, and
	// `nil` - signalling that there was no error.

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
