package main

import (
	"log"

	"github.com/rohit25sep/Fake-Product-Detection-Hyperledger/asset-transfer/product-chaincode/chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer chaincode: %v", err)
	}
}
