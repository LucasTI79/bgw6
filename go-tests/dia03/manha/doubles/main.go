package main

import (
	"fmt"

	"github.com/bgw6/doubles/service"
)

func main() {
	engine := service.NewSearchEngine(nil)

	fmt.Printf("Search Engine Version: %s\n", engine.GetVersion())
}
