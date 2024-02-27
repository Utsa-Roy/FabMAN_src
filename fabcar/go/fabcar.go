/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"errors"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"
	"time"
	"io/ioutil"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
	//"github.com/docker/docker/client"
	//"golang.org/x/net/context"
)

func main() {
	n := 0
	//a := 100.0
	//b := 10.0
	//c := 0.012

	//latestBlock := 0
	//for i := 0; i < 10; i++ {
	start := time.Now()
	/*for t := 0; t < 600; t++ {

		//tno := int(gompertz(a, b, c, float64(t)))

		//for i := 0; i < tno; i++ {
			k := 1000 / tno
			go sendTxn(n)
			n++
			time.Sleep(time.Duration(k) * time.Millisecond)
		}
		fmt.Printf("\n-------------------------Time %d----------------------\n", t)

	}*/

	for i := 0; i < 10; i++ {
		for j := 0; j < 100; j++ {
			go sendTxn(n)
			n++

		}
		time.Sleep(1 * time.Second)
		fmt.Printf("\n-------------------------Time %d Sec----------------------\n", i)
	}
	sendTxn(n)
	elapsed := time.Since(start)
	fmt.Printf("--------------%s  ---------------\n", elapsed)
	//}

	for {

	}

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
	//cert, err := os.ReadFile(filepath.Clean(certPath))
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))

	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	//files, err := os.ReadDir(keyDir)
	fileInfo, err := ioutil.ReadDir(keyDir)
	if err != nil {
    		return err
	}
	if len(fileInfo) != 1 {
    		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, fileInfo[0].Name())

	/*if err != nil {
		return err
	}
	if len(files) != 1 {
		return errors.New("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())*/
	//key, err := os.ReadFile(filepath.Clean(keyPath))
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))

	
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

func gompertz(a, b, c, t float64) float64 {
	return 1 + a*math.Exp(-b*math.Exp(-c*t))
}

/*

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
*/
