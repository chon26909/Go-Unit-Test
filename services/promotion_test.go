package services_test

import (
	"go-test/repositories"
	"go-test/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	promoRepo := repositories.NewPromotionRepositoryMock()
	promoRepo.On("GetPromotion").Return(repositories.Promotion{
		Id:              1,
		PurchaseMin:     100,
		DiscountPercent: 20,
	}, nil)

	promoService := services.NewPromotionService(promoRepo)

	// Act
	discount, _ := promoService.CalculateDiscount(100)
	expected := 80

	// Assert
	assert.Equal(t, expected, discount)
}
