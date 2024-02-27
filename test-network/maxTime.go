package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type DataItem struct {
	Timestamp string `json:"timestamp"`
	// Add other fields as needed
}

type Block20Data struct {
	Data []struct {
		Data []DataItem `json:"data"`
	} `json:"data"`
}

func main() {
	// Open and read the JSON file
	file, err := os.Open("block20.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Parse JSON data
	var blockData Block20Data
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&blockData)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Access the timestamp at the desired location
	index := 10
	if index < len(blockData.Data[0].Data) {
		timestampStr := blockData.Data[0].Data[index].Timestamp

		// Parse timestamp string into time.Time
		timestamp, err := time.Parse(time.RFC3339Nano, timestampStr)
		if err != nil {
			fmt.Println("Error parsing timestamp:", err)
			return
		}

		// Calculate time difference
		currentTime := time.Now().UTC()
		timeDifference := currentTime.Sub(timestamp)

		fmt.Printf("Timestamp: %s\n", timestamp)
		fmt.Printf("Time difference: %s\n", timeDifference)
	} else {
		fmt.Println("Index out of range")
	}
}
