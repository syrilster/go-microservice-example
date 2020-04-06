package currencyconversion

import (
	"fmt"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/syrilster/go-microservice-example/internal/util"
	"net/http"
	"strconv"
)

func handler(rateFetcher ExchangeRateFetcher) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		contextLogger := log.WithContext(ctx)

		request := parseFetchRequest(r)

		calculatedAmount, err := rateFetcher.FetchExchangeRate(ctx, request)
		if err != nil {
			contextLogger.WithError(err).Error("Failed to fetch currency exchange rates")
			util.WithBodyAndStatus(nil, http.StatusInternalServerError, w)
			return
		}
		util.WithBodyAndStatus(calculatedAmount, http.StatusOK, w)
	}
}

func parseFetchRequest(r *http.Request) Request {
	fromCurrency := mux.Vars(r)["from"]
	toCurrency := mux.Vars(r)["to"]
	q := mux.Vars(r)["quantity"]
	quantity, err := strconv.ParseFloat(q, 2)
	if err != nil {
		fmt.Printf("invalid quantity, err: %v", err)
	}
	return Request{
		FromCurrency: fromCurrency,
		ToCurrency:   toCurrency,
		Quantity:     quantity,
	}
}
