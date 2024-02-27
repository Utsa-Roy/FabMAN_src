/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"

	"github.com/docker/docker/client"
	"golang.org/x/net/context"
)

func main() {

	os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		fmt.Printf("Failed to create wallet: %s\n", err)
		os.Exit(1)
	}

	if !wallet.Exists("appUser") {
		err = populateWallet(wallet)
		if err != nil {
			fmt.Printf("Failed to populate wallet contents: %s\n", err)
			os.Exit(1)
		}
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, "appUser"),
	)
	if err != nil {
		fmt.Printf("Failed to connect to gateway: %s\n", err)
		os.Exit(1)
	}
	defer gw.Close()

	network, err := gw.GetNetwork("mychannel")
	if err != nil {
		fmt.Printf("Failed to get network: %s\n", err)
		os.Exit(1)
	}

	contract := network.GetContract("fabcar")

	start := time.Now()
	n := 0
	chnl := 1
	ch := "mychannel"
	for j := 0; j < 60; j++ {
		for i := 0; i <= 1000; i++ {

			id := "Request" + strconv.Itoa(n)

			go contract.SubmitTransaction("AccessRequestVerifier", id, "U2D", "Read", "User0", "Device1", "5")
			n++

		}

		id := "Request" + strconv.Itoa(n+1)
		result, err := contract.SubmitTransaction("AccessRequestVerifier", id, "U2D", "Read", "User0", "Device1", "5")
		n++

		if err != nil {
			fmt.Printf("Failed to submit transaction user: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(string(result))

		////////////////////////////////////////////////////////

		elapsed := time.Since(start)
		fmt.Printf("Time--- %s \n", elapsed)

		s, err := getFileSizeFromContainer("3557e594e49f", "/var/hyperledger/production/ledgersData/chains/chains/"+ch+"/blockfile_000000")
		if err != nil {
			fmt.Printf("Geting error %s \n", err)
			os.Exit(1)
		}
		fmt.Printf(" Size of the file %d \n", s)

		if s >= 52428800 {

			_, err1 := newChannel(chnl)
			if err != nil {
				fmt.Printf("Failed to Create Channel: %s\n", err1)
				os.Exit(1)
			}

			ch = "channel" + strconv.Itoa(chnl)
			network, err = gw.GetNetwork(ch)
			if err != nil {
				fmt.Printf("Failed to get network: %s\n", err)
				os.Exit(1)
			}
			contract = network.GetContract("fabcar")
			chnl++

		}

		time.Sleep(60 * time.Second)
	}

	//////////////////////////////////////////////////////////////

	elapsed := time.Since(start)
	fmt.Printf("Time taken--- %s", elapsed)

	/*result, err := contract.EvaluateTransaction("QueryAllAccessRequest")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Printf("Ledger Size --")
	fmt.Println(string(result))

	/*result, err := contract.EvaluateTransaction("QueryUser", "User0")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	user := User{}
	json.Unmarshal(result, &user)

	status := user.UStatus

	fmt.Println(string(status))

	result, err := contract.EvaluateTransaction("QueryAllUser")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))

	result, err = contract.SubmitTransaction("createCar", "CAR10", "VW", "Polo", "Grey", "Mary")
	if err != nil {
		fmt.Printf("Failed to submit transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))

	result, err = contract.EvaluateTransaction("queryCar", "CAR10")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))

	_, err = contract.SubmitTransaction("changeCarOwner", "CAR10", "Archie")
	if err != nil {
		fmt.Printf("Failed to submit transaction: %s\n", err)
		os.Exit(1)
	}

	result, err = contract.EvaluateTransaction("queryCar", "CAR10")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))*/
}

func populateWallet(wallet *gateway.Wallet) error {
	credPath := filepath.Join(
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := os.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := os.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := os.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity("Org1MSP", string(cert), string(key))

	err = wallet.Put("appUser", identity)
	if err != nil {
		return err
	}
	return nil
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

	return stats.Size, nil
}
func newChannel(chnl int) (int, error) {
	start := time.Now()

	scriptDir := "/home/utsa/WorkSpace2nd/Version3/fabric-samples/test-network"
	ch := "channel" + strconv.Itoa(chnl)
	// Set the path to the script file
	scriptPath := scriptDir + "/channelCreate.sh"

	// Check if the script file exists
	if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
		fmt.Println("Script file does not exist")
		return 0, err
	}

	// Set the command to execute the script
	cmd := exec.Command("/bin/bash", scriptPath, ch)

	// Set the working directory for the command
	cmd.Dir = scriptDir

	// Run the command and wait for it to finish
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error running script:", err)
		return 0, err
	}

	// Print a message indicating that the script has finished executing

	elapsed := time.Since(start)
	fmt.Printf("Time taken to create channel and install chaincode--- %s \n", elapsed)
	return 1, nil
}
