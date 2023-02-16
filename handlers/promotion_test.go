package handlers_test

import (
	"fmt"
	"go-test/handlers"
	"go-test/services"
	"io"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestPromotionCalculateDiscount(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		amount := 100
		expected := 80

		promoService := services.NewPromotionServiceMock()
		promoService.On("CalculateDiscount", amount).Return(expected, nil)

		promoHandler := handlers.NewPromotionHandler(promoService)

		app := fiber.New()
		app.Get("/calculate", promoHandler.CalculateDiscount)

		req := httptest.NewRequest("GET", fmt.Sprintf("/calculate?amount=%v", amount), nil)
		res, _ := app.Test(req)

		if assert.Equal(t, fiber.StatusOK, res.StatusCode) {

			body, _ := io.ReadAll(res.Body)
			assert.Equal(t, strconv.Itoa(expected), string(body))
		}
	})
}
