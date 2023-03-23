#!/usr/bin/env bash
#
# Works in windows by using/installing 'git bash'
# Works native on linux/macOS
#
# You're welcome.
###
FILES="cards.go decks.go hands.go main.go rounds.go strategies.go bets.go bettingstrategies.go dlogger.go"
STRATEGIES_DIR="strategies"

if [ x$1 != "x" ]
then
	STRATEGY=$1
else 
	STRATEGY="passive"
	STRATEGY="always_hit"
	STRATEGY="always_stand"
fi

BETTINGSTRATEGY="bet_streaks"

RESULTSFILE="results_out.txt"

GAMES=1
VERBOSE="true"

BINARY="blackjack-simulator"

echo "Run strategy: ${STRATEGY}"
set -x
#go run ${FILES} --verbose=${VERBOSE} --games=${GAMES} --resultsfile=${RESULTSFILE} --bettingstrategy=${STRATEGIES_DIR}/${BETTINGSTRATEGY} --strategy="${STRATEGIES_DIR}/${STRATEGY}"
./${BINARY} --verbose=${VERBOSE} --games=${GAMES}  --resultsfile=${RESULTSFILE} --bettingstrategy=${STRATEGIES_DIR}/${BETTINGSTRATEGY} --strategy="${STRATEGIES_DIR}/${STRATEGY}"