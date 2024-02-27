arg1="$1"

cat "$arg1" | jq '.data.data | length'