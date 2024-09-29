package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/stub"
)

type Asset struct {
	DEALERID  string `json:"dealerId"`
	MSISDN    string `json:"msisdn"`
	MPIN      string `json:"mpin"`
	BALANCE   int    `json:"balance"`
	STATUS    string `json:"status"`
	TRANSAMOUNT int    `json:"transAmount"`
	TRANSTYPE  string `json:"transType"`
	REMARKS    string `json:"remarks"`
}

func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "createAsset" {
		return s.createAsset(stub, args)
	} else if function == "updateAsset" {
		return s.updateAsset(stub, args)
	} else if function == "getAsset" {
		return s.getAsset(stub, args)
	} else if function == "getAssetHistory" {
		return s.getAssetHistory(stub, args)
	}

	return shim.Error("Invalid function name")
}

func (s *SmartContract) createAsset(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 8 {
		return shim.Error("Invalid number of arguments")
	}

	asset := Asset{
		DEALERID:  args[0],
		MSISDN:    args[1],
		MPIN:      args[2],
		BALANCE:   parseInt(args[3]),
		STATUS:    args[4],
		TRANSAMOUNT: parseInt(args[5]),
		TRANSTYPE:  args[6],
		REMARKS:    args[7],
	}

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(args[0], assetJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (s *SmartContract) updateAsset(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 8 {
		return shim.Error("Invalid number of arguments")
	}

	asset, err := s.getAsset(stub, args[:1])
	if err != nil {
		return shim.Error(err.Error())
	}

	asset.DELAERID = args[0]
	asset.MSISDN = args[1]
	asset.MPIN = args[2]
	asset.BALANCE = parseInt(args[3])
	asset.STATUS = args[4]
	asset.TRANSAMOUNT = parseInt(args[5])
	asset.TRANSTYPE = args[6]
	asset.REMARKS = args[7]

	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(args[0], assetJSON)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func (s *SmartContract) getAsset(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Invalid number of arguments")
	}

	assetJSON, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	var asset Asset
	err = json.Unmarshal(assetJSON, &asset)
	if err != nil {
		return shim.Error(err.Error())
	}

	assetJSON, err = json.Marshal(asset)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(assetJSON)
}

func (s *SmartContract) getAssetHistory(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Invalid number of arguments")
	}

	history, err := stub.GetHistoryForKey(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	var historyJSON []byte
	for _, block := range history {
		blockJSON, err := json.Marshal(block)
		if err != nil {
			return shim.Error(err.Error())
		}
		historyJSON = append(historyJSON, blockJSON...)
	}

	return shim.Success(historyJSON)
}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}
	return i
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}
}