#!/usr/bin/env bash

FILES="cards.go decks.go hands.go main.go rounds.go strategies.go"
STRATEGIES_DIR="strategies"

if [ x$1 != "x" ]
then
	STRATEGY=$1
else 
	STRATEGY="passive"
fi

echo "Run strategy: ${{TRATEGY}"
set -x
go run ${FILES} --strategy="${STRATEGIES_DIR}/${STRATEGY}"