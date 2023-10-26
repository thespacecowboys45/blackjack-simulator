package main

import (
	"flag"
	"log"
	"os"
	"fmt"
	"bufio"
	"math"
	graphite "github.com/jtaczanowski/go-graphite-client"
)


var version string = "1.6"
var strategyFile string
var bettingStrategyFile string
var bettingStrategyFile2 string
var resultsFile string
var verbose bool
var games int
var num_decks int

var totalHands int



func init() {
	flag.StringVar(&strategyFile, "strategy", "", "strategy file path")
	flag.StringVar(&bettingStrategyFile, "bettingstrategy", "", "bettingstrategy file path")
	flag.StringVar(&bettingStrategyFile2, "bettingstrategy2", "", "bettingstrategy2 file path")
	flag.IntVar(&games, "games", 10, "number of games to play")
	flag.IntVar(&num_decks, "num_decks", DEFAULT_DECKS, "number of decks to play from")
	flag.BoolVar(&verbose, "verbose", false, "should output steps")	
	flag.StringVar(&resultsFile, "resultsfile", "", "results file")
	
	// DAVB @TODO -
	// starting bank roll
	// default starting wager
	flag.Parse()
}

// Find the percentage between 
// top - value to estimate percentage
// bottom - value to divide by to estimate percentage
//
// Ex: top=50, bottom=100, percentage == 50.00 == (0.5*100)
//
func pct(top, bottom int) float64 {
	return (float64(top) / float64(bottom)) * 100.0
}

// From: https://gosamples.dev/round-float/
// used to find precision of loss/win percentages
func roundFloat(val float64, precision uint) float64 {
	log.Printf("[main.go][roundFloat()]")
	ratio := math.Pow(10, float64(precision))
	
	result := math.Round(val*ratio) / ratio
	
	 log.Printf("[main.go][roundFloat][val=%f precision=%d ratio = %v result=%f\n", val, precision, ratio, result) 
	return result
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
	log.Printf("[main.go][testGraphite()][entry][version %s]", version)
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
	fmt.Printf("fmt printf[version %s] Blackjack Simulator\n", version)
	log.Printf("[version %s] Blackjack Simulator\n", version)
	dlog(fmt.Sprintf("dlog[version %s] Blackjack Simulator\n", version))
	
	log.Printf("[version %s] Playing %d games per round.\n", version, int(games))
	
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
	fmt.Printf("Load playing strategy from file: %s\n", strategyFile)
	strategy := LoadStrategy(strategyFile)
	fmt.Printf("Strategy: %v\n", strategy)
	
	// DAVB - add betting strategy
	fmt.Printf("Load betting strategy 1 from file: %s\n", bettingStrategyFile)
	bettingStrategy1 := LoadBettingStrategy(bettingStrategyFile)
	fmt.Printf("Betting Strategy1: %v\n", bettingStrategy1)

	// setup the callback function for the bettingStrategy
	/*
	bettingStrategy1fn := func(streak Streak) BettingAction {
		return bettingStrategy1.GetBettingAction(streak.ConsecutiveLosses,
												streak.ConsecutiveWins)
	}
	*/

	// DAVB - add betting strategy2
	fmt.Printf("Load betting strategy 2 from file: %s\n", bettingStrategyFile2)
	bettingStrategy2 := LoadBettingStrategy(bettingStrategyFile2)
	fmt.Printf("Betting Strategy2: %v\n", bettingStrategy2)	

	// setup the callback function for the bettingStrategy
	/*	

	bettingStrategy2fn := func(streak Streak) BettingAction {
		return bettingStrategy2.GetBettingAction(streak.ConsecutiveLosses,
												streak.ConsecutiveWins)
	}
	*/

	
	// DAVB - reset
	bankRoll := NewBankRoll(DEFAULT_BANKROLL)
	fmt.Printf("Starting bankroll: %s\n", bankRoll.String())
	
	
	for i := 0; i < games; i += 1 {
		//deck := NewMultipleDeck(DEFAULT_DECKS)
		deck := NewMultipleDeck(num_decks)
		
		// DAVB - display the deck before starting
		log.Printf("[main.go][game #%d] Deck before shuffle: %s\n", i, deck.String())

		// This shuffles all decks together, however many there are
		round := NewRound(deck.Shuffle())

		// DAVB - display the deck before starting
		log.Printf("[main.go][game #%d] Deck after shuffle: %s\n", i, round.deck.String())
		
		
		//
		// dxb - seriously curious (why) are these here and not outside the loop
		//
		strategy := func(round Round) Action {
			return strategy.GetAction(round.Player, round.Dealer)
		}

/*
move outside of loop		
		// why is this inside the loop - i do not understand that
*/			
		bettingStrategy1fn := func(streak Streak) BettingAction {
			return bettingStrategy1.GetBettingAction(streak.ConsecutiveLosses,
													streak.ConsecutiveWins)
		}
		bettingStrategy2fn := func(streak Streak) BettingAction {
			return bettingStrategy2.GetBettingAction(streak.ConsecutiveLosses,
													streak.ConsecutiveWins)
		}
		
	
		
		// DAVB - @TODO implement some swizzle to incorporate
		// calculating/passing in the bettingstrategy to the computer
		// for now, stub:
		//s := Streak{}
		//fmt.Printf("BettingStrategy: %v\n", bettingStrategy.GetBettingAction(s, 2))
		
		
		// @TODO - splits1
		// @TODO - for splits this needs to be collapsed into the Hand itself.
		// Like, the 'hand' structure needs to become more complex, encompasing the
		// wager, as well as the cards the player owns for the hand.
		//
		
		// Make a new wager
		wager := Wager{}
		
		
		
		// dxb - why do I have this here ? Is this to initialize the function with something?
		// why not use bankRoll.streak instead ? 
		s := Streak{}
		
		
		
		// This could be any betting strategy - it's just INITIALIZATION
		if (i < (games / 2)) {
//		if (true) {
			fmt.Printf("[main.go][i=%d][i mod 2=%d - use strategy ONE - on INIT\n", i, (i %2))
			// temporarily fenagle, unused variable betteingStrategy2fn
			wager = wager.NewWager(OUTCOME_INIT, s, bettingStrategy2fn)
		} else {
			fmt.Printf("[main.go][i=%d][i mod 2=%d - use strategy TWO - on INIT\n", i, (i %2))
			wager = wager.NewWager(OUTCOME_INIT, s, bettingStrategy2fn)	
		}
		
		// Is this needed ?- or remove completely.  Decision is not made here.
		//wagerAction := BETTINGACTION_RESET

		for {

			outcome := round.Play(strategy)
			totalHands += 1
			
			bankRoll = bankRoll.tallyOutcome(outcome, wager)

			// Play 'till we can't play no mo!
			// Basically: the shoe has run out of cards.
			if outcome == OUTCOME_ABORT {
				break
			} else {
				// DAVB - Track how many unique outcomes we have (win/loss/push)
				outcomes[outcome] += 1
			}
			
			// DAVB - @TODO this is where to implement a change in wager logic
			// put this somewhere else, add all functoinality here for now
			streak := bankRoll.streak
			fmt.Printf("[main.go] here streak: %s\n", streak.String())

			
			/**
			 * Oct '23
			 * @TODO - okay, this looks like a work in-progress.
			 * The concept looks like to build in an action based on the results of the last bet
			 * and the betting streak going on
			 */
			 
			// On second glance - this looks to be development code which did not work.
			// Instead - the bettingStrategy is set above.  We pass the function
			// 'bettingStrategy1fn' to the NewWager function in order that it may determine
			// a new wager inside itself.
			
			// yeah, that.
			
			
			// For now, just keep the same logic
			// Toggle between two betting strategies every game
			
			// Flip-flop strategies half-way through games
			if (i < (games / 2)) {
			//if (true) {
				fmt.Printf("[main.go][i=%d][i mod 2=%d - use strategy ONE - on PLAY\n", i, (i %2))
				wager = wager.NewWager(outcome, streak, bettingStrategy1fn)
			} else {
				fmt.Printf("[main.go][i=%d][i mod 2=%d - use strategy TWO - on PLAY\n", i, (i %2))
				wager = wager.NewWager(outcome, streak, bettingStrategy2fn)	
			}
		}
	}


	log.Printf("Total Hands\t\t%d", totalHands)
	log.Printf("Total Wins\t\t%d\t(%0.03f%%)", outcomes[OUTCOME_WIN], pct(outcomes[OUTCOME_WIN], totalHands))
	log.Printf("Total Losses\t%d\t(%0.03f%%)", outcomes[OUTCOME_LOSS], pct(outcomes[OUTCOME_LOSS], totalHands))
	log.Printf("Total Pushes\t%d\t(%0.03f%%)", outcomes[OUTCOME_PUSH], pct(outcomes[OUTCOME_PUSH], totalHands))
	
	winPct := roundFloat(pct(outcomes[OUTCOME_WIN], totalHands), 2)
	log.Printf("Total Wins percentage two: %0.03f%%", winPct)
	
	// Send data to remote ---------------------------
	// metrics map
	metricsMap := map[string]float64{
		fmt.Sprintf("%ddecks.%dgames.outcome.total_hands", num_decks, games) : float64(totalHands),
		fmt.Sprintf("%ddecks.%dgames.outcome.total_wins", num_decks, games) : float64(outcomes[OUTCOME_WIN]),
		fmt.Sprintf("%ddecks.%dgames.outcome.total_losses", num_decks, games) : float64(outcomes[OUTCOME_LOSS]),
		fmt.Sprintf("%ddecks.%dgames.outcome.total_pushes", num_decks, games) : float64(outcomes[OUTCOME_PUSH]),
		fmt.Sprintf("%ddecks.%dgames.outcome.percent_wins", num_decks, games) : roundFloat(pct(outcomes[OUTCOME_WIN], totalHands), 2),
		fmt.Sprintf("%ddecks.%dgames.outcome.percent_losses", num_decks, games) : roundFloat(pct(outcomes[OUTCOME_LOSS], totalHands), 2),
		fmt.Sprintf("%ddecks.%dgames.outcome.percent_pushes", num_decks, games) : roundFloat(pct(outcomes[OUTCOME_PUSH], totalHands), 2),
		
	}
	log.Printf("Send MetricsMap 1: %v\n", metricsMap)
	sendGraphite(metricsMap)


	log.Printf("Bank Roll\t%v", bankRoll)
	metricsMap = map[string]float64{
		fmt.Sprintf("%ddecks.%dgames.bankroll.amount", num_decks, games) : float64(bankRoll.Amount),
		fmt.Sprintf("%ddecks.%dgames.bankroll.min", num_decks, games) : float64(bankRoll.Min),
		fmt.Sprintf("%ddecks.%dgames.bankroll.max", num_decks, games) : float64(bankRoll.Max), 
	}
	log.Printf("Send MetricsMap 2: %v\n", metricsMap)	
	sendGraphite(metricsMap)
	
	metricsMap = map[string]float64{
		fmt.Sprintf("%ddecks.%dgames.bankroll.streak.Wins", num_decks, games) : float64(bankRoll.streak.Wins),
		fmt.Sprintf("%ddecks.%dgames.bankroll.streak.Losses", num_decks, games) : float64(bankRoll.streak.Losses),
		fmt.Sprintf("%ddecks.%dgames.bankroll.streak.ConsecutiveWins", num_decks, games) : float64(bankRoll.streak.ConsecutiveWins), 
		fmt.Sprintf("%ddecks.%dgames.bankroll.streak.ConsecutiveLosses", num_decks, games) : float64(bankRoll.streak.ConsecutiveLosses),
		fmt.Sprintf("%ddecks.%dgames.bankroll.streak.MaxConsecutiveWins", num_decks, games) : float64(bankRoll.streak.MaxConsecutiveWins),
		fmt.Sprintf("%ddecks.%dgames.bankroll.streak.MaxConsecutiveLosses", num_decks, games) : float64(bankRoll.streak.MaxConsecutiveLosses),
		fmt.Sprintf("%ddecks.%dgames.bankroll.streak.MaxWagerWon", num_decks, games) : float64(bankRoll.streak.MaxWagerWon),
		fmt.Sprintf("%ddecks.%dgames.bankroll.streak.MaxWagerLost", num_decks, games) : float64(bankRoll.streak.MaxWagerLost),
	}
	log.Printf("Send MetricsMap 3: %v\n", metricsMap)	
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
    output := fmt.Sprintf("games=%d,tot=%d,win=%d,win_pct=%0.03f%%,loss=%d,loss_pct=%0.03f%%,push=%d,push_pct=%0.03f%%\n%v\n\n", 
		games, 
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
