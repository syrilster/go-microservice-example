package currencyconversion

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/syrilster/go-microservice-example/internal/currencyexchange"
	"testing"
)

type MockCurrencyExchangeClient struct {
	mock.Mock
}

func (client MockCurrencyExchangeClient) GetExchangeRate(ctx context.Context, request currencyexchange.Request) (*currencyexchange.Response, error) {
	args := client.Called(ctx, request)
	return args.Get(0).(*currencyexchange.Response), args.Error(1)
}

func TestCurrencyConversion(t *testing.T) {
	var currencyExchangeRequest = currencyexchange.Request{
		FromCurrency: "AUD",
		ToCurrency:   "INR",
	}
	var mockResponse = &currencyexchange.Response{
		FromCurrency:       "AUD",
		ToCurrency:         "INR",
		ConversionMultiple: "49.5",
	}

	var currencyConversionReq = Request{
		FromCurrency: "AUD",
		ToCurrency:   "INR",
		Quantity:     12.5,
	}

	var expectedResponse = &Response{Amount: 618.75}

	t.Run("Success", func(t *testing.T) {
		mockClient := new(MockCurrencyExchangeClient)
		mockClient.On("GetExchangeRate", context.Background(), currencyExchangeRequest).Return(mockResponse, nil)
		service := NewService(mockClient)

		resp, err := service.FetchExchangeRate(context.Background(), currencyConversionReq)
		assert.Nil(t, err)
		assert.Equal(t, expectedResponse, resp)
		mockClient.AssertExpectations(t)
	})

	t.Run("error during currency conversion", func(t *testing.T) {
		mockClient := new(MockCurrencyExchangeClient)
		mockClient.On("GetExchangeRate", context.Background(), currencyExchangeRequest).Return(&currencyexchange.Response{}, errors.New("failed to get exchange rates"))
		service := NewService(mockClient)

		_, err := service.FetchExchangeRate(context.Background(), currencyConversionReq)
		assert.Error(t, err)
		assert.Equal(t, "failed to get exchange rates", err.Error())
		mockClient.AssertExpectations(t)
	})

	t.Run("when the service gets an invalid response from client", func(t *testing.T) {
		var invalidResponse = &currencyexchange.Response{
			FromCurrency:       "AUD",
			ToCurrency:         "INR",
			ConversionMultiple: "Test",
		}

		mockClient := new(MockCurrencyExchangeClient)
		mockClient.On("GetExchangeRate", context.Background(), currencyExchangeRequest).Return(invalidResponse, nil)
		service := NewService(mockClient)

		_, err := service.FetchExchangeRate(context.Background(), currencyConversionReq)
		assert.Error(t, err)
		assert.Equal(t, "strconv.ParseFloat: parsing \"Test\": invalid syntax", err.Error())
		mockClient.AssertExpectations(t)
	})
}
