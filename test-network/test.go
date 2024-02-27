package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// Specify the file path
	filePath := "/home/utsa/WorkSpace2nd/Version3/fabric-samples/test-network-nano-bash/data/peer0.org1.example.com/ledgersData/chains/chains/mychannel/blockfile_000000"

	// Open the file in read-only mode
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Get file information
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	// Get the file size in bytes
	fileSize := fileInfo.Size()

	// Print the file size
	fmt.Printf("File size: %d bytes\n", fileSize)
}
