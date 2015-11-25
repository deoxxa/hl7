package hl7terser

import (
	"testing"
)

func TestNew(t *testing.T) {
	q1 := New("MSH", 0, 0, 0, 0, 0, 0, 0)
	if q1.Segment != "MSH" {
		t.Errorf("segment should be MSH")
	}
	if q1.Component != 1 {
		t.Errorf("component should be 1")
	}
	if q1.SubComponent != 1 {
		t.Errorf("subcomponent should be 1")
	}
}
