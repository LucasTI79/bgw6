package main

import (
	"fmt"

	"github.com/bgw6/functions/calc"
)

func main() {
	// fmt.Println(calc.Calculate(calc.Sum, 2, 3, 4))
	// sum, sub, mult, div := calc.Operations(6, 2)

	// fmt.Println("Summatory:\t\t", sum)
	// fmt.Println("Subtraction:\t\t", sub)
	// fmt.Println("Multiplication:\t", mult)
	// fmt.Println("Division:\t", div)

	// result, err := calc.Division(6, 0)

	// if err != nil {
	// 	switch {
	// 	case errors.Is(err, calc.ErrCannotDivideByZero):
	// 		fmt.Println("o zé ruela, não pode dividir por 0")
	// 	default:
	// 		{
	// 			fmt.Printf("ops..., um erro inesperado aconteceu: %v\n", err)
	// 		}
	// 	}

	// 	return
	// }

	// fmt.Println(result)

	sumFn := calc.CalculateCurrying(calc.Sum)
	fmt.Println("valor da soma é: ", sumFn(2, 3))
}
