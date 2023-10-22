#!/usr/bin/env bash
echo "run_loop.sh"

cd $(dirname $0)/..

#LOOPCOUNT=50000
LOOPCOUNT=10000
SLEEPTIME=10

i=1
while [ True ]
do
	echo ""
		echo ""
		echo "PLAYING NEW ROUND # ${i}...... of LOOPCOUNT: ${LOOPCOUNT}"
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