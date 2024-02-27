#!/bin/bash

# Set environment variables
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
export CHANNEL_NAME=mychannel
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051


START_BLOCK=7
END_BLOCK=20

total_transactions=0
total_time=0

for ((block_num = 7; block_num <= 75; block_num= block_num +68)); do
    # Fetch the block
    peer channel fetch $block_num block$block_num.pb -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com -c $CHANNEL_NAME --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
    
    # Decode the block and extract transaction timestamps
    transaction_count=$(configtxlator proto_decode --input block$block_num.pb --type common.Block | jq '.data.data | length')
    echo "$transaction_count"
    for ((tx_index = 0; tx_index < transaction_count; tx_index++)); do
        transaction_timestamp=$(configtxlator proto_decode --input block$block_num.pb --type common.Block | jq -r ".data.data[$tx_index].payload.header.channel_header.timestamp")
          echo "$transaction_timestamp"

        #if [[ $tx_index -gt 0 ]]; then
          #  prev_transaction_timestamp=$(configtxlator proto_decode --input block$block_num.pb --type common.Block | jq -r ".data.data[$((tx_index - 1))].payload.header.channel_header.timestamp")
          #  echo "$prev_transaction_timestamp"
          #  echo "$transaction_timestamp"
          #  time_diff=$(date -d "$transaction_timestamp" +%s.%N) - $(date -d "$prev_transaction_timestamp" +%s.%N)
           # total_time=$(echo "$total_time + $time_diff" | bc)
     #   fi
        
     #   total_transactions=$((total_transactions + 1))
    done
done

#throughput=$(echo "scale=2; $total_transactions / $total_time" | bc)

#echo "Total transactions: $total_transactions"
#echo "Total time: $total_time seconds"
#echo "Throughput: $throughput transactions per second"
