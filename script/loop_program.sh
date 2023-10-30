#!/usr/bin/env bash
echo "[" $(date) "] Executing: $0"

cd $(dirname $0)/..

#LOOPCOUNT=50000
LOOPCOUNT=1
#LOOPCOUNT=30
SLEEPTIME=5


if [ x$1 != "x" ]
then
	STRATEGY=$1
else
	
	echo "NO strategy specified.  Required"
	exit
	 
	#STRATEGY="passive"
	
	#STRATEGY="always_hit"
	#STRATEGY="always_stand"
	# best is wizard so far
	STRATEGY1="wizard_simple"
	STRATEGY2="no_bust"
	STRATEGY3="no_bust2"
fi

#BETTINGSTRATEGY1="bet_streaks"
#BETTINGSTRATEGY1="bet_flat"
#BETTINGSTRATEGY1="bet_breakit1"
#BETTINGSTRATEGY1="bet_breakit2"
BETTINGSTRATEGY2="bet_flat"
#BETTINGSTRATEGY2="bet_breakit1"
#BETTINGSTRATEGY2="bet_breakit2"
#BETTINGSTRATEGY2="bet_breakit1"



i=1
MODULUS=2
while [ True ]
do
	START_TIME=$(date +%s)
	echo ""
	echo "==============================================================================="
	echo "[${START_TIME}] PLAYING NEW ROUND # ${i} of LOOPCOUNT: ${LOOPCOUNT} rounds.... "
	echo ""
	echo "i: ${i} loopcount: ${LOOPCOUNT}"
	echo "-------------------------------------------------------------------------------"

	# deprecate the above nonsense.  Use loop_looper_program.sh instead.
	STRATEGY_TO_USE=${STRATEGY}
		
	set -x
	./script/run.sh ${STRATEGY_TO_USE} binary
	set +x
	
	END_TIME=$(date +%s)
	LOOP_TIME=$((END_TIME-START_TIME))
	echo "[${END_TIME}] FINISHED ROUND # ${i} of LOOPCOUNT: ${LOOPCOUNT} in ${LOOP_TIME} secs.... "
	
	METRIC="programming.dev.blackjack_simulator.loop_program.runtime.execution_time ${LOOP_TIME} ${END_TIME}"
	./script/bash_to_graphite.sh ${METRIC}
	
	METRIC="programming.dev.blackjack_simulator.loop_program.runtime.selected_strategy.${STRATEGY_TO_USE} 1 ${END_TIME}"
	./script/bash_to_graphite.sh ${METRIC}

	#
	# Try to send to graphite
	#
	
	echo "Sleeping for ${SLEEPTIME} seconds."
#	sleep ${SLEEPTIME} 

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