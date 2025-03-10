package interfaces

import "fmt"

type Veiculo interface {
	AbriPorta()
	Soma()
}

type Carro struct{}

func (c Carro) AbriPorta() {
	fmt.Println("ABRIU A PORTA UAI")
}
func (c Carro) Soma() {
	fmt.Println("SOMEU OS VALORES UÉ")
}

type Mosquito struct{}

func (m Mosquito) AbriPorta() {
	fmt.Println("ABRIU A PORTA UAI")
}
func (m Mosquito) Soma() {
	fmt.Println("SOMEU OS VALORES UÉ")
}
