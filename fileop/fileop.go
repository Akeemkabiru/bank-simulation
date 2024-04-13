package fileop

import (
	"fmt"
	"os"
	"strconv"
)

func WriteFile(fileName string, value float64) {
	err := os.WriteFile(fileName, []byte(fmt.Sprintf("%v", value)), 0644)
	if err != nil {
		fmt.Printf("error writing file")
		os.Exit(1)

	}
}

func ReadFile(fileName string) float64 {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("error writing file")
		os.Exit(1)

	}
	floatValue, err := strconv.ParseFloat(string(data), 64)
	return floatValue
}
