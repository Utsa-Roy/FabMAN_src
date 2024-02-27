#!/bin/bash

# Block 1
block1="0x1234567890abcdef"
block1_time=$(docker exec -it peer0.org1.example.com peer channel fetch oldest mychannel | grep "$block1" | awk '{print $5}')
block1_timestamp=$(date -d "$block1_time" +"%s%3N")

# Block 2
block2="0xabcdef0123456789"
block2_time=$(docker exec -it peer0.org1.example.com peer channel fetch oldest mychannel | grep "$block2" | awk '{print $5}')
block2_timestamp=$(date -d "$block2_time" +"%s%3N")

# Calculate time gap
time_gap=$((block2_timestamp - block1_timestamp))
echo "Time gap between blocks $block1 and $block2 is $time_gap milliseconds."
