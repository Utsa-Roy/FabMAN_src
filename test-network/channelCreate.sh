
arg1="$1"

./network.sh createChannel -c "$arg1"
./network.sh deployCC -c "$arg1" -ccn fabcar -ccp ../chaincode/fabcar/go -ccl go

