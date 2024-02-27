package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

// Define the User structure, with 8 attributes.
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

// Define the Device structure, with 9 attributes.
type Device struct {
	DID         string `json:"did"`
	DNID        string `json:"dnid"`
	DPubKey     string `json:"dpubkey"`
	DType       string `json:""dtype`
	SLevel      string `json:"slevel"`
	Dzone       string `json:"Dzone"`
	TATimeStart int    `json:"tatimestart"`
	TATimeEnd   int    `json:"tatimeend"`
	DTrustLevel int    `json:"dtrustlevel"`
	DStatus     string `json:"dstatus"`
}

// Define the Request Structure, With 6 Properties
type Request struct {
	RType       string `json:"type"`
	ActionType  string `json:"actiontype"`
	RequesterID string `json:"rid"`
	DeviceID    string `json:"did"`
	Time        int    `json:"time"`
	Permission  string `json:"permission"`
}

// define 3 counter to keep track of the user, device, and request number.
type EntryCounter struct {
	UserCount    int `json:"usercount"`
	DeviceCount  int `json:"devicecount"`
	RequestCount int `json:"requestcount"`
}

type QueryResultU struct {
	Key    string `json:"Key"`
	Record *User
}

type QueryResultD struct {
	Key    string `json:"Key"`
	Record *Device
}

type QueryResultR struct {
	Key    string `json:"Key"`
	Record *Request
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	request := []Request{
		Request{RType: "U2D", ActionType: "Read", RequesterID: "User1", DeviceID: "Device2", Time: 10, Permission: "ALLOW"},
		Request{RType: "U2D", ActionType: "Read", RequesterID: "User1", DeviceID: "Device2", Time: 10, Permission: "ALLOW"},
	}

	i := 0
	for i < len(request) {
		requestAsBytes, _ := json.Marshal(request[i])
		err := ctx.GetStub().PutState("Request"+strconv.Itoa(i), requestAsBytes)
		fmt.Println("Added", request[i])
		i = i + 1

		if err != nil {
			return fmt.Errorf("Failed to put to world state. %s", err.Error())
		}
	}

	return nil
}

func (s *SmartContract) AccessRequestVerifier(ctx contractapi.TransactionContextInterface, args0 string, args1 string, args2 string, args3 string, args4 string, args5 string) error {

	//n, _ := strconv.ParseInt(args0, 10, 64)

	//time.Sleep(time.Duration(n) * 966 * time.Microsecond)

	if args1 == "U2D" {
		var tm, _ = strconv.Atoi(args5)
		id := "Request" + args0
		var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "ALLOW"}
		requestAsBytes, _ := json.Marshal(request)
		ctx.GetStub().PutState(id, requestAsBytes)

		return nil
	} else {
		var tm, _ = strconv.Atoi(args5)
		var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "ALLOW"}
		requestAsBytes, _ := json.Marshal(request)
		ctx.GetStub().PutState(args0, requestAsBytes)

		return nil
	}

}

// accessRequest function is the access control policy
/*func (s *SmartContract) AccessRequestVerifier(ctx contractapi.TransactionContextInterface, args0 string, args1 string, args2 string, args3 string, args4 string, args5 string, args6 string, args7 string, args8 string, args9 string, args10 string, args11 string, args12 string, args13 string, args14 string, args15 string, args16 string, args17 string, args18 string, args19 string, args20 string, args21 string, args22 string) error {

	if args1 == "U2D" {
		/*userAsBytes, err := queryUsrWallet(args3) //ctx.GetStub().GetState(args3) // ******Here......query*********
		if err != nil {
			return err
		}
		user := userAsBytes
		//json.Unmarshal(userAsBytes, &user)
		var user *User

		user.UID = args3
		user.UNID = args6
		user.UPubKey = args7
		user.UserLevel = args8
		user.ASLevel = args9
		user.UserZone = args10
		user.validity = args11
		user.UTrustLevel, _ = strconv.Atoi(args12)
		user.UStatus = args13

		/*deviceAsBytes, err := queryDevWallet(args4) //ctx.GetStub().GetState(args4) // *****Here......query*********
		if err != nil {
			return err
		}
		device := deviceAsBytes

		var device *Device

		device.DID = args4
		device.DNID = args14
		device.DPubKey = args15
		device.DType = args16
		device.SLevel = args17
		device.Dzone = args18
		device.TATimeStart, _ = strconv.Atoi(args19)
		device.TATimeEnd, _ = strconv.Atoi(args20)
		device.DTrustLevel, _ = strconv.Atoi(args21)
		device.DStatus = args22rid

		if args2 == "Read" && (device.DType == "Sensor" || device.DType == "Both") {
			if user.UserLevel == "Admin" && user.UNID == device.DNID {

				var tm, _ = strconv.Atoi(args5)
				var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "ALLOW"}
				requestAsBytes, _ := json.Marshal(request)
				ctx.GetStub().PutState(args0, requestAsBytes)

				return nil
			} else if user.UserLevel == "Guest" && user.validity == "Not valid" {

				var tm, _ = strconv.Atoi(args5)
				var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "DENY"}
				requestAsBytes, _ := json.Marshal(request)
				ctx.GetStub().PutState(args0, requestAsBytes)

				return nil
			} else {
				if user.UserZone == device.Dzone {

					var tm, _ = strconv.Atoi(args5)
					var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "ALLOW"}
					requestAsBytes, _ := json.Marshal(request)
					ctx.GetStub().PutState(args0, requestAsBytes)

					return nil
				} else {
					var time, _ = strconv.Atoi(args5)
					if (device.TATimeStart <= time) && (time <= device.TATimeEnd) {

						var tm, _ = strconv.Atoi(args4)
						var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "ALLOW"}
						requestAsBytes, _ := json.Marshal(request)
						ctx.GetStub().PutState(args0, requestAsBytes)

						return nil
					} else {

						var tm, _ = strconv.Atoi(args5)
						var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "DENY"}
						requestAsBytes, _ := json.Marshal(request)
						ctx.GetStub().PutState(args0, requestAsBytes)

						return nil
					}
				}

			}

		} else if args2 == "Action" && (device.DType == "Actuator" || device.DType == "Both") {
			if user.UserLevel == "Admin" && user.UNID == device.DNID {

				var tm, _ = strconv.Atoi(args5)
				var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "ALLOW"}
				requestAsBytes, _ := json.Marshal(request)
				ctx.GetStub().PutState(args0, requestAsBytes)

				return nil
			} else if user.UserLevel == "Guest" && user.validity == "Not valid" {
				var tm, _ = strconv.Atoi(args5)
				var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "DENY"}
				requestAsBytes, _ := json.Marshal(request)
				ctx.GetStub().PutState(args0, requestAsBytes)

				return nil
			} else {
				if user.UserZone == device.Dzone {

					var tm, _ = strconv.Atoi(args5)
					var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "ALLOW"}
					requestAsBytes, _ := json.Marshal(request)
					ctx.GetStub().PutState(args0, requestAsBytes)

					return nil
				} else {
					var time, _ = strconv.Atoi(args5)
					if (device.TATimeStart <= time) && (time <= device.TATimeEnd) {
						if device.SLevel == user.ASLevel {

							var tm, _ = strconv.Atoi(args5)
							var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "ALLOW"}
							requestAsBytes, _ := json.Marshal(request)
							ctx.GetStub().PutState(args0, requestAsBytes)

							return nil
						}
					} else {

						var tm, _ = strconv.Atoi(args5)
						var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "DENY"}
						requestAsBytes, _ := json.Marshal(request)
						ctx.GetStub().PutState(args0, requestAsBytes)

						return nil
					}
				}

			}

		} else {

			var tm, _ = strconv.Atoi(args5)
			var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "DENY"}
			requestAsBytes, _ := json.Marshal(request)
			ctx.GetStub().PutState(args0, requestAsBytes)

			return nil
		}

	} else if args1 == "D2D" {

		rdevice := Device{}

		rdevice.DID = args3
		rdevice.DNID = args6
		rdevice.DPubKey = args7
		rdevice.DType = args8
		rdevice.SLevel = args9
		rdevice.Dzone = args10
		rdevice.TATimeStart, _ = strconv.Atoi(args19)
		rdevice.TATimeEnd, _ = strconv.Atoi(args20)
		rdevice.DTrustLevel = 100
		rdevice.DStatus = args13

		device := Device{}

		device.DID = args4
		device.DNID = args14
		device.DPubKey = args15
		device.DType = args16
		device.SLevel = args17
		device.Dzone = args18
		device.TATimeStart, _ = strconv.Atoi(args19)
		device.TATimeEnd, _ = strconv.Atoi(args20)
		device.DTrustLevel, _ = strconv.Atoi(args21)
		device.DStatus = args22

		if (args2 == "Read" && (device.DType == "Sensor" || device.DType == "Both")) || (args2 == "Action" && (device.DType == "Actuator" || device.DType == "Both")) {
			if rdevice.DNID == device.DNID {
				if rdevice.Dzone == device.Dzone {
					var tm, _ = strconv.Atoi(args5)
					var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "ALLOW"}
					requestAsBytes, _ := json.Marshal(request)
					ctx.GetStub().PutState(args0, requestAsBytes)

					return nil

				} else {
					var time, _ = strconv.Atoi(args5)
					if (device.TATimeStart <= time) && (time <= device.TATimeEnd) {
						if rdevice.SLevel == device.SLevel {

							var tm, _ = strconv.Atoi(args5)
							var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "ALLOW"}
							requestAsBytes, _ := json.Marshal(request)
							ctx.GetStub().PutState(args0, requestAsBytes)

							return nil

						}
					} else {

						var tm, _ = strconv.Atoi(args5)
						var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "DENY"}
						requestAsBytes, _ := json.Marshal(request)
						ctx.GetStub().PutState(args0, requestAsBytes)

						return nil

					}
				}
			} else {

				var tm, _ = strconv.Atoi(args5)
				var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "DENY"}
				requestAsBytes, _ := json.Marshal(request)
				ctx.GetStub().PutState(args0, requestAsBytes)

				return nil
			}
		} else {

			var tm, _ = strconv.Atoi(args4)
			var request = Request{RType: args1, ActionType: args2, RequesterID: args3, DeviceID: args4, Time: tm, Permission: "DENY"}
			requestAsBytes, _ := json.Marshal(request)
			ctx.GetStub().PutState(args0, requestAsBytes)

			return nil
		}

	}
	return fmt.Errorf("Error!!!! Unsupported Request")
} */

//end  accessRequest function

// update Trust level function
/*func (s *SmartContract) TrustLevelUpdater(ctx contractapi.TransactionContextInterface, args0 string, args1 string, args2 string) error {

	requestAsBytes, _ := ctx.GetStub().GetState(args0)
	request := Request{}
	json.Unmarshal(requestAsBytes, &request)
	var RID = request.RequesterID
	var DID = request.DeviceID

	userAsBytes, _ := ctx.GetStub().GetState(RID)
	user := User{}
	json.Unmarshal(userAsBytes, &user)
	var um = user.UTrustLevel

	deviceAsBytes, _ := ctx.GetStub().GetState(DID)
	device := Device{}
	json.Unmarshal(deviceAsBytes, &device)
	var dm = device.DTrustLevel
	if args1 == "Satisfactory" && um < 100 {
		um = um + 1
		user.UTrustLevel = um
		userAsBytes, _ = json.Marshal(user)
		ctx.GetStub().PutState(RID, userAsBytes)
	} else {
		um = um - 1

		//if um <= 0 {
		//		var ip[] string = {RID}
		//		s.deleteUser(APIstub, ip)
		//	}
	}
	if args2 == "Satisfactory" && dm < 100 {
		dm = dm + 1
		device.DTrustLevel = dm
		deviceAsBytes, _ = json.Marshal(device)
		ctx.GetStub().PutState(args1, deviceAsBytes)
	} else {
		dm = dm - 1
		device.DTrustLevel = dm
		deviceAsBytes, _ = json.Marshal(device)
		ctx.GetStub().PutState(args1, deviceAsBytes)
		///if dm <= 0 {
		//	var ip[] string = {DID}
		//	s.deleteDevice(APIstub, ip)
		//
		//		}
	}

	userAsBytes, _ = json.Marshal(user)
	ctx.GetStub().PutState(args1, userAsBytes)

	return nil
}
*/
// query Request
func (s *SmartContract) QueryAccessRequest(ctx contractapi.TransactionContextInterface, rid string) (*Request, error) {

	requestAsBytes, err := ctx.GetStub().GetState(rid)

	if err != nil {
		return nil, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if requestAsBytes == nil {
		return nil, fmt.Errorf("%s does not exist", rid)
	}

	request := new(Request)
	_ = json.Unmarshal(requestAsBytes, request)

	return request, nil
}

// QueryAllAccessRequest returns all Request found in world state
func (s *SmartContract) QueryAllAccessRequest(ctx contractapi.TransactionContextInterface) ([]QueryResultR, error) {
	startKey := "Request0"
	endKey := "Request9999"

	resultsIterator, err := ctx.GetStub().GetStateByRange(startKey, endKey)

	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	results := []QueryResultR{}

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()

		if err != nil {
			return nil, err
		}

		request := new(Request)
		_ = json.Unmarshal(queryResponse.Value, request)

		queryResultR := QueryResultR{Key: queryResponse.Key, Record: request}
		results = append(results, queryResultR)
	}

	return results, nil
}

func adminOrNot(args string) bool {

	/*userAsBytes, _ := APIstub.GetState(args)
	user := User{}
	json.Unmarshal(userAsBytes, &user)
	if user.UserLevel == "Admin" {
		return true
	} else {
		return false
	}*/
	return true
}

func (s *SmartContract) QueryPermission(ctx contractapi.TransactionContextInterface, rid string) (string, error) {

	requestAsBytes, err := ctx.GetStub().GetState(rid)
	var rse = "DENY"
	if err != nil {
		return rse, fmt.Errorf("Failed to read from world state. %s", err.Error())
	}

	if requestAsBytes == nil {
		return rse, fmt.Errorf("%s does not exist", rid)
	}

	request := Request{}

	_ = json.Unmarshal(requestAsBytes, &request)
	var p = request.Permission
	//request := new(Request)
	//_ = json.Unmarshal(requestAsBytes, request)
	//
	return p, nil
}

/*
//////////////////////////////////query using wallet ///////////////////////////////

	func queryUsrWallet(qID string) (*User, error) {
		os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
		wallet, err := gateway.NewFileSystemWallet("wallet")
		if err != nil {
			return err
		}

		if !wallet.Exists("appUser") {
			err = populateWallet(wallet)
			if err != nil {
				return err
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
			return err
		}
		defer gw.Close()

		network, err := gw.GetNetwork("mychannel")
		if err != nil {
			return err
		}

		contract := network.GetContract("fabcar")

		result, err := contract.EvaluateTransaction("QueryUser", qID)
		if err != nil {
			return err
		}
		return result
	}

	func queryDevWallet(qID string) (*Device, error) {
		os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
		wallet, err := gateway.NewFileSystemWallet("wallet")
		if err != nil {
			return err
		}

		if !wallet.Exists("appUser") {
			err = populateWallet(wallet)
			if err != nil {
				return err
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
			return err
		}
		defer gw.Close()

		network, err := gw.GetNetwork("mychannel")
		if err != nil {
			return err
		}

		contract := network.GetContract("fabcar")

		result, err := contract.SubmitTransaction("QueryDevice", qID)
		if err != nil {
			return err
		}
		return result
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
*/
func main() {

	chaincode, err := contractapi.NewChaincode(new(SmartContract))

	if err != nil {
		fmt.Printf("Error create fabcar chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting fabcar chaincode: %s", err.Error())
	}
}
