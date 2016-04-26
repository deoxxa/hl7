package hl7

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	a := assert.New(t)

	q := New("MSH", 0, 0, 0, 0, 0)
	a.Equal(q, Query{Segment: "MSH"})
}
