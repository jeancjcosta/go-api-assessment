package utils

import (
	"os"
	"testing"
)

func TestLoadFile(test *testing.T) {
	// Create a temporary file with test data
	tmpfile, err := os.CreateTemp("", "testfile")
	if err != nil {
		test.Fatalf("Failed to create temp file: %s", err)
	}
	defer func(name string) {
		err := os.Remove(name)
		if err != nil {
			test.Fatalf("Failed to remove temp file: %s", err)
		}
	}(tmpfile.Name()) // Clean up the temp file after the test

	// Write test data to the temporary file
	testData := "1 2 3\n4 5 6\n"
	if _, err := tmpfile.Write([]byte(testData)); err != nil {
		test.Fatalf("Failed to write to temp file: %s", err)
	}
	if err := tmpfile.Close(); err != nil {
		test.Fatalf("Failed to close temp file: %s", err)
	}

	// Call the LoadFile function with the path to the temporary file
	LoadFile(tmpfile.Name())

	// Verify the loaded data
	expectedData := []int{1, 2, 3, 4, 5, 6}
	if len(Data) != len(expectedData) {
		test.Fatalf("Expected %d elements, got %d", len(expectedData), len(Data))
	}
	for i, value := range Data {
		if value != expectedData[i] {
			test.Errorf("Expected %d at index %d, got %d", expectedData[i], i, value)
		}
	}
}
