package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	/*s, err := getFileSizeFromContainer("0bd91ae1d992", "/var/hyperledger/production/ledgersData/chains/chains/mychannel/blockfile_000000")
		if err != nil {
			fmt.Printf("Geting error %s \n", err)
			os.Exit(1)
		}

		fmt.Printf("Size of the file %d ", s)
	}

	func createDockerClient() (*client.Client, error) {
		cli, err := client.NewEnvClient()
		if err != nil {
			return nil, err
		}
		return cli, nil
	}

	func getFileSizeFromContainer(containerID, filePath string) (int64, error) {
		// Create Docker client
		cli, err := createDockerClient()
		if err != nil {
			return 0, err
		}

		// Get container file stats
		stats, err := cli.ContainerStatPath(context.Background(), containerID, filePath)
		if err != nil {
			return 0, err
		}

		// Return file size
		return stats.Size, nil*/

	/*start := time.Now()
	chnl := 3
	ch := "channel" + strconv.Itoa(chnl)
	fmt.Printf("Channel Name-- %s \n", ch)

	cmd := exec.Command("/bin/bash", "/home/utsa/WorkSpace2nd/Version3/fabric-samples/test-network/channelCreate.sh")

	out, err := cmd.Output()
	if err != nil {
		fmt.Printf("The errors are %s \n", err)

	}

	fmt.Printf(string(out))
	elapsed := time.Since(start)
	fmt.Printf("Time taken to create channel and install chaincode--- %s \n", elapsed)*/

	// Set the path to the directory that contains the script
	scriptDir := "/home/utsa/WorkSpace2nd/Version3/fabric-samples/test-network"
	ch := "channel5"
	// Set the path to the script file
	scriptPath := scriptDir + "/channelCreate.sh"

	// Check if the script file exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		fmt.Println("Script file does not exist")
		return
	}

	// Set the command to execute the script
	cmd := exec.Command("/bin/bash", scriptPath, ch)

	// Set the working directory for the command
	cmd.Dir = scriptDir

	// Run the command and wait for it to finish
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running script:", err)
		return
	}

	// Print a message indicating that the script has finished executing
	fmt.Println("Script has finished executing")
}
