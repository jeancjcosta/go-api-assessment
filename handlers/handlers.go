package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"go-api-assessment/models"
	"go-api-assessment/utils"
)

// SearchHandler handles the search requests
// @Summary Search for a value
// @Description Get a value by its identifier
// @Tags search
// @Param value path string true "Value Identifier"
// @Success 200 {object} models.Response
// @Failure 400 {object} string
// @Failure 404 {object} string
// @Router /endpoint/{value} [get]
func SearchHandler(w http.ResponseWriter, r *http.Request) {
	// Extract the value from the URL path variables
	vars := mux.Vars(r)
	value, err := strconv.Atoi(vars["value"])
	if err != nil {
		// If the value is not a valid integer, return a 400 Bad Request error
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	// Perform a binary search on the data to find the index of the value
	index, found := binarySearch(utils.Data, value)
	response := models.Response{
		Value: value,
		Index: index,
	}

	// If the value is not found, set a message indicating the closest match is returned
	if !found && index != -1 {
		response.Message = "Value not found, returning closest match"
	}

	// If the value is not found, set a message indicating the closest match is returned
	if !found && index == -1 {
		response.Message = "Value not found. No values within a 10% margin were located"
	}

	// Set the response content type to JSON
	w.Header().Set("Content-Type", "application/json")
	// Encode the response as JSON and write it to the response writer
	json.NewEncoder(w).Encode(response)
}

func binarySearch(data []int, target int) (int, bool) {
	left, right := 0, len(data)-1
	// Perform binary search
	for left <= right {
		mid := left + (right-left)/2
		if data[mid] == target {
			// If the target is found, return the index and true
			return mid, true
		} else if data[mid] < target {
			// If the target is greater, ignore the left half
			left = mid + 1
		} else {
			// If the target is smaller, ignore the right half
			right = mid - 1
		}
	}

	// Check if the closest values are within 10% of the target
	if left < len(data) && float64(data[left]) <= float64(target)*1.1 {
		return left, false
	}
	if right >= 0 && float64(data[right]) >= float64(target)*0.9 {
		return right, false
	}

	// If the target is not found, and the closest match is above the threshold, return -1 and false
	return -1, false
}
