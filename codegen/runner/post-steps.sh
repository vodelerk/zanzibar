#!/usr/bin/env bash
set -e

if [ -z $1 ]; then
	echo "prefix argument (\$1) is missing"
	exit 1
fi

PREFIX="$1"
DIRNAME="$(dirname $0)"
EASY_JSON_RAW_DIR="$DIRNAME/../../scripts/easy_json"
EASY_JSON_DIR="`cd "${EASY_JSON_RAW_DIR}";pwd`"
EASY_JSON_FILE="$EASY_JSON_DIR/easy_json.go"

echo "Generating JSON Marshal/Unmarshal for rest"
for file in $(find "$PREFIX/gen-code" -name "*_structs.go"); do
    go run "$EASY_JSON_FILE" -- "$file"
done
for file in $(find "$PREFIX/clients" -name "*_structs.go"); do
    go run "$EASY_JSON_FILE" -- "$file"
done
for file in $(find "$PREFIX/endpoints" -name "*_structs.go"); do
    go run "$EASY_JSON_FILE" -- "$file"
done