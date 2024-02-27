export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051

# Fetch the most recent block and convert it to JSON format
peer channel fetch oldest -c mychannel -o json 
peer channel fetch newest -c mychannel -o json

configtxlator proto_decode --type common.Block --input mychannel_oldest.block --output mychannel_oldest.json
configtxlator proto_decode --type common.Block --input mychannel_newest.block --output mychannel_newest.json


TIMESTAMP1=$(jq -r '.data.data[0].payload.header.channel_header.timestamp' mychannel_oldest.json)

TIMESTAMP2=$(jq -r '.data.data[0].payload.header.channel_header.timestamp' mychannel_newest.json)


# Print the timestamp value
echo "The timestamp1 is $TIMESTAMP1"
echo "The timestamp2 is $TIMESTAMP2"

# Convert the timestamps to milliseconds
SECONDS1=$(date -d $TIMESTAMP1 +%s)
NANOSECONDS1=$(echo $TIMESTAMP1 | awk -F'[TZ.]' '{print $3}')
MILLISECONDS1=$(($NANOSECONDS1/1000000))

SECONDS2=$(date -d $TIMESTAMP2 +%s)
NANOSECONDS2=$(echo $TIMESTAMP2 | awk -F'[TZ.]' '{print $3}')
MILLISECONDS2=$(($NANOSECONDS2/1000000))

# Calculate the difference in milliseconds
DIFFERENCE=$((($SECONDS2 - $SECONDS1) * 1000 + $MILLISECONDS2 - $MILLISECONDS1))

# Print the difference
echo $DIFFERENCE
