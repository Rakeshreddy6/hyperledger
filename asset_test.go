package main

import (
	"testing"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/stub"
)

func TestCreateAsset(t *testing.T) {
	stub := shim.NewMockStub("asset", new(SmartContract))
	stub.MockInvoke("createAsset", [][]byte{
		[]byte("dealer1"),
		[]byte("msisdn1"),
		[]byte("mpin1"),
		[]byte("100"),
		[]byte("active"),
		[]byte("10"),
		[]byte("credit"),
		[]byte("remarks1"),
	})
	if stub.MockInvoke("getAsset", [][]byte{[]byte("dealer1")}).Status != shim.OK {
		t.Errorf("Expected status %d but got %d", shim.OK, stub.MockInvoke("getAsset", [][]byte{[]byte("dealer1")}).Status)
	}
}

func TestUpdateAsset(t *testing.T) {
	stub := shim.NewMockStub("asset", new(SmartContract))
	stub.MockInvoke("createAsset", [][]byte{
		[]byte("dealer1"),
		[]byte("msisdn1"),
		[]byte("mpin1"),
		[]byte("100"),
		[]byte("active"),
		[]byte("10"),
		[]byte("credit"),
		[]byte("remarks1"),
	})
	stub.MockInvoke("updateAsset", [][]byte{
		[]byte("dealer1"),
		[]byte("msisdn2"),
		[]byte("mpin2"),
		[]byte("200"),
		[]byte("inactive"),
		[]byte("20"),
		[]byte("debit"),
		[]byte("remarks2"),
	})
	if stub.MockInvoke("getAsset", [][]byte{[]byte("dealer1")}).Status != shim.OK {
		t.Errorf("Expected status %d but got %d", shim.OK, stub.MockInvoke("getAsset", [][]byte{[]byte("dealer1")}).Status)
	}
}

func TestGetAsset(t *testing.T) {
	stub := shim.NewMockStub("asset", new(SmartContract))
	stub.MockInvoke("createAsset", [][]byte{
		[]byte("dealer1"),
		[]byte("msisdn1"),
		[]byte("mpin1"),
		[]byte("100"),
		[]byte("active"),
		[]byte("10"),
		[]byte("credit"),
		[]byte("remarks1"),
	})
	if stub.MockInvoke("getAsset", [][]byte{[]byte("dealer1")}).Status != shim.OK {
		t.Errorf("Expected status %d but got %d", shim.OK, stub.MockInvoke("getAsset", [][]byte{[]byte("dealer1")}).Status)
	}
}

func TestGetAssetHistory(t *testing.T) {
	stub := shim.NewMockStub("asset", new(SmartContract))
	stub.MockInvoke("createAsset", [][]byte{
		[]byte("dealer1"),
		[]byte("msisdn1"),
		[]byte("mpin1"),
		[]byte("100"),
		[]byte("active"),
		[]byte("10"),
		[]byte("credit"),
		[]byte("remarks1"),
	})
	stub.MockInvoke("updateAsset", [][]byte{
		[]byte("dealer1"),
		[]byte("msisdn2"),
		[]byte("mpin2"),
		[]byte("200"),
		[]byte("inactive"),
		[]byte("20"),
		[]byte("debit"),
		[]byte("remarks2"),
	})
	if stub.MockInvoke("getAssetHistory", [][]byte{[]byte("dealer1")}).Status != shim.OK {
		t.Errorf("Expected status %d but got %d", shim.OK, stub.MockInvoke("getAssetHistory", [][]byte{[]byte("dealer1")}).Status)
	}
}