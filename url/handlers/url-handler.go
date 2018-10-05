package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/asahasrabuddhe/go-api-base/response"
	"github.com/asahasrabuddhe/url-shortener/url/datastore"
	"github.com/asahasrabuddhe/url-shortener/url/generator"
	"github.com/asahasrabuddhe/url-shortener/url/requests"
	"net/http"
)

func ShrinkURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	var req requests.ShrinkURLRequest
	var res response.Response

	decoder.Decode(&req)

	if err := req.Validate(); err != nil {
		// Error
		w.WriteHeader(http.StatusUnprocessableEntity)

		res = response.Response{
			Success: false,
			Message: err.Error(),
		}
	} else {
		// Logic
		key := generator.DefaultGenerator()
		// Ensure key is not already present in the bucket
		for {
			if v := datastore.DB.Get(key); v == "" {
				break
			}
		}
		key = generator.DefaultGenerator()
		// Store key and url in the bucket
		datastore.DB.Set(key, req.URL)
	}

	encoder.Encode(&res)
}

func ResolveURL(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, r.UserAgent())
	fmt.Fprintln(w, r.Host)
}
