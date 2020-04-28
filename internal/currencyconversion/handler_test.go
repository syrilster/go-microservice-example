package currencyconversion

import (
	"context"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const fromCurrency string = "AUD"
const toCurrency string = "INR"
const quantity string = "12.55"

func TestCurrencyConversionHandler(t *testing.T) {
	var validPayload = fmt.Sprintf(`{
		"from": "%v",
		"to":"%v",
		"quantity":"%v",
	}`, fromCurrency, toCurrency, quantity)

	var inValidPayload = fmt.Sprintf(`{
		"from": "%v",
		"to":"%v",
	}`, fromCurrency, toCurrency)

	t.Run("valid request", func(t *testing.T) {
		req := request(t, strings.NewReader(validPayload), fromCurrency, toCurrency, quantity)
		rr := httptest.NewRecorder()
		mockHandler := new(MockExchangeRateFetcher)

		handler := http.HandlerFunc(handler(mockHandler))
		handler.ServeHTTP(rr, req)

		t.Run("success response", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rr.Code)
		})
	})

	t.Run("Invalid request", func(t *testing.T) {
		req := request(t, strings.NewReader(inValidPayload), fromCurrency, toCurrency, "")
		rr := httptest.NewRecorder()
		mockHandler := new(MockExchangeRateFetcher)

		handler := http.HandlerFunc(handler(mockHandler))
		handler.ServeHTTP(rr, req)

		t.Run("response is 400", func(t *testing.T) {
			assert.Equal(t, http.StatusBadRequest, rr.Code)
		})
	})

	t.Run("downstream error", func(t *testing.T) {
		req := request(t, strings.NewReader(inValidPayload), fromCurrency, toCurrency, "999")
		rr := httptest.NewRecorder()
		mockHandler := new(MockExchangeRateFetcher)

		handler := http.HandlerFunc(handler(mockHandler))
		handler.ServeHTTP(rr, req)

		t.Run("response is 500", func(t *testing.T) {
			assert.Equal(t, 500, rr.Code)
		})
	})
}

type MockExchangeRateFetcher struct {
	mock.Mock
}

func (m MockExchangeRateFetcher) FetchExchangeRate(ctx context.Context, req Request) (*Response, error) {
	quantity := fmt.Sprintf("%f", req.Quantity)
	if strings.EqualFold(quantity, "999.000000") {
		return &Response{}, errors.New("exchange rate fetch error")
	}
	return nil, nil
}

func request(t *testing.T, reader io.Reader, fromCurrency, toCurrency, quantity string) *http.Request {
	req, err := http.NewRequest("GET", "/currency-converter/from/"+fromCurrency+"/to/"+toCurrency+"/quantity/"+quantity, reader)

	//This is required even though the payload already has it!!
	vars := map[string]string{
		"from":     fromCurrency,
		"to":       toCurrency,
		"quantity": quantity,
	}
	req = mux.SetURLVars(req, vars)

	if err != nil {
		t.Fatalf("problem creating request: %+v", err)
	}
	return req
}
