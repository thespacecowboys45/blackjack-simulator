#!/usr/bin/env bash

START_TIME=$(date +%s)

function usage() {
	echo "$0"
}

FILES="dev_splitter.go globals.go"
BINARY="dev_slitter"


echo "Run $0"
OPTIONS=""

if [ x$1 == "xbinary" ]
then
	echo "run BINARY version"
	./${BINARY} ${OPTIONS}
else
	echo "run GOLANG version from source"
	go run ${FILES} ${OPTIONS}
fi	


echo "--------------"

END_TIME=$(date +%s)

LOOP_TIME=$((END_TIME-START_TIME))

echo "[${END_TIME}] Run of $0 ... complete in ${LOOP_TIME} secs...."