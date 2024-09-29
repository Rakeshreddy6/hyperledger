package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/identity"
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

func main() {
	// Create a new gateway client
	gateway, err := client.Connect(
		client.WithAddress("localhost:7051"),
		client.WithTLS(true),
		client.WithUser("Admin"),
	)
	if err != nil {
		log.Fatalf("Failed to create gateway client: %v", err)
	}
	defer gateway.Close()

	// Get the network
	network, err := gateway.GetNetwork("mychannel")
	if err != nil {
		log.Fatalf("Failed to get network: %v", err)
	}

	// Get the contract
	contract := network.GetContract("asset")

	http.HandleFunc("/createAsset", func(w http.ResponseWriter, r *http.Request) {
		var asset Asset
		err := json.NewDecoder(r.Body).Decode(&asset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = contract.SubmitTransaction("createAsset", asset.DEALERID, asset.MSISDN, asset.MPIN, strconv.Itoa(asset.BALANCE), asset.STATUS, strconv.Itoa(asset.TRANSAMOUNT), asset.TRANSTYPE, asset.REMARKS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
	})

	http.HandleFunc("/getAsset", func(w http.ResponseWriter, r *http.Request) {
		assetID := r.URL.Query().Get("assetID")
		if assetID == "" {
			http.Error(w, "assetID is required", http.StatusBadRequest)
			return
		}

		result, err := contract.EvaluateTransaction("getAsset", assetID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var asset Asset
		err = json.Unmarshal(result, &asset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(asset)
	})

	http.HandleFunc("/updateAsset", func(w http.ResponseWriter, r *http.Request) {
		var asset Asset
		err := json.NewDecoder(r.Body).Decode(&asset)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		_, err = contract.SubmitTransaction("updateAsset", asset.DEALERID, asset.MSISDN, asset.MPIN, strconv.Itoa(asset.BALANCE), asset.STATUS, strconv.Itoa(asset.TRANSAMOUNT), asset.TRANSTYPE, asset.REMARKS)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	http.HandleFunc("/getAssetHistory", func(w http.ResponseWriter, r *http.Request) {
		assetID := r.URL.Query().Get("assetID")
		if assetID == "" {
			http.Error(w, "assetID is required", http.StatusBadRequest)
			return
		}

		result, err := contract.EvaluateTransaction("getAssetHistory", assetID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
	})

	log.Fatal(http .ListenAndServe(":8080", nil))
}