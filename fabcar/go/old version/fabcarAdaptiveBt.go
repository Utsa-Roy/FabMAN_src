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
	"strings"
	"time"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	//"github.com/docker/docker/client"
	//"golang.org/x/net/context"
)

func main() {
	n := 1412
	latestBlock := 0
	//for i := 0; i < 10; i++ {
	start := time.Now()
	for j := 0; j < 300; j++ {
		go sendTxn(n)
		n++
		time.Sleep(200 * time.Millisecond)

		/////////////////////////////////////////////////////

		scriptDir := "/home/utsa/WorkSpace2nd/Version3/fabric-samples/test-network"
		scriptPath := scriptDir + "/bn.sh"
		if _, err := os.Stat(scriptPath); os.IsNotExist(err) {
			fmt.Println("Script file does not exist")
			return
		}

		// Set the command to execute the script
		cmd := exec.Command("/bin/bash", scriptPath)

		// Set the working directory for the command
		cmd.Dir = scriptDir

		// Run the command and capture the output
		output, err := cmd.Output()
		if err != nil {
			fmt.Println("Error running script:", err)
			return
		}

		trimmedOutput := strings.TrimSpace(string(output))
		blockNumber, err := strconv.Atoi(trimmedOutput)
		if err != nil {
			fmt.Println("Error converting to integer:", err)
			return
		}
		if blockNumber > latestBlock {
			latestBlock = blockNumber

			fmt.Printf("\n NewBlock Number %d\n", latestBlock)

			if latestBlock == 33 {
				break
			}

		}

		////////////////////////////////////////////////////
	}
	elapsed := time.Since(start)
	fmt.Printf("--------------%s  ---------------\n", elapsed)
	//}

	for {

	}

	/*s, err := getFileSizeFromContainer("d2efda028fc2", "/var/hyperledger/production/ledgersData/chains/chains/mychannel/blockfile_000000")
	if err != nil {
		fmt.Printf("Geting error %s \n", err)
		os.Exit(1)
	}

	if s >= 67000000 {
		s1, err1 := getFileSizeFromContainer("d2efda028fc2", "/var/hyperledger/production/ledgersData/chains/chains/mychannel/blockfile_000001")
		if err1 != nil {
			fmt.Printf("Geting error %s \n", err1)
		} else {
			fmt.Printf(" 2nd Size of the file %d ", s1)
		}

		if s1 >= 67000000 {
			s2, err2 := getFileSizeFromContainer("d2efda028fc2", "/var/hyperledger/production/ledgersData/chains/chains/mychannel/blockfile_000002")
			if err2 != nil {
				fmt.Printf("Geting error %s \n", err2)
			} else {
				fmt.Printf(" 3rd Size of the file %d ", s2)
			}

			if s2 >= 67000000 {
				s3, err3 := getFileSizeFromContainer("d2efda028fc2", "/var/hyperledger/production/ledgersData/chains/chains/mychannel/blockfile_000003")
				if err3 != nil {
					fmt.Printf("Geting error %s \n", err3)
				} else {
					fmt.Printf(" 4th file Size of the file %d ", s3)
				}

				if s3 >= 67000000 {
					s4, err4 := getFileSizeFromContainer("d2efda028fc2", "/var/hyperledger/production/ledgersData/chains/chains/mychannel/blockfile_000004")
					if err4 != nil {
						fmt.Printf("Geting error %s \n", err4)
					} else {
						fmt.Printf(" 5th file Size of the file %d ", s4)
					}
				}

			}
		}
	}

	fmt.Printf(" Size of the file %d ", s)
	*/

	//////////////////////////////////////////////////////////////

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

func sendTxn(txnNo int) {
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
	n := txnNo

	id := strconv.Itoa(n)

	result, err := contract.SubmitTransaction("AccessRequestVerifier", id, "U2D", "Read", "User0", "Device1", "5")

	if err != nil {
		fmt.Printf("Failed to submit transaction user: %s\n", err)
		os.Exit(1)
	}
	fmt.Println(string(result))

	elapsed := time.Since(start)
	fmt.Printf("Txn No- %s Time--- %s", id, elapsed)
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

/*
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
*/
