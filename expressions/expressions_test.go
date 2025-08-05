package expressions

import (
	"fmt"
	"testing"
)

func TestExpressionsBasic(t *testing.T) {
	tests := []struct {
		expr     string
		env      Env
		expected string
	}{
		{"add(3, 5)", Env{}, "8"},
		{"subtract(9, 6)", Env{}, "3"},
		{"multiply(9, 8)", Env{}, "72"},
	}

	for _, test := range tests {
		expr, _ := Parse(test.expr)
		fmt.Println(expr.Evaluate())
	}
}
