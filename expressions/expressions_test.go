package expressions

import (
	"fmt"
	"go-book/assert"
	"testing"
)

const (
	delta = 0.01
)

func TestExpressionsBasic(t *testing.T) {
	tests := []struct {
		expr     string
		env      Env
		expected float64
	}{
		{"add(3, 5)", Env{}, 8},
		{"subtract(100, 36)", Env{}, 64},
		{"multiply(13, 15)", Env{}, 195},
		{"divide(144, 12)", Env{}, 12},
		{"sqrt(81)", Env{}, 9},
		{"sin(1)", Env{}, 0.84147},
	}

	for _, test := range tests {
		expr, err := Parse(test.expr)
		assert.Nil(t, err)
		result, err := expr.Evaluate()
		assert.Nil(t, err)
		assert.FloatsInDelta(t, result, test.expected, delta)
	}
}

func TestExpressionsNested(t *testing.T) {
	tests := []struct {
		expr     string
		env      Env
		expected float64
	}{
		{"add(multiply(3, 5), 27)", Env{}, 42},
		{"add(multiply(3, 5), subtract(123, 85))", Env{}, 53},
		//{"subtract(1000, add(multiply(3, 5), 27))", Env{}, 958},
	}

	for _, test := range tests {
		expr, err := Parse(test.expr)
		assert.Nil(t, err)
		result, err := expr.Evaluate()
		fmt.Println(result)
		assert.Nil(t, err)
		assert.FloatsInDelta(t, result, test.expected, delta)
	}
}
