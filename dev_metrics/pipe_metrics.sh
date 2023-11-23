#!/usr/bin/env bash

HOST="myvps2"
PORT=2003

METRIC_NAME="dev.programming.dev.dev_metrics"
METRIC_BASENAME="value"

echo "pipe_metrics.sh running ...."

# from: https://linuxhint.com/generate-unix-timestamps-linux/
initial_ts=`date -d "2023-10-20 00:00:00" +%s`
ts_increment=30


echo "initial_ts=${initial_ts}"

binary="ncat" # windows
#binary="nc" # linux


N_to_SEND=10
SLEEPTIME=0.1
VALUE_RANGE=100 # 0 to 100 random number

I=0
ts=${initial_ts}
while true
do
	echo "[${I}] Send metric"


	# random value
	value=$((RANDOM %= VALUE_RANGE))
	
	# create the metric to send
	METRIC="${METRIC_NAME}.${METRIC_BASENAME} $value $ts"

	# send it	
	echo ${METRIC} | ${binary} ${HOST} ${PORT}

	# logging
	echo "sent ${METRIC} to ${HOST} on port ${PORT}"

	# add time
	ts=$((ts+ts_increment))
	
	I=$((I+1))
	sleep ${SLEEPTIME}
	
	if [ ${I} -gt 9 ]
	then
		echo "Terminate program loop."
		break
	fi
done

echo "end of line"