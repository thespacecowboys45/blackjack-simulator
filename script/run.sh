#!/usr/bin/env bash
#
# Works in windows by using/installing 'git bash'
# Works native on linux/macOS
#
# You're welcome.
###

cd $(dirname $0)/..

FILES="cards.go decks.go hands.go main.go rounds.go strategies.go bets.go bettingstrategies.go dlogger.go"
STRATEGIES_DIR="strategies"

if [ x$1 != "x" ]
then
	STRATEGY=$1
else 
	#STRATEGY="passive"
	#STRATEGY="always_hit"
	#STRATEGY="always_stand"
	STRATEGY="no_bust"
fi

BETTINGSTRATEGY="bet_streaks"

RESULTSFILE="results_out.txt"

# Number of games to play per round (run of program)
#GAMES=1
#GAMES=10
GAMES=100
#GAMES=1000
VERBOSE="true"

BINARY="blackjack-simulator"

echo "Building ..."
#./script/build.sh
echo "Build complete."

echo "Run strategy: ${STRATEGY}"
set -x

START_TIME=$(date +%s)

#go run ${FILES} --verbose=${VERBOSE} --games=${GAMES} --resultsfile=${RESULTSFILE} --bettingstrategy=${STRATEGIES_DIR}/${BETTINGSTRATEGY} --strategy="${STRATEGIES_DIR}/${STRATEGY}"
./${BINARY} --verbose=${VERBOSE} --games=${GAMES}  --resultsfile=${RESULTSFILE} --bettingstrategy=${STRATEGIES_DIR}/${BETTINGSTRATEGY} --strategy="${STRATEGIES_DIR}/${STRATEGY}"

END_TIME=$(date +%s)
LOOP_TIME=$((END_TIME-START_TIME))

METRIC="programming.dev.blackjack_simulator.run_program.runtime.execution_loop_time ${LOOP_TIME} ${END_TIME}"
./script/bash_to_graphite.sh ${METRIC}

METRIC="programming.dev.blackjack_simulator.run_program.runtime.execution_games_count ${GAMES} ${END_TIME}"
./script/bash_to_graphite.sh ${METRIC}