// Package assert contains helper functions used for testing.
package assert

import (
	"math"
	"testing"
)

func Equal(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("assert.Equal failed: Expected %v, got %v", a, b)
	}
}

func True(t *testing.T, v bool) {
	if v != true {
		t.Errorf("assert.True failed")
	}
}

func False(t *testing.T, v bool) {
	if v != false {
		t.Errorf("assert.False failed")
	}
}

func Nil(t *testing.T, v interface{}) {
	if v != nil {
		t.Errorf("assert.Nil failed")
	}
}

func NonNil(t *testing.T, v interface{}) {
	if v == nil {
		t.Errorf("assert.NonNil failed")
	}
}

func FloatsInDelta(t *testing.T, a, b, delta float64) {
	if math.Abs(a-b) >= delta {
		t.Errorf("assert.FloatsInDelta failed with |a - b| >= delta: a = %v, b = %v, delta = %v", a, b, delta)
	}
}
