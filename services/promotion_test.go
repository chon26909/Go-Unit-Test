package services_test

import (
	"errors"
	"go-test/repositories"
	"go-test/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {

	type testCase struct {
		name            string
		purchaseMin     int
		discountPercent int
		amount          int
		expected        int
	}

	cases := []testCase{
		{name: "applied 100", purchaseMin: 100, discountPercent: 20, amount: 100, expected: 80},
		{name: "applied 200", purchaseMin: 100, discountPercent: 20, amount: 200, expected: 160},
		{name: "applied 300", purchaseMin: 100, discountPercent: 20, amount: 300, expected: 240},
		{name: "not applied 50", purchaseMin: 100, discountPercent: 20, amount: 50, expected: 50},
	}

	for _, c := range cases {

		t.Run(c.name, func(t *testing.T) {
			promoRepo := repositories.NewPromotionRepositoryMock()
			promoRepo.On("GetPromotion").Return(repositories.Promotion{
				Id:              1,
				PurchaseMin:     c.purchaseMin,
				DiscountPercent: c.discountPercent,
			}, nil)

			promoService := services.NewPromotionService(promoRepo)

			discount, _ := promoService.CalculateDiscount(c.amount)

			assert.Equal(t, c.expected, discount)
		})

	}

	t.Run("purchase amount zero", func(t *testing.T) {
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{
			Id:              1,
			PurchaseMin:     100,
			DiscountPercent: 20,
		})

		promoService := services.NewPromotionService(promoRepo)

		// ถ้า discount เป็น 0
		_, err := promoService.CalculateDiscount(0)
		// จบการทำงาน แล้ว return error ออกมา
		assert.ErrorIs(t, err, services.ErrZeroAmount)
		// และต้องไม่เรียก GetPromotion
		promoRepo.AssertNotCalled(t, "GetPromotion")

	})

	t.Run("repository error", func(t *testing.T) {
		promoRepo := repositories.NewPromotionRepositoryMock()
		promoRepo.On("GetPromotion").Return(repositories.Promotion{}, errors.New(""))

		promoService := services.NewPromotionService(promoRepo)

		_, err := promoService.CalculateDiscount(100)

		assert.ErrorIs(t, err, services.ErrRepository)
	})

}
