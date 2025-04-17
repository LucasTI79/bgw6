package doubles_test

import (
	"testing"

	"github.com/bgw6/doubles/doubles"
	"github.com/bgw6/doubles/service"
	"github.com/stretchr/testify/assert"
)

func TestSearchByNameMocked(t *testing.T) {
	myMockSearchEngine := doubles.MockSearchEngine{}
	engine := service.NewSearchEngine(&myMockSearchEngine)
	expectedPhone := "12345678"

	result := engine.Engine.SearchByName("Nacho")

	assert.Equal(t, expectedPhone, result)
	assert.True(t, myMockSearchEngine.SearchByNameWasCalled)
}

func TestSearchByNameMockedNotCalled(t *testing.T) {
	myMockSearchEngine := doubles.MockSearchEngine{}
	engine := service.NewSearchEngine(&myMockSearchEngine)
	expectedPhone := ""

	result := engine.Engine.SearchByName("Nac")

	assert.Equal(t, expectedPhone, result)
	assert.False(t, myMockSearchEngine.SearchByNameWasCalled)
}
