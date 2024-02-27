package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	// Set the parameters
	blockNumber := "1"
	blockFileName := "block.pb"
	decodedBlockFileName := "block.json"

	// Fetch the block
	cmdFetch := exec.Command("peer", "channel", "fetch", blockNumber, blockFileName, "-c", "<mychannel->")
	if err := cmdFetch.Run(); err != nil {
		log.Fatal(err)
	}

	// Decode the block
	cmdDecode := exec.Command("configtxlator", "proto_decode", "--input", blockFileName, "--type", "common.Block", "--output", decodedBlockFileName)
	if err := cmdDecode.Run(); err != nil {
		log.Fatal(err)
	}

	// Count the transactions
	decodedBlockFile, err := os.Open(decodedBlockFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer decodedBlockFile.Close()

	decodedBlockBytes, err := ioutil.ReadAll(decodedBlockFile)
	if err != nil {
		log.Fatal(err)
	}

	decodedBlockString := string(decodedBlockBytes)
	transactionCountString := strings.TrimSpace(execCommand(fmt.Sprintf(`echo '%s' | jq '.data.data | length'`, decodedBlockString)))
	transactionCount, err := strconv.Atoi(transactionCountString)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The block number %s contains %d transactions.\n", blockNumber, transactionCount)
}

func execCommand(command string) string {
	cmd := exec.Command("bash", "-c", command)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(out)
}
