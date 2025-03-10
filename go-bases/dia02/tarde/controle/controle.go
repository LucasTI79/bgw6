package controle

import "fmt"

var teste string

func ControleComIfElse() {
	fmt.Println("📌 Exemplo 1: Verificação de número")
	num := 0
	if num > 0 {
		fmt.Println(num, " é positivo")
	} else if num < 0 {
		fmt.Println(num, " é negativo")
	} else {
		fmt.Println("é Zero")
	}
}

func ControleComSwitch() {
	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Hoje é segunda feira !")
	case "Tuesday":
		fmt.Println("Hoje é Terça-feira!")
	case "Wednesday":
		fmt.Println("Hoje é Quarta-feira!")
	case "Thursday":
		fmt.Println("Hoje é Quinta-feira!")
	case "Friday":
		fmt.Println("Hoje é Sexta-feira!")
	case "Saturday", "Sunday":
		fmt.Println("É fim de semana!")
	default:
		fmt.Println("Dia inválido")
	}
}

func ControleComSwitchMultiplosCases() {
	day := "Tuesday"
	switch day {
	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("É um dia útil!")
	case "Saturday", "Sunday":
		fmt.Println("É fim de semana!")
	default:
		fmt.Println("Dia inválido!")
	}
}

func ControleComSwitchComFallthrough() {
	nota := 80

	switch {
	case nota >= 90:
		fmt.Println("Nota A")
		fallthrough
	case nota >= 80:
		fmt.Println("Nota B") // Executado devido ao fallthrough
		fallthrough
	case nota >= 70:
		fmt.Println("Nota C") // Executado devido ao fallthrough
	default:
		fmt.Println("Nota abaixo de C")
	}
}
