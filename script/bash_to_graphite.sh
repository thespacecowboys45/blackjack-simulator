#!/usr/bin/env bash
#
# @author dxb The Space Cowboy
# @date Oct '23
#
# DESCRIPTION: All good programs need a description.
#
# Bash abstraction layer for graphite metrics interface
HOST=myvps2
PORT=2003

METRIC_BASENAME=$1
VALUE=$2
TS=$3

METRIC="${METRIC_BASENAME} ${VALUE} ${TS}"

if [ x${METRIC_BASENAME} == 'x' ]
then
	echo "no metric sent"
	break
fi


if [ $OSTYPE == 'darwin20' ]
then
	BINARY=nc #linux / MacOS
else
	BINARY=ncat #windows
fi


echo "Bash to graphite send: ${METRIC} through ${BINARY} to ${HOST} on port ${PORT}"
echo "${METRIC}"| ${BINARY} ${HOST} ${PORT}