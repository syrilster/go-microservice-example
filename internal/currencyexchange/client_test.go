package currencyexchange

import (
	"bytes"
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

// MockHTTPClient is to mock the external calls
type MockHTTPClient struct {
	mock.Mock
}

// Do is to mock the external calls
func (mockClient *MockHTTPClient) Do(r *http.Request) (*http.Response, error) {
	args := mockClient.Called(r)
	return args.Get(0).(*http.Response), args.Error(1)
}

func TestCurrencyExchangeClient(t *testing.T) {
	var testURL = "http://test-url.com/"
	var invalidJSONResponse = `
		{
			"results": [
				{
					"document": {
						"name": "test_name"
					},
				}
			],
		}`
	var request = Request{
		FromCurrency: "AED",
		ToCurrency:   "INR",
	}
	var expected = &Response{
		FromCurrency:       "AED",
		ToCurrency:         "INR",
		ConversionMultiple: "20.72",
	}

	t.Run("Fetch Exchange Rate", func(t *testing.T) {
		var validJSON = `{
						  "from": "AED",
						  "to": "INR",
						  "conversion_multiple" : "20.72"
						}`
		t.Run("Success", func(t *testing.T) {
			mockHTTPClient := new(MockHTTPClient)
			mockResponse := &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(validJSON))),
			}

			mockHTTPClient.On("Do", mock.Anything).Return(mockResponse, nil)
			client := client{
				URL:         testURL,
				HTTPCommand: mockHTTPClient,
			}

			actual, err := client.GetExchangeRate(context.Background(), request)
			assert.Equal(t, expected, actual)
			assert.Nil(t, err)
		})

		t.Run("Should return error for invalid response", func(t *testing.T) {
			mockHTTPClient := new(MockHTTPClient)
			mockResponse := &http.Response{
				StatusCode: 200,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(invalidJSONResponse))),
			}

			mockHTTPClient.On("Do", mock.Anything).Return(mockResponse, nil)
			client := client{
				URL:         testURL,
				HTTPCommand: mockHTTPClient,
			}

			_, err := client.GetExchangeRate(context.Background(), request)
			assert.Error(t, err)
		})

		t.Run("Should return error for http client error", func(t *testing.T) {
			mockHTTPClient := new(MockHTTPClient)
			mockHTTPClient.On("Do", mock.Anything).Return(&http.Response{},
				&url.Error{Op: "Post", URL: testURL, Err: errors.New("test error")})
			client := client{testURL, mockHTTPClient}

			_, err := client.GetExchangeRate(context.Background(), request)
			assert.Error(t, err)
		})

		t.Run("Should return error when exchange rate not found", func(t *testing.T) {
			mockHTTPClient := new(MockHTTPClient)
			mockResponse := &http.Response{
				Status:     "Rate not available!!",
				StatusCode: 400,
				Body:       ioutil.NopCloser(bytes.NewBuffer([]byte(""))),
			}

			mockHTTPClient.On("Do", mock.Anything).Return(mockResponse, nil)
			client := client{
				URL:         testURL,
				HTTPCommand: mockHTTPClient,
			}

			response, err := client.GetExchangeRate(context.Background(), request)
			assert.Nil(t, response)
			assert.Error(t, err)
		})
	})
}
