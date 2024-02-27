package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// Set the parameters
	blockNumber1 := "9"
	blockNumber2 := "10"

	// Get the timestamps of the blocks
	timestamp1, err := getBlockTimestamp(blockNumber1)
	if err != nil {
		log.Fatal(err)
	}

	timestamp2, err := getBlockTimestamp(blockNumber2)
	if err != nil {
		log.Fatal(err)
	}

	// Calculate the time gap
	timeGap := timestamp2 - timestamp1

	fmt.Printf("The time gap between block %s and block %s is %d seconds.\n", blockNumber1, blockNumber2, timeGap)
}

func getBlockTimestamp(blockNumber string) (int64, error) {
	// Fetch the block
	blockFileName := fmt.Sprintf("block_%s.pb", blockNumber)
	cmdFetch := exec.Command("peer", "channel", "fetch", blockNumber, blockFileName, "-c", "mychannel")
	if err := cmdFetch.Run(); err != nil {
		return 0, err
	}

	// Decode the block
	decodedBlockFileName := fmt.Sprintf("block_%s.json", blockNumber)
	cmdDecode := exec.Command("configtxlator", "proto_decode", "--input", blockFileName, "--type", "common.Block", "--output", decodedBlockFileName)
	if err := cmdDecode.Run(); err != nil {
		return 0, err
	}

	// Get the timestamp
	decodedBlockFile, err := exec.Command("cat", decodedBlockFileName).Output()
	if err != nil {
		return 0, err
	}

	decodedBlockString := string(decodedBlockFile)
	timestampString := strings.TrimSpace(execCommand(fmt.Sprintf(`echo '%s' | jq '.header.channel_header.timestamp'`, decodedBlockString)))
	timestampString = strings.Trim(timestampString, `"`)
	timestamp, err := strconv.ParseInt(timestampString, 10, 64)
	if err != nil {
		return 0, err
	}

	return timestamp, nil
}

func execCommand(command string) string {
	cmd := exec.Command("bash", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
