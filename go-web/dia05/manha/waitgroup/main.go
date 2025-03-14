package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup // Criando um WaitGroup

func tarefa(id int) {
	defer wg.Done() // Indica que a goroutine terminou
	time.Sleep(time.Second * 2)
	fmt.Println("Tarefa", id, "concluída")
}

func main() {
	for i := 1; i <= 5; i++ {
		wg.Add(1) // Adiciona 1 ao contador do WaitGroup
		go tarefa(i)
	}

	wg.Wait() // Aguarda todas as goroutines finalizarem

	fmt.Println("Fim do programa")
}
