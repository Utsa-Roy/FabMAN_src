/*package main

import (
	"fmt"
	"os/exec"
)

/*
	type User struct {
		UID         string `json:"uid"`
		UNID        string `json:"unid"`
		UPubKey     string `json:"upubkey"`
		UserLevel   string `json:"userlevel"`
		ASLevel     string `json:"aslevel"`
		UserZone    string `json:"userzone"`
		validity    string `json:"validity"`
		UTrustLevel int    `json:"utrustlevel`
		UStatus     string `json:"status"`
	}
*/
/*func main() {
	out, err := exec.Command("/bin/sh", "/home/utsa/WorkSpace2nd/Version3/fabric-samples/fabcar/go/runfabcar.sh").Output()
	if err != nil {
		fmt.Println("Error running the script", err)
		return
	}
	//user := User{}
	//json.Unmarshal(out, &user)
	output := string(out)
	fmt.Println(output)

}*/


/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

func main() {

	type User struct {
		UID         string `json:"uid"`
		UNID        string `json:"unid"`
		UPubKey     string `json:"upubkey"`
		UserLevel   string `json:"userlevel"`
		ASLevel     string `json:"aslevel"`
		UserZone    string `json:"userzone"`
		validity    string `json:"validity"`
		UTrustLevel int    `json:"utrustlevel`
		UStatus     string `json:"status"`
	}

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

	/*start := time.Now()
	for i := 10; i < 20; i++ {
		id := "User" + strconv.Itoa(1+i)
		result, err := contract.SubmitTransaction("AddUser", "User0", id, "127.2.2.0", "sjhohsov23hjv23", "Admin", "High", "A", "N/A", "100", "Active")
		if err != nil {
			fmt.Printf("Failed to submit transaction: %s\n", err)
			os.Exit(1)
		}
		fmt.Println(string(result))
	}
	elapsed := time.Since(start)
	fmt.Printf("Time taken--- %s", elapsed)*/

	result, err := contract.EvaluateTransaction("QueryUser", "User0")
	if err != nil {
		fmt.Printf("Failed to evaluate transaction: %s\n", err)
		os.Exit(1)
	}
	user := User{}
	json.Unmarshal(result, &user)

	status := user.UStatus

	fmt.Println(string(status))

	/*result, err := contract.EvaluateTransaction("QueryAllUser")
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
