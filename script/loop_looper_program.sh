#!/usr/bin/env bash
#
# @date October 23
# @author dxb The Space Cowboy
#
####
#
# DESCRIPTION:
# mit-license implied
#
# Feeds the main loop, cycling through strategies and betting strategies.
#
# Creates, in essence, enough data by running enouch cycles through the
# looper.#
# For us to evaluate over time and see differences.
#
# If any.
#
######
LOOPER_BIN="./script/loop_program.sh"

cd $(dirname $0)/..


echo "[" $(date) "] Executing: $0"
#print " Executing: "

SLEEPTIME=6

# all available strategies to test
STRATEGIES=()


STRATEGIES+=("almost_always_hit")
STRATEGIES+=("always_hit")
STRATEGIES+=("advanced")
STRATEGIES+=("aggressive")
STRATEGIES+=("passive")
STRATEGIES+=("wizard_simple")
STRATEGIES+=("always_stand")
STRATEGIES+=("no_bust")
STRATEGIES+=("no_bust2")


#STRATEGIES+=("autostrat_100000")
#STRATEGIES+=("autostrat_200000")

######### RUN 1
#
#STRATEGIES+=("autostrat_1000000")
#STRATEGIES+=("autostrat_2000000")
#STRATEGIES+=("autostrat_3000000")
#STRATEGIES+=("autostrat_4000000")
#STRATEGIES+=("autostrat_5000000")
#STRATEGIES+=("autostrat_6000000")
#STRATEGIES+=("autostrat_7000000")
#STRATEGIES+=("autostrat_8000000")
#STRATEGIES+=("autostrat_9000000")




#STRATEGIES+=("autostrat_100000000")
#STRATEGIES+=("autostrat_200000000")
#STRATEGIES+=("autostrat_300000000")
#STRATEGIES+=("autostrat_400000000")
#STRATEGIES+=("autostrat_500000000")
#STRATEGIES+=("autostrat_600000000")
#STRATEGIES+=("autostrat_700000000")
#STRATEGIES+=("autostrat_800000000")
#STRATEGIES+=("autostrat_900000000")






# @TODO - loop through betting strategies
#BETTING_STRATEGIES=("bet_streaks" "bet_flat" "bet_break1" "bet_break2")
#BETTING_STRATEGIES=("bet_streaks" "bet_flat")
BETTING_STRATEGIES=("bet_flat")
for strategy in ${STRATEGIES[@]}
do
	echo "[" `date` "]Run strategy: ${strategy}"
	for betting_strategy in ${BETTING_STRATEGIES[@]}
	do
		echo "[" `date` "]Run strategy: ${strategy} agasint betting_strategy: ${betting_strategy}"
		./${LOOPER_BIN} ${strategy} ${betting_strategy}	
		echo "[" `date` "]Finished running strategy: ${strategy} agasint betting_strategy: ${betting_strategy}"
		
	done

	echo "[" `date` "]Finished running ${strategy}"
	sleep ${SLEEPTIME}
done	
	
	
print " END OF LINE: "