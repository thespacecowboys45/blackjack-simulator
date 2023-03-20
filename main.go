package main

import (
	"flag"
	"log"
	"os"
	"fmt"
)

var strategyFile string
var bettingStrategyFile string
var verbose bool
var games int

var totalHands int

func init() {
	flag.StringVar(&strategyFile, "strategy", "", "strategy file path")
	flag.StringVar(&bettingStrategyFile, "bettingstrategy", "", "bettingstrategy file path")
	flag.IntVar(&games, "games", 10, "number of games to play")
	flag.BoolVar(&verbose, "verbose", false, "should output steps")
	
	// DAVB @TODO -
	// starting bank roll
	// default starting wager
	flag.Parse()
}

func pct(top, bottom int) float64 {
	return (float64(top) / float64(bottom)) * 100.0
}

func main() {
	outcomes := make(map[Outcome]int)
	// 'strategy' has two types: softStrategies, and hardStrategies
	strategy := LoadStrategy(strategyFile)
	fmt.Printf("Strategy: %v\n", strategy)
	
	// DAVB - add betting strategy
	bettingStrategy := LoadBettingStrategy(bettingStrategyFile)
	fmt.Printf("BS: %v\n", bettingStrategy)
	
	// DAVB - reset
	bankRoll := NewBankRoll(DEFAULT_BANKROLL)
	fmt.Printf("Starting bankroll: %s\n", bankRoll.String())

	
	for i := 0; i < games; i += 1 {
		deck := NewMultipleDeck(DEFAULT_DECKS)
		
		// DAVB - display the deck before starting
		log.Printf("Deck: %s\n", deck.String())

		// This shuffles all decks together, however many there are
		round := NewRound(deck.Shuffle())

		// DAVB - display the deck before starting
		log.Printf("Deck NOW: %s\n", round.deck.String())
		
		strategy := func(round Round) Action {
			return strategy.GetAction(round.Player, round.Dealer)
		}
		
		// DAVB - @TODO implement some swizzle to incorporate
		// calculating/passing in the bettingstrategy to the computer
		// for now, stub:
		//s := Streak{}
		//fmt.Printf("BettingStrategy: %v\n", bettingStrategy.GetBettingAction(s, 2))
		
		// Make a new wager
		wager := Wager{}
		wager = wager.NewWager(OUTCOME_INIT)
		wagerAction := BETTINGACTION_RESET

		for {

			outcome := round.Play(strategy)
			totalHands += 1
			
			bankRoll = bankRoll.tallyOutcome(outcome, wager)

			// Play 'till we can't play no mo!
			if outcome == OUTCOME_ABORT {
				break
			} else {
				// DAVB - Track how many unique outcomes we have (win/loss/push)
				outcomes[outcome] += 1
			}
			
			// DAVB - @TODO this is where to implement a change in wager logic
			// put this somewhere else, add all functoinality here for now
			s := bankRoll.streak
			fmt.Printf("here streak: %s\n", s.String())
			//wagerAction = bettingStrategy.GetBettingAction(bankRoll.streak, outcome)
			//wagerAction = bettingStrategy.GetBettingAction(s, 2)
			fmt.Printf("WAGER_ACTION: %d\n", wagerAction)
			
			// For now, just keep the same logic
			wager = wager.NewWager(outcome)			
		}
	}

	log.Printf("Total Hands\t\t%d", totalHands)
	log.Printf("Total Wins\t\t%d\t(%0.03f%%)", outcomes[OUTCOME_WIN], pct(outcomes[OUTCOME_WIN], totalHands))
	log.Printf("Total Losses\t%d\t(%0.03f%%)", outcomes[OUTCOME_LOSS], pct(outcomes[OUTCOME_LOSS], totalHands))
	log.Printf("Total Pushes\t%d\t(%0.03f%%)", outcomes[OUTCOME_PUSH], pct(outcomes[OUTCOME_PUSH], totalHands))
	
	log.Printf("Bank Roll\t%v", bankRoll)
	
	os.Exit(0)
	
}
