package service_test

import (
	"testing"

	"github.com/bgw6/doubles/doubles"
	"github.com/bgw6/doubles/service"
	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	// given
	dummyEngine := &doubles.SearchEngineDummy{}
	searchEngine := service.NewSearchEngine(dummyEngine)
	expectedVersion := "1.0.0"

	// when
	result := searchEngine.GetVersion()

	//then
	assert.Equal(t, expectedVersion, result)
}
