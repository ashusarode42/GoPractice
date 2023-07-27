package Polynomail

import (
	"errors"
	"fmt"
	"strconv"
)

// Polynomial represents polynomial with integer coefficients.
type Polynomial struct {
	poly_coefficients []int
}

// New initializes a new polynomial expressiom.
func New(coefficients []int) Polynomial {
	return Polynomial{poly_coefficients: coefficients}
}

// Eval will subsitiued the interger value and evaluate the polynomial expression.
func (p Polynomial) Eval(x int) int {
	var counter int
	result := p.poly_coefficients[0]
	counter = len(p.poly_coefficients) - 1
	power := 1
	// Using Honers Mehtod to evaluate value of ploynomial.
	for i := 1; i <= counter; i++ {
		power = x * power
		result = p.poly_coefficients[i]*power + result
	}
	return result
}

//Order func will calculate the highest order of the polynomial expression.
func (p Polynomial) Order() int {

	var result int
	counter := len(p.poly_coefficients) - 1
	for i := 0; i <= counter; i++ {
		if p.poly_coefficients[i] != 0 {
			result = p.poly_coefficients[i]
			break
		} else {
			return p.poly_coefficients[i]
		}

	}
	return result
}

// Multiply func perform multiplication of two Polynomials expression.
func (p Polynomial) Multiply(q Polynomial) Polynomial {

	Polynomial_1 := len(p.poly_coefficients) - 1
	Polynomial_2 := len(q.poly_coefficients) - 1
	result := Polynomial{make([]int, Polynomial_1+Polynomial_2+1)}

	for i := 0; i <= Polynomial_1; i++ {
		for j := 0; j <= Polynomial_2; j++ {
			// Multiplying every term of second polynomial with first polynomail term.
			result.poly_coefficients[i+j] = p.poly_coefficients[i]*q.poly_coefficients[j] + result.poly_coefficients[i+j]
		}
	}
	return result
}

// Str function return the string form of polynomial.
func (p Polynomial) Str() string {
	d := len(p.poly_coefficients) - 1

	if d == -1 {
		return "0"
	}

	s := ""
	//getting coefficent with highest order .
	if p.poly_coefficients[0] == -1 {
		if d == 0 {
			s += "-1"
		} else {
			s += "-"
		}
	} else if p.poly_coefficients[0] != 1 || d == 0 {
		s += strconv.Itoa(p.poly_coefficients[0])
	}
	if d == 1 {
		s += "x"
	} else if d > 1 {
		s += "x^" + strconv.Itoa(d)
	}
	for i := d - 1; i >= 0; i-- {
		// Operators and signs representation.
		if p.poly_coefficients[d-i] > 0 {
			s += " + "
			if p.poly_coefficients[d-i] != 1 || i == 0 {
				s += strconv.Itoa(p.poly_coefficients[d-i])
			}
		} else if p.poly_coefficients[d-i] < 0 {
			s += " - "
			if p.poly_coefficients[d-i] != -1 || i == 0 {
				s += strconv.Itoa(p.poly_coefficients[d-i] * -1)
			}
		}
		// variable and power representation.
		if p.poly_coefficients[d-i] != 0 {
			if i == 1 {
				s += "x"
			} else if i > 1 {
				s += "x^" + strconv.Itoa(i)
			}
		}
	}

	return s
}

//Deg return the degree of polynomial.
func (p Polynomial) Deg() int {
	i := len(p.poly_coefficients) - 1
	return i
}

// Add will add the 2 polynomial.
func Add(p, q Polynomial) Polynomial {
	var m, n, result Polynomial
	// Check which polynomial is bigger.
	if p.Deg() > q.Deg() {
		m = p
		n = q
	} else {
		m = q
		n = p
	}
	//Assign the result to polynomial
	result = m
	for i := 0; i <= n.Deg(); i += 1 {
		result.poly_coefficients[i] += n.poly_coefficients[i]
	}
	return result
}

// Substract will find the difference between 2 polynomial.
func Substract(p, q Polynomial) Polynomial {
	degp := p.Deg()
	degq := q.Deg()
	big := degp
	if degp < degq {
		big = degq
	}
	r := Polynomial{make([]int, big+1)}
	for i := 0; i <= big; i++ {
		if i <= degp {
			r.poly_coefficients[i] = p.poly_coefficients[i]
		}
		if i <= degq {
			r.poly_coefficients[i] -= q.poly_coefficients[i]
		}
	}
	return r
}

//Divide will divide the polynomial and return the quotient,remainder and error.
func (p Polynomial) Divide(q Polynomial) (Polynomial, Polynomial, error) {
	var quo Polynomial

	for q.Deg() <= p.Deg() {
		dega := p.Deg()
		degb := q.Deg()
		if p.poly_coefficients[dega] < q.poly_coefficients[degb] {
			return p, q, errors.New("Invalid inputs, a should be greater than b")
		}
		qc := p.poly_coefficients[dega] / q.poly_coefficients[degb]
		qp := dega - degb
		qi := Polynomial{make([]int, qp+1)}
		qi.poly_coefficients[qp] = qc
		p = Substract(p, qi.Multiply(q))
		quo = Add(quo, qi)
	}
	return quo, p, nil
}

func main() {
	// Creating New polynomials for arithmatic operation.
	poly1 := New([]int{2, 5, -7, 4})
	poly2 := New([]int{0, 4, 0, 5})
	x := 2

	fmt.Println("First Polynomial poly1: ", poly1)
	fmt.Println("Second Polynomial poly2 : ", poly2)

	fmt.Println("Evaluation of poly1 : ", poly1.Eval(x))
	fmt.Println("Evaluation of poly2 : ", poly2.Eval(x))

	fmt.Println("Highest Order coefficient of poly1: ", poly1.Order())
	fmt.Println("Highest Order  coefficient of poly1: ", poly2.Order())

	result := poly1.Multiply(poly2)

	fmt.Println("Multiplied Polynomial is:", result)
	fmt.Println("(", poly1.Str(), ")*(", poly2.Str(), ") = ", result.Str())

	quo, p, err := poly1.Divide(poly2)
	if err == nil {
		fmt.Println(err)
	} else {
		fmt.Println("Quotient: ", quo.Str())
		fmt.Println("Remainder: ", p.Str())

	}
}
