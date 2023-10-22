#!/usr/bin/env bash
#
# @date Oct '23
# @author dxb The Space Cowboy
#
# DESCRIPTION:
#
# script to test bash script integration with grahite

#source ./bash_to_graphite.sh

TS=`date +%s`
VALUE=5
METRIC_BASENAME="dev.scratch.metric1"

METRIC="${METRIC_BASENAME} ${VALUE} ${TS}"
echo "Send METRIC: ${METRIC}"

./bash_to_graphite.sh ${METRIC}
