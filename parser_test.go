package hl7parser

import (
	"io/ioutil"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func mustReadFile(f string) []byte {
	d, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return d
}

var (
	allElementsContent = mustReadFile("testdata/all_elements.hl7")
	pdfGeneticsContent = mustReadFile("testdata/pdf_genetics.hl7")
	sampleContent      = mustReadFile("testdata/sample.hl7")
	simpleContent      = mustReadFile("testdata/simple.hl7")
	simpleNohexContent = mustReadFile("testdata/simple_nohex.hl7")
	vaersLongContent   = mustReadFile("testdata/vaers_long.hl7")
)

func TestParseOneSegment(t *testing.T) {
	a := assert.New(t)

	m, d, err := Parse([]byte(`MSH|^~\&|IPM|1919|SUPERHOSPITAL|1919|20160101000000||ADT^A08|555544444|D|2.4|||AL|NE`))
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{FieldItem{Component{Subcomponent("MSH")}}},
			Field{FieldItem{Component{Subcomponent("|^~\\&")}}},
			Field{FieldItem{Component{Subcomponent("IPM")}}},
			Field{FieldItem{Component{Subcomponent("1919")}}},
			Field{FieldItem{Component{Subcomponent("SUPERHOSPITAL")}}},
			Field{FieldItem{Component{Subcomponent("1919")}}},
			Field{FieldItem{Component{Subcomponent("20160101000000")}}},
			nil,
			Field{FieldItem{
				Component{Subcomponent("ADT")},
				Component{Subcomponent("A08")},
			}},
			Field{FieldItem{Component{Subcomponent("555544444")}}},
			Field{FieldItem{Component{Subcomponent("D")}}},
			Field{FieldItem{Component{Subcomponent("2.4")}}},
			nil,
			nil,
			Field{FieldItem{Component{Subcomponent("AL")}}},
			Field{FieldItem{Component{Subcomponent("NE")}}},
		},
	}, m)
}

func TestParseTwoSegments(t *testing.T) {
	a := assert.New(t)

	m, d, err := Parse([]byte(strings.Join([]string{
		`MSH|^~\&|IPM|1919|SUPERHOSPITAL|1919|20160101000000||ADT^A08|555544444|D|2.4|||AL|NE`,
		`EVN|A08|20160101000001||BATMAN_U|SHBOLTONM^Bolton, Michael^^^^^^USERS`,
	}, "\r")))
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{FieldItem{Component{Subcomponent("MSH")}}},
			Field{FieldItem{Component{Subcomponent("|^~\\&")}}},
			Field{FieldItem{Component{Subcomponent("IPM")}}},
			Field{FieldItem{Component{Subcomponent("1919")}}},
			Field{FieldItem{Component{Subcomponent("SUPERHOSPITAL")}}},
			Field{FieldItem{Component{Subcomponent("1919")}}},
			Field{FieldItem{Component{Subcomponent("20160101000000")}}},
			nil,
			Field{FieldItem{
				Component{Subcomponent("ADT")},
				Component{Subcomponent("A08")},
			}},
			Field{FieldItem{Component{Subcomponent("555544444")}}},
			Field{FieldItem{Component{Subcomponent("D")}}},
			Field{FieldItem{Component{Subcomponent("2.4")}}},
			nil,
			nil,
			Field{FieldItem{Component{Subcomponent("AL")}}},
			Field{FieldItem{Component{Subcomponent("NE")}}},
		},
		Segment{
			Field{FieldItem{Component{Subcomponent("EVN")}}},
			Field{FieldItem{Component{Subcomponent("A08")}}},
			Field{FieldItem{Component{Subcomponent("20160101000001")}}},
			nil,
			Field{FieldItem{Component{Subcomponent("BATMAN_U")}}},
			Field{FieldItem{
				Component{Subcomponent("SHBOLTONM")},
				Component{Subcomponent("Bolton, Michael")},
				nil,
				nil,
				nil,
				nil,
				nil,
				Component{Subcomponent("USERS")},
			}},
		},
	}, m)
}

// func TestParseSimpleContent(t *testing.T) {
// 	a := assert.New(t)

// 	m, d, err := Parse(simpleContent)
// 	a.NoError(err)
// 	a.Equal(&Delimiters{"|", "^", "~", "\\", "&"}, d)
// 	a.Equal(Message{}, m)
// }

func BenchmarkAllElementsContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(allElementsContent)
	}
}

func BenchmarkSampleContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(sampleContent)
	}
}

func BenchmarkSimpleContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(simpleContent)
	}
}

func BenchmarkSimpleNohexContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(simpleNohexContent)
	}
}

func BenchmarkVaersLongContent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(vaersLongContent)
	}
}
