package assert

import "testing"

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
