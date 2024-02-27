
export PATH=${PWD}/../../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/../organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/../organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051



# Set the new values for batch size and batch timeout
NEW_BATCH_SIZE=10
NEW_BATCH_TIMEOUT=5s

# Fetch the latest config block for the ordering service
peer channel fetch config config_block.pb -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com -c mychannel --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

# Decode the config block and save it as JSON
configtxlator proto_decode --input config_block.pb --type common.Block --output config_block.json

# Extract the config section of the block
jq .data.data[0].payload.data.config config_block.json > config.json

# Update the batch size and batch timeout
jq '.channel_group.groups.Orderer.values.BatchSize.value.max_message_count = $newBatchSize | .channel_group.groups.Orderer.values.BatchTimeout.value.timeout = $newBatchTimeout' --arg newBatchSize "${NEW_BATCH_SIZE}" --arg newBatchTimeout "${NEW_BATCH_TIMEOUT}" config.json > modified_config.json

# Encode the modified config as a protobuf message
configtxlator proto_encode --input config.json --type common.Config --output config.pb
configtxlator proto_encode --input modified_config.json --type common.Config --output modified_config.pb

# Compute the difference between the original and modified configs as a protobuf message
configtxlator compute_update --channel_id mychannel --original config.pb --updated modified_config.pb --output config_update.pb

# Decode the protobuf message into JSON
configtxlator proto_decode --input config_update.pb --type common.ConfigUpdate --output config_update.json

# Wrap the JSON into an envelope message
echo '{"payload":{"header":{"channel_header":{"channel_id":"mychannel", "type":2}},"data":{"config_update":'$(cat config_update.json)'}}}' | jq . > config_update_in_envelope.json

# Encode the envelope message as a protobuf message
configtxlator proto_encode --input config_update_in_envelope.json --type common.Envelope --output config_update_in_envelope.pb

# Submit the updated configuration to the ordering service
peer channel update -f config_update_in_envelope.pb -c mychannel -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

echo "Batch size and batch timeout updated successfully."

