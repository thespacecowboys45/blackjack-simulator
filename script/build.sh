#!/usr/bin/env bash
echo "Building..."

cd $(dirname $0)/..

FILES="cards.go decks.go hands.go main.go rounds.go strategies.go bets.go bettingstrategies.go dlogger.go"
BINARY="blackjack-simulator"

go build -o ${BINARY} ${FILES}
echo "Build complete"
echo "Run: ${BINARY}"