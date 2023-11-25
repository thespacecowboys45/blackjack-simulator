#!/usr/bin/env bash
echo "Building..."

cd $(dirname $0)/..

FILES="main.go bets.go bettingstrategies.go cards.go decks.go dlogger_config.go hands.go players.go rounds.go strategies.go"
BINARY="blackjack-simulator"

go build -o ${BINARY} ${FILES}
echo "Build complete"
echo "Run: ${BINARY}"