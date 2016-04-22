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
			Field{FieldItem{Component{"MSH"}}},
			Field{FieldItem{Component{"|^~\\&"}}},
			Field{FieldItem{Component{"IPM"}}},
			Field{FieldItem{Component{"1919"}}},
			Field{FieldItem{Component{"SUPERHOSPITAL"}}},
			Field{FieldItem{Component{"1919"}}},
			Field{FieldItem{Component{"20160101000000"}}},
			nil,
			Field{FieldItem{
				Component{"ADT"},
				Component{"A08"},
			}},
			Field{FieldItem{Component{"555544444"}}},
			Field{FieldItem{Component{"D"}}},
			Field{FieldItem{Component{"2.4"}}},
			nil,
			nil,
			Field{FieldItem{Component{"AL"}}},
			Field{FieldItem{Component{"NE"}}},
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
			Field{FieldItem{Component{"MSH"}}},
			Field{FieldItem{Component{"|^~\\&"}}},
			Field{FieldItem{Component{"IPM"}}},
			Field{FieldItem{Component{"1919"}}},
			Field{FieldItem{Component{"SUPERHOSPITAL"}}},
			Field{FieldItem{Component{"1919"}}},
			Field{FieldItem{Component{"20160101000000"}}},
			nil,
			Field{FieldItem{
				Component{"ADT"},
				Component{"A08"},
			}},
			Field{FieldItem{Component{"555544444"}}},
			Field{FieldItem{Component{"D"}}},
			Field{FieldItem{Component{"2.4"}}},
			nil,
			nil,
			Field{FieldItem{Component{"AL"}}},
			Field{FieldItem{Component{"NE"}}},
		},
		Segment{
			Field{FieldItem{Component{"EVN"}}},
			Field{FieldItem{Component{"A08"}}},
			Field{FieldItem{Component{"20160101000001"}}},
			nil,
			Field{FieldItem{Component{"BATMAN_U"}}},
			Field{FieldItem{
				Component{"SHBOLTONM"},
				Component{"Bolton, Michael"},
				nil,
				nil,
				nil,
				nil,
				nil,
				Component{"USERS"},
			}},
		},
	}, m)
}

func TestParseSampleContent(t *testing.T) {
	a := assert.New(t)

	m, d, err := Parse(sampleContent)
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{FieldItem{Component{"MSH"}}},
			Field{FieldItem{Component{"|^~\\&"}}},
			Field{FieldItem{Component{"EP^IC"}}},
			Field{FieldItem{Component{"EPICADT"}}},
			Field{FieldItem{Component{"SMS"}}},
			Field{FieldItem{Component{"SMSADT"}}},
			Field{FieldItem{Component{"199912271408"}}},
			Field{FieldItem{Component{"CHARRIS"}}},
			Field{FieldItem{Component{"ADT"}, Component{"A04"}}},
			Field{FieldItem{Component{"1817457"}}},
			Field{FieldItem{Component{"D"}}},
			Field{FieldItem{Component{"2.5"}}},
		},
		Segment{
			Field{FieldItem{Component{"PID"}}},
			nil,
			Field{FieldItem{Component{"0493575"}, nil, nil, Component{"2"}, Component{"ID 1"}}},
			Field{FieldItem{Component{"454721"}}},
			nil,
			Field{FieldItem{Component{"DOE"}, Component{"JOHN"}, nil, nil, nil}},
			Field{FieldItem{Component{"DOE"}, Component{"JOHN"}, nil, nil, nil}},
			Field{FieldItem{Component{"19480203"}}},
			Field{FieldItem{Component{"M"}}},
			nil,
			Field{FieldItem{Component{"B"}}},
			Field{FieldItem{Component{"254 MYSTREET AVE"}, nil, Component{"MYTOWN"}, Component{"OH"}, Component{"44123"}, Component{"USA"}}},
			nil,
			Field{FieldItem{Component{"(216)123-4567"}}},
			nil,
			nil,
			Field{FieldItem{Component{"M"}}},
			Field{FieldItem{Component{"NON"}}},
			Field{FieldItem{Component{"400003403"}}, FieldItem{Component{"1129086"}}}},
		Segment{
			Field{FieldItem{Component{"NK1"}}},
			nil,
			Field{FieldItem{Component{"ROE"}, Component{"MARIE"}, nil, nil, nil}},
			Field{FieldItem{Component{"SPO"}}},
			nil,
			Field{FieldItem{Component{"(216)123-4567"}}},
			nil,
			Field{FieldItem{Component{"EC"}}},
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
		},
		Segment{
			Field{FieldItem{Component{"PV1"}}},
			nil,
			Field{FieldItem{Component{"O"}}},
			Field{FieldItem{Component{"168 "}}, FieldItem{Component{"219"}}, FieldItem{Component{"C"}}, FieldItem{Component{"PMA"}, nil, nil, nil, nil, nil, nil, nil, nil}},
			nil,
			nil,
			nil,
			Field{FieldItem{Component{"277"}, Component{"ALLEN MYLASTNAME"}, Component{"BONNIE"}, nil, nil, nil}},
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			Field{FieldItem{Component{" "}}},
			nil,
			Field{FieldItem{Component{"2688684"}}},
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			nil,
			Field{FieldItem{Component{"199912271408"}}},
			nil,
			nil,
			nil,
			nil,
			nil,
			Field{FieldItem{Component{"002376853\""}}},
		},
	}, m)

	t.Logf("%#v\n", m)
}

func TestParseSimpleNohexContent(t *testing.T) {
	a := assert.New(t)

	m, d, err := Parse(simpleNohexContent)
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{FieldItem{Component{"MSH"}}},
			Field{FieldItem{Component{"|^~\\&"}}},
			Field{FieldItem{Component{"field"}}},
			Field{FieldItem{
				Component{"\\|~^&HEY"},
			}},
			Field{FieldItem{
				Component{"component1"},
				Component{"component2"},
			}},
			Field{FieldItem{
				Component{"subcomponent1a", "subcomponent2a"},
				Component{"subcomponent1b", "subcomponent2b"},
			}},
			Field{
				FieldItem{Component{"component1a"}, Component{"component2a"}},
				FieldItem{Component{"component1b"}, Component{"component2b"}},
			},
		},
	}, m)
}

func TestParseSimpleContent(t *testing.T) {
	a := assert.New(t)

	m, d, err := Parse(simpleContent)
	a.NoError(err)
	a.Equal(&Delimiters{'|', '^', '~', '\\', '&'}, d)
	a.Equal(Message{
		Segment{
			Field{FieldItem{Component{"MSH"}}},
			Field{FieldItem{Component{"|^~\\&"}}},
			Field{FieldItem{Component{"field"}}},
			Field{FieldItem{
				Component{"\\|~^&\\X484559"},
			}},
			Field{FieldItem{
				Component{"component1"},
				Component{"component2"},
			}},
			Field{FieldItem{
				Component{"subcomponent1a", "subcomponent2a"},
				Component{"subcomponent1b", "subcomponent2b"},
			}},
			Field{
				FieldItem{Component{"component1a"}, Component{"component2a"}},
				FieldItem{Component{"component1b"}, Component{"component2b"}},
			},
		},
	}, m)
}

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
