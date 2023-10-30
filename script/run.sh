#!/usr/bin/env bash
#
# Works in windows by using/installing 'git bash'
# Works native on linux/macOS
#
# You're welcome.
###

cd $(dirname $0)/..

FILES="bets.go bettingstrategies.go cards.go decks.go  dlogger.go hands.go main.go players.go rounds.go strategies.go"
STRATEGIES_DIR="strategies"

if [ x$1 != "x" ]
then
	STRATEGY=$1
else 
	#STRATEGY="passive"
	#STRATEGY="always_hit"
	#STRATEGY="always_stand"
	# best is wizard so far
	#STRATEGY="wizard_simple"
	#STRATEGY="no_bust"
	#STRATEGY="no_bust2"
	echo "NO strategy specified.  Required"
	exit
fi

BETTINGSTRATEGY1="bet_streaks"
#BETTINGSTRATEGY1="bet_flat"
#BETTINGSTRATEGY1="bet_breakit1"
#BETTINGSTRATEGY1="bet_breakit2"
#BETTINGSTRATEGY2="bet_flat"
#BETTINGSTRATEGY2="bet_breakit1"
#BETTINGSTRATEGY2="bet_breakit2"
#BETTINGSTRATEGY2="bet_breakit1"

BETTINGSTRATEGY2=${BETTINGSTRATEGY1}

RESULTSFILE="results_out.txt"

# Number of games to play per round (run of program)
#GAMES=1
#GAMES=10
GAMES=100
#GAMES=500
#GAMES=1000
#GAMES=5000
#GAMES=10000

NUM_DECKS=6
NUM_PLAYERS=6

VERBOSE="true"

BINARY="blackjack-simulator"

#echo "Building ..." ./script/build.sh;  echo "Build complete."


START_TIME=$(date +%s)

echo "[${START_TIME}][" `date` "] Run strategy: ${STRATEGY}"
set -x


if [ x$2 == "xbinary" ]
then
	echo "run BINARY version"
	#./${BINARY} --verbose=${VERBOSE} --games=${GAMES}  --resultsfile=${RESULTSFILE} --num_decks=${NUM_DECKS} --num_players=${NUM_PLAYERS} --bettingstrategy=${STRATEGIES_DIR}/${BETTINGSTRATEGY1} --bettingstrategy2=${STRATEGIES_DIR}/${BETTINGSTRATEGY2} --strategy="${STRATEGIES_DIR}/${STRATEGY}"
else
	echo "run GOLANG version from source"
	#go run ${FILES} --verbose=${VERBOSE} --games=${GAMES} --resultsfile=${RESULTSFILE} --num_decks=${NUM_DECKS} --num_players=${NUM_PLAYERS} --bettingstrategy=${STRATEGIES_DIR}/${BETTINGSTRATEGY1} --bettingstrategy2=${STRATEGIES_DIR}/${BETTINGSTRATEGY2} --strategy="${STRATEGIES_DIR}/${STRATEGY}"
fi


END_TIME=$(date +%s)
LOOP_TIME=$((END_TIME-START_TIME))

echo "[${END_TIME}] Run of strategy ...${STRATEGY}... complete in ${LOOP_TIME} secs...."

METRIC="programming.dev.blackjack_simulator.run_program.runtime.execution_loop_time ${LOOP_TIME} ${END_TIME}"
./script/bash_to_graphite.sh ${METRIC}

METRIC="programming.dev.blackjack_simulator.run_program.runtime.execution_games_count ${GAMES} ${END_TIME}"
./script/bash_to_graphite.sh ${METRIC}