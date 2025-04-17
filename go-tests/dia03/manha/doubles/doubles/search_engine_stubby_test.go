package doubles_test

import (
	"testing"

	"github.com/bgw6/doubles/doubles"
	"github.com/bgw6/doubles/service"
	"github.com/stretchr/testify/assert"
)

func TestFindByName(t *testing.T) {
	// given
	engine := &doubles.SearchEngineStub{}
	searchEngine := service.NewSearchEngine(engine)
	expectedPhone := "123456"

	// when

	result := searchEngine.Engine.SearchByName("John Doe")

	// then

	assert.Equal(t, expectedPhone, result)
}
