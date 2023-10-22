#!/usr/bin/env bash
echo "run_loop.sh"

cd $(dirname $0)/..

#LOOPCOUNT=50000
LOOPCOUNT=10000
SLEEPTIME=2

i=1
while [ True ]
do
	START_TIME=$(date +%s)
	echo ""
		echo ""
		echo "[${START_TIME}] PLAYING NEW ROUND # ${i} of LOOPCOUNT: ${LOOPCOUNT} rounds.... "
		echo ""
		echo ""
		echo ""
		echo ""
		echo ""
		echo ""
		echo ""
		echo ""
	#echo "i: ${i} loopcount: ${LOOPCOUNT}"

	#cmd="./run.sh"
	# works
	./script/run.sh
	#eval ${cmd}
	
	END_TIME=$(date +%s)
	LOOP_TIME=$((END_TIME-START_TIME))
	echo "[${END_TIME}] FINISHED ROUND # ${i} of LOOPCOUNT: ${LOOPCOUNT} in ${LOOP_TIME} secs.... "
	
	METRIC="programming.dev.blackjack_simulator.loop_program.runtime.execution_time ${LOOP_TIME} ${END_TIME}"
	./script/bash_to_graphite.sh ${METRIC}
	
	#
	# Try to send to graphite
	#
	
	echo "Sleeping for ${SLEEPTIME} seconds."
	sleep ${SLEEPTIME} 

	i=$((i+1))
	if [ $i -gt ${LOOPCOUNT} ]
	then
		echo ""
		echo ""
		echo "[ END OF LINE ]"
		echo ""
		echo ""
		echo ""
		echo ""
		echo ""
		echo ""
		echo ""
		echo ""
		exit
	fi
done