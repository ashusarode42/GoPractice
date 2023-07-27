package main

import (
	"pkg/Polynomial"
)


func main() {

	poly1 := New([]int{2, 5, -7, 4})
	poly2 := New([]int{0, 4, 0, 5})
	x := 2

	log.Println("First Polynomial poly1: ", poly1)
	log.Println("Second Polynomial poly2 : ", poly2)

	log.Println("Evaluation of poly1 : ", poly1.Eval(x))
	log.Println("Evaluation of poly2 : ", poly2.Eval(x))

	log.Println("Highest Order coefficient of poly1: ", poly1.Order())
	log.Println("Highest Order  coefficient of poly1: ", poly2.Order())

	result := poly1.Multiply(poly2)

	log.Println("Multiplied Polynomial is:", result)
	log.Println("(", poly1.Str(), ")*(", poly2.Str(), ") = ", result.Str())

	quo, p, err := poly1.Divide(poly2)
	if err == nil {
		log.Println(err)
	} else {
		log.Println("Quotient: ", quo.Str())
		log.Println("Remainder: ", p.Str())

	}
}
