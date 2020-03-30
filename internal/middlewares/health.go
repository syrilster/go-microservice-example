package middlewares

import (
	"github.com/syrilster/go-microservice-example/internal/util"
	"net/http"
)

func RuntimeHealthCheck() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		util.WithBodyAndStatus("All OK", http.StatusOK, w)
	}
}
