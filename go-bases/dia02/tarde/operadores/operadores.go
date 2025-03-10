package operadores

import "fmt"

func Operators() {
	fmt.Println("Operadores Aritméticos")
	a, b := 15, 4
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)

	fmt.Println("\n🎯 Operadores de Atribuição")
	x := 10
	fmt.Println("x é igual a: ", x)

	x += 5 // x + 5
	fmt.Println("X é igual a: ", x)

	x -= 5 // x - 5
	fmt.Println("X é igual a: ", x)

	x *= 5
	fmt.Println("X é igual a:", x)

	x /= 5
	fmt.Println("X é igual a:", x)

	// fmt.Println("\n⚖️ Operadores de Comparação")
	p, q := 10, 20
	fmt.Println("p == q", p == q)
	fmt.Println("p != q", p != q)
	fmt.Println("p > q", p > q)
	fmt.Println("p < q", p < q)
	fmt.Println("p <= q", p <= q)
	fmt.Println("p >= q", p >= q)

	fmt.Println("\n🤔 Operadores Lógicos")
	// E OU =  AND OR
	cond1, cond2 := true, false
	fmt.Print("Cond1 && cond2: ", cond1 && cond2)
	fmt.Print("Cond1 || cond2: ", cond1 || cond2)
	fmt.Println("!Cond1: ", !cond1)

	fmt.Println("\n📌 Operadores de Endereço")
	var prt *int
	prt = &p
	fmt.Println("Endereço de p →", prt)        // Mostra o endereço de memória de p
	fmt.Println("Valor da variavel p →", *prt) // Mostra Valor da variavel p

}
