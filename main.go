package main

import (
	"flag"
	"log"
	"os"
	"fmt"
	"bufio"
	graphite "github.com/jtaczanowski/go-graphite-client"
)


var version string = "1.3"
var strategyFile string
var bettingStrategyFile string
var resultsFile string
var verbose bool
var games int

var totalHands int



func init() {
	flag.StringVar(&strategyFile, "strategy", "", "strategy file path")
	flag.StringVar(&bettingStrategyFile, "bettingstrategy", "", "bettingstrategy file path")
	flag.IntVar(&games, "games", 10, "number of games to play")
	flag.BoolVar(&verbose, "verbose", false, "should output steps")	
	flag.StringVar(&resultsFile, "resultsfile", "", "results file")
	
	// DAVB @TODO -
	// starting bank roll
	// default starting wager
	flag.Parse()
}

func pct(top, bottom int) float64 {
	return (float64(top) / float64(bottom)) * 100.0
}

func sendGraphite(data map[string]float64) error {
	log.Printf("[main.go][testGraphite()][entry]")
	graphiteClient := graphite.NewClient("myvps2", 2003, "programming.dev.blackjack_simulator", "tcp")
	// graphiteClient.SendData(data map[string]float64) error - this method expects a map of metrics as an argument
	if err := graphiteClient.SendData(data); err != nil {
		log.Printf("Error sending metrics: %v", err)
		return err
	}	
	return nil	
}

func testGraphite() error {
	log.Printf("[main.go][testGraphite()][entry]")
	graphiteClient := graphite.NewClient("myvps2", 2003, "programming.dev.dev_metrics.prefix", "tcp")
	 
	 // metrics map
	metricsMap := map[string]float64{
		"test_metric1":  1234.1234,
		"test_metric2": 12345.12345,
	}
	
	// graphiteClient.SendData(data map[string]float64) error - this method expects a map of metrics as an argument
	if err := graphiteClient.SendData(metricsMap); err != nil {
		log.Printf("Error sending metrics: %v", err)
		return err
	}	
	
	log.Printf("[main.go][testGraphite()][exit]")
	return nil
}

func main() {
	log.Printf("Blackjack Simulator version: %s\n", version)
	
/*
dev code - take this out, was put in to test if get working	
	 // metrics map
	metricsMap := map[string]float64{
		"test_metric1":  55,
		"test_metric2": 55.0,
		"test_metric3": 55.5,
		"test_metric4": 555.55,
	}
	
	//testGraphite()
	sendGraphite(metricsMap)
*/	
	
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
		
		// DAVB - idk try it out
		bettingStrategy2 := func(streak Streak) BettingAction {
			return bettingStrategy.GetBettingAction(streak.ConsecutiveLosses,
													streak.ConsecutiveWins)
		}
		
		// DAVB - @TODO implement some swizzle to incorporate
		// calculating/passing in the bettingstrategy to the computer
		// for now, stub:
		//s := Streak{}
		//fmt.Printf("BettingStrategy: %v\n", bettingStrategy.GetBettingAction(s, 2))
		
		// Make a new wager
		wager := Wager{}
		s := Streak{}
		wager = wager.NewWager(OUTCOME_INIT, s, bettingStrategy2)
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
			streak := bankRoll.streak
			fmt.Printf("here streak: %s\n", streak.String())
			//wagerAction = bettingStrategy.GetBettingAction(bankRoll.streak, outcome)
			//wagerAction = bettingStrategy.GetBettingAction(s, 2)
			fmt.Printf("WAGER_ACTION: %d\n", wagerAction)
			
			// For now, just keep the same logic
			wager = wager.NewWager(outcome, streak, bettingStrategy2)			
		}
	}


	log.Printf("Total Hands\t\t%d", totalHands)
	log.Printf("Total Wins\t\t%d\t(%0.03f%%)", outcomes[OUTCOME_WIN], pct(outcomes[OUTCOME_WIN], totalHands))
	log.Printf("Total Losses\t%d\t(%0.03f%%)", outcomes[OUTCOME_LOSS], pct(outcomes[OUTCOME_LOSS], totalHands))
	log.Printf("Total Pushes\t%d\t(%0.03f%%)", outcomes[OUTCOME_PUSH], pct(outcomes[OUTCOME_PUSH], totalHands))
	
	// Send data to remote ---------------------------
	// metrics map
	metricsMap := map[string]float64{
		"outcome.total_hands":  float64(totalHands),
		"outcome.total_wins": float64(outcomes[OUTCOME_WIN]),
		"outcome.total_losses": float64(outcomes[OUTCOME_LOSS]),
		"outcome.total_pushes": float64(outcomes[OUTCOME_PUSH]),
	}
	sendGraphite(metricsMap)


	log.Printf("Bank Roll\t%v", bankRoll)
	metricsMap = map[string]float64{
		"bankroll.amount": float64(bankRoll.Amount),
		"bankroll.min": float64(bankRoll.Min),
		"bankroll.max": float64(bankRoll.Max), 
	}
	sendGraphite(metricsMap)
	
	metricsMap = map[string]float64{
		"bankroll.streak.Wins": float64(bankRoll.streak.Wins),
		"bankroll.streak.Losses": float64(bankRoll.streak.Losses),
		"bankroll.streak.ConsecutiveWins": float64(bankRoll.streak.ConsecutiveWins), 
		"bankroll.streak.ConsecutiveLosses": float64(bankRoll.streak.ConsecutiveLosses),
		"bankroll.streak.MaxConsecutiveWins": float64(bankRoll.streak.MaxConsecutiveWins),
		"bankroll.streak.MaxConsecutiveLosses": float64(bankRoll.streak.MaxConsecutiveLosses),
	}
	sendGraphite(metricsMap)	
	

	

	// Write results file ---------------------------
    // create file
    f, err := os.OpenFile(resultsFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file
    defer f.Close()

    // create new buffer
    buffer := bufio.NewWriter(f)
    output := fmt.Sprintf("tot=%d,win=%d,win_pct=%0.03f%%,loss=%d,loss_pct=%0.03f%%,push=%d,push_pct=%0.03f%%\n%v\n\n", 
		totalHands, 
		outcomes[OUTCOME_WIN],  pct(outcomes[OUTCOME_WIN], totalHands),
		outcomes[OUTCOME_LOSS], pct(outcomes[OUTCOME_LOSS], totalHands),
		outcomes[OUTCOME_PUSH], pct(outcomes[OUTCOME_PUSH], totalHands),bankRoll)
		
	_, err = buffer.WriteString(output)
    // flush buffered data to the file
    if err := buffer.Flush(); err != nil {
        log.Fatal(err)
    }

	// exit normally
	os.Exit(0)

}
