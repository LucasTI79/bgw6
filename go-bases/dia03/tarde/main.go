package main

import "interfaces/exemplo"

func main() {
	// autor := composition.Autor{
	// 	Nome:      "Natan",
	// 	Sobrenome: "Menezes",
	// 	Endereco: composition.Address{
	// 		Rua:    "R. Guajajaras",
	// 		Bairro: "Centro",
	// 		Numero: 1522,
	// 	},
	// }

	// autor.Detail()

	// abada := 42
	// // fmt.Println(abada)

	// ponteiro.TrocaValorDoAbada(&abada, 45)

	// fmt.Println("Deveria ser 99, mas o preço é ", abada)

	// var i interface{} = "3"

	// s := i.(string)
	// fmt.Println(s)

	// s, ok := i.(string)
	// fmt.Println("Valor :", s, "Sucsso:", ok)

	// f, ok := i.(float64)
	// fmt.Println("Valor :", f, "Sucsso:", ok)

	t := exemplo.TrioEletrico{Nome: "Axé Bahia"}
	x := "Fantasia dos Mágicos"
	exemplo.Exemplo(t)
	exemplo.Exemplo(x)
}
