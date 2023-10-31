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


#STRATEGIES+=("advanced")
#STRATEGIES+=("aggressive")
#STRATEGIES+=("always_hit")
#STRATEGIES+=("always_stand")
#STRATEGIES+=("no_bust")
#STRATEGIES+=("no_bust2")
#STRATEGIES+=("passive")
#STRATEGIES+=("wizard_simple")

STRATEGIES+=("auto_strats/autostrat_100000.txt")
STRATEGIES+=("auto_strats/autostrat_200000.txt")
STRATEGIES+=("auto_strats/autostrat_1000000.txt")
STRATEGIES+=("auto_strats/autostrat_2000000.txt")


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