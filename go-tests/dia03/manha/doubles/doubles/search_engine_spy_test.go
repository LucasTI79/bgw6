package doubles_test

import (
	"testing"

	"github.com/bgw6/doubles/doubles"
	"github.com/bgw6/doubles/service"
	"github.com/stretchr/testify/assert"
)

func TestSearchByPhone(t *testing.T) {
	mySpySearchEngine := doubles.SpySearchEngine{SearchByPhoneWasCalled: false}
	engine := service.NewSearchEngine(&mySpySearchEngine)
	phone := "12345678" // em nosso mecanismo, não realizaremos pesquisas se o telefone
	// tem um comprimento inferior a 5

	engine.Engine.SearchByPhone(phone)

	assert.True(t, mySpySearchEngine.SearchByPhoneWasCalled)
}

func TestSearchByPhoneNotCalled(t *testing.T) {
	mySpySearchEngine := doubles.SpySearchEngine{}
	engine := service.NewSearchEngine(&mySpySearchEngine)

	engine.Engine.SearchByPhone("1234") // o telefone tem menos de 5 caracteres

	assert.False(t, mySpySearchEngine.SearchByPhoneWasCalled)
}
