package products_test

import (
	"testing"

	"github.com/bgw6/pod/internal/domain"
	"github.com/bgw6/pod/internal/products"
	"github.com/bgw6/pod/tests/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetById(t *testing.T) {
	t.Run("should be return a product", func(t *testing.T) {
		// given
		repository := new(mocks.ProductRepositoryMock)
		product := domain.Product{
			Name: "Batata",
		}
		service := products.NewService(repository)
		expectedResult := &product
		productId := 1

		// when

		repository.Mock.On(
			"GetByID",
			mock.AnythingOfType("int"),
		).Return(
			&product,
			nil,
		)
		result, err := service.GetByID(productId)

		// then

		assert.NoError(t, err)
		assert.Equal(t, expectedResult, result)
		repository.AssertCalled(t, "GetByID")
		repository.AssertExpectations(t)
	})
}
