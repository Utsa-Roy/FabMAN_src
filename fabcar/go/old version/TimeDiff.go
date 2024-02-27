package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	// Set the path to the directory that contains the script
	scriptDir := "/home/utsa/WorkSpace2nd/Version3/fabric-samples/test-network/Adaptive"
	// Set the path to the script file
	scriptPath := scriptDir + "/S1.sh"

	BlockNumber1 := "10"
	BlockNumber2 := "11"

	BlockName1 := "mychannel_" + BlockNumber1 + ".block"
	BlockName2 := "mychannel_" + BlockNumber2 + ".block"

	file1 := "mychannel_" + BlockNumber1 + ".json"
	file2 := "mychannel_" + BlockNumber2 + ".json"

	// Check if the script file exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		fmt.Println("Script file does not exist")
		return
	}

	// Set the command to execute the script
	cmd := exec.Command("/bin/bash", scriptPath, BlockNumber1, BlockNumber2, BlockName1, BlockName2, file1, file2)

	// Set the working directory for the command
	cmd.Dir = scriptDir

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running script:", err)
		return
	}

	// Convert the output to a string and trim any whitespace
	outputStr := strings.TrimSpace(string(output))

	// Print the time difference
	fmt.Println("Time difference:", outputStr, "milliseconds")

	txnNo, err := countTxn()
	if err != nil {
		fmt.Printf("Error Reading the Block %s \n", err)
		return
	}
	fmt.Printf("Number of transaction on the Block %s \n", txnNo)

}

func countTxn() (string, error) {
	// Set the path to the directory that contains the script
	scriptDir := "/home/utsa/WorkSpace2nd/Version3/fabric-samples/test-network/Adaptive"
	// Set the path to the script file
	scriptPath := scriptDir + "/txnNo.sh"

	BlockNumber1 := "10"

	file1 := "mychannel_" + BlockNumber1 + ".json"

	// Check if the script file exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		fmt.Println("Script file does not exist")
		return "null", err
	}

	// Set the command to execute the script
	cmd := exec.Command("/bin/bash", scriptPath, file1)

	// Set the working directory for the command
	cmd.Dir = scriptDir

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error running script:", err)
		return "null", err
	}

	// Convert the output to a string and trim any whitespace
	outputStr := strings.TrimSpace(string(output))

	return outputStr, nil
}
