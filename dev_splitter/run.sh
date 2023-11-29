#!/usr/bin/env bash

START_TIME=$(date +%s)

function usage() {
	echo "$0"
}

FILES="dev_splitter.go globals.go"


echo "Run $0"

go run ${FILES} 

echo "--------------"

END_TIME=$(date +%s)

LOOP_TIME=$((END_TIME-START_TIME))

echo "[${END_TIME}] Run of $0 ... complete in ${LOOP_TIME} secs...."