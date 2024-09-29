#!/bin/bash

# Create the network
docker-compose -f network/docker-compose.yml up -d

# Create the channel
docker exec -it orderer peer channel create -c mychannel -f network/configtx.yaml

# Join the peer to the channel
docker exec -it peer0.org1 peer channel join -b mychannel.block

# Install the chaincode
docker exec -it peer0.org1 peer chaincode install -n asset -v 1.0 -p github.com/asset

# Instantiate the chaincode
docker exec -it peer0.org1 peer chaincode instantiate -n asset -v 1.0 -c '{"Args":[]}' -C mychannel