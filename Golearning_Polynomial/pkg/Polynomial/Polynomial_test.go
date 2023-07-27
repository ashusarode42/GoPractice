package package

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	t.Parallel()
	TestPolynomial := New([]int{3, 0, 1, 0})
	if len(TestPolynomial.poly_coefficients) != 0 {
		t.Logf("Polynomial is created")
	} else {
		t.Fatalf("Polynomial cannot be 0")
	}
}

func TestEval(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		input []int
		x     int
		want  int
	}{
		{
			name:  "Polynomail created",
			input: []int{2, 5, -7, 4},
			x:     2,
			want:  16,
		},
		{
			name:  "zeros are present ",
			input: []int{0, 4, 0, 5},
			x:     2,
			want:  48,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			polynomial := New(tc.input)
			expected := polynomial.Eval(tc.x)
			if expected == tc.want {
				return
			}
			t.Fatalf("Error Occured got %v want %v", expected, tc.want)
		})
	}
}
func TestOrder(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		input []int
		order int
	}{
		{
			name:  "Polynomail created",
			input: []int{2, 5, -7, 4},
			order: 2,
		},
		{
			name:  "Polynomail created",
			input: []int{0, 4, 0, 5},
			order: 0,
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			polynomial := New(tc.input)
			expected := polynomial.Order()
			if expected == tc.order {
				return
			}
			t.Fatalf("Error Occured  got %v want %v", expected, tc.order)
		})
	}
}
func TestMultiply(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name   string
		input1 []int
		input2 []int
		want   []int
		err    error
	}{
		{
			name:   "Multiply ",
			input1: []int{2, 5, -7, 4},
			input2: []int{0, 4, 0, 5},
			want:   []int{0, 8, 20, -18, 41, -35, 20},
		},
	}

	for _, testCase := range testCases {
		tc := testCase
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			polynomial_1 := New(tc.input1)
			polynomial_2 := New(tc.input2)
			want := New(tc.want)
			expected := polynomial_1.Multiply(polynomial_2)
			if reflect.DeepEqual(expected, want) {
				return
			} else {
				if tc.err != nil {
					return
				}
			}
			t.Fatalf("Error Occured got %v want %v", expected, tc.want)
		})
	}
}
func TestDivide(t *testing.T) {
	TestPolyQuotient, TestPolyRemainder, error := Polynomial.Divide(
		Polynomial{
			poly_coefficients: []int{4, 2}},
		Polynomial{
			poly_coefficients: []int{3, 7, 4, 2}})
	if len(TestPolyQuotient.poly_coefficients) != 0 || len(TestPolyRemainder.poly_coefficients) != 0 {
		t.Logf("Polynomial is divided")
	}
	if error != nil {
		t.Fatalf("Error in dividing Polynomial")
	}

}
