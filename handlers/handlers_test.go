package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"go-api-assessment/utils"
)

func TestSearchHandler(test *testing.T) {
	// Initialize the data for the test
	utils.Data = []int{0, 10, 20, 100, 1000000}

	tests := []struct {
		name           string
		value          string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Exact match",
			value:          "100",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"value":100,"index":3}`,
		},
		{
			name:           "Within 10% range",
			value:          "110",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"value":110,"index":3,"message":"Value not found, returning closest match"}`,
		},
		{
			name:           "Outside 10% range",
			value:          "200",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"value":200,"index":4,"message":"Value not found. No values within a 10% margin were located"}`,
		},
		{
			name:           "Invalid value",
			value:          "invalid",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   "Invalid value",
		},
		{
			name:           "Value below range",
			value:          "-10",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"value":-10,"index":0,"message":"Value not found, returning closest match"}`,
		},
		{
			name:           "Value above range",
			value:          "2000000",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"value":2000000,"index":5,"message":"Value not found, returning closest match"}`,
		},
	}

	for _, testCase := range tests {
		test.Run(testCase.name, func(t *testing.T) {
			// Create a new HTTP GET request for the endpoint
			req, err := http.NewRequest("GET", "/endpoint/"+testCase.value, nil)
			if err != nil {
				test.Fatal(err) // Fail the test if the request creation fails
			}

			// Create a ResponseRecorder to record the response
			rr := httptest.NewRecorder()

			// Create a new router and register the handler
			router := mux.NewRouter()
			router.HandleFunc("/endpoint/{value}", SearchHandler)

			// Serve the HTTP request
			router.ServeHTTP(rr, req)

			// Check if the status code is as expected
			if status := rr.Code; status != testCase.expectedStatus {
				test.Errorf("handler returned wrong status code: got %v want %v", status, testCase.expectedStatus)
			}

			// Compare the actual response body with the expected response body
			if strings.ReplaceAll(rr.Body.String(), "\n", "") != testCase.expectedBody {
				test.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), testCase.expectedBody)
			}
		})
	}
}
