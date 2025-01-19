package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var Data []int

func LoadFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil {
				log.Fatalf("Failed to convert value to int: %s", err)
			}
			Data = append(Data, num)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %s", err)
	}
}
