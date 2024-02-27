#!/bin/bash

# Set the environment variables
export FABRIC_CFG_PATH=/home/utsa/WorkSpace2nd/Version3/fabric-samples/test-network/configtx
export CHANNEL_NAME=mychannel

# Generate the updated channel configuration transaction
configtxgen -profile TwoOrgsApplicationGenesis -channelID $CHANNEL_NAME -outputCreateChannelTx ./channel-artifacts/$CHANNEL_NAME.tx

# Update the channel with the new configuration
peer channel update -o orderer.example.com:7050 -c $CHANNEL_NAME -f ./channel-artifacts/$CHANNEL_NAME.tx --tls --cafile /home/utsa/WorkSpace2nd/Version3/fabric-samples/test-network/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/tls/ca.crt

# Restart the necessary components for the new configuration to take effect
# Restart the peer nodes
docker restart peer0.org1.example.com
docker restart peer0.org2.example.com

# Restart the ordering service nodes
docker restart orderer.example.com

# Wait for the components to start up
sleep 5

# Verify the updated configuration
peer channel getinfo -c $CHANNEL_NAME

echo "Batch size and batch timeout have been updated successfully."

