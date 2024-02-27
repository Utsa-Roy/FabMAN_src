export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export CHANNEL_NAME=mychannel

# Specify the block number and transaction index
BLOCK_NUMBER="newest"
TRANSACTION_INDEX=1

# Fetch the block containing the desired transaction

peer channel fetch $BLOCK_NUMBER $CHANNEL_NAME.pb -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com -c mychannel --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pe

# Extract the timestamp of the desired transaction from the block
TRANSACTION_TIMESTAMP=$(configtxlator proto_decode --input $CHANNEL_NAME.pb --type common.Block | jq -r ".data.data[$TRANSACTION_INDEX].payload.header.channel_header.timestamp")

echo "Timestamp of transaction $TRANSACTION_INDEX in block $BLOCK_NUMBER: $TRANSACTION_TIMESTAMP"F