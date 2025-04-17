package doubles_test

import (
	"testing"

	"github.com/bgw6/doubles/doubles"
	"github.com/bgw6/doubles/service"
	"github.com/stretchr/testify/assert"
)

func TestSearchByNameFaked(t *testing.T) {
	// given
	myFakeSearchEngine := doubles.FakeSearchEngine{}
	engine := service.NewSearchEngine(&myFakeSearchEngine)

	testValues := map[string]string{"Nacho": "123456", "Nico": "234567"}

	for key := range testValues {
		engine.Engine.AddEntry(key, testValues[key])
	}

	resultNacho := engine.Engine.SearchByName("Nacho")
	resultNico := engine.Engine.SearchByName("Nico")
	resultPhone := engine.Engine.SearchByPhone("123456")

	assert.Equal(t, testValues["Nacho"], resultNacho)
	assert.Equal(t, testValues["Nico"], resultNico)
	assert.Equal(t, "Nacho", resultPhone)

}
