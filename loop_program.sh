#!/usr/bin/env bash
echo "run_loop.sh"


#LOOPCOUNT=50000
LOOPCOUNT=2

i=1
while [ True ]
do
	echo ""
		echo ""
		echo "PLAYING NEW ROUND ${i}...... MAX: ${LOOPCOUNT}"
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
	./run.sh
	#eval ${cmd} 

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