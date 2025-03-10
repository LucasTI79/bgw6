package exemplo

import "fmt"

type Veiculo interface{}

type TrioEletrico struct {
	Nome string
}

type CarroAlegorico struct {
	Tema string
}

func Exemplo(v Veiculo) {
	if trio, ok := v.(TrioEletrico); ok {
		fmt.Println("Trio elétrico:", trio.Nome)
	} else {
		fmt.Println("tipo de veiculo desconhecido")
	}
}
