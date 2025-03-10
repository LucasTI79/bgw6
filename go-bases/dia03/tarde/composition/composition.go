package composition

import "fmt"

type Autor struct {
	Nome      string
	Sobrenome string
	Endereco  Address
}

type Address struct {
	Rua    string
	Bairro string
	Numero int
}

func (ad *Address) ExibeBairro() {
	fmt.Println(ad.Rua)
}

func (a *Autor) Detail() {
	fmt.Println(a.Nome, a.Sobrenome, a.Endereco.Rua)
	a.Endereco.ExibeBairro()
}
