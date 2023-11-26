package main
/***
```
Copyright (c) 2023 dxb The Space Cowboy

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE
```
***/
import (
	"flag"
	"log"
	"os"
	"fmt"
	"bufio"
	"math"
	graphite "github.com/jtaczanowski/go-graphite-client"
	dlog "bitbucket.org/thespacecowboys45/dlogger"
)

var version string = "1.6.5"
var strategyFile string
var bettingStrategyFile string
var bettingStrategyFile2 string
var resultsFile string
var verbose bool
var games int
var num_decks int
var num_players int
var totalHands int
var totalHandsRecalc int

func init() {
	flag.StringVar(&strategyFile, "strategy", "", "strategy file path")
	flag.StringVar(&bettingStrategyFile, "bettingstrategy", "", "bettingstrategy file path")
	flag.StringVar(&bettingStrategyFile2, "bettingstrategy2", "", "bettingstrategy2 file path")
	flag.IntVar(&games, "games", 10, "number of games to play")
	flag.IntVar(&num_decks, "num_decks", DEFAULT_DECKS, "number of decks to play from")
	flag.IntVar(&num_players, "num_players", MIN_PLAYERS, "number of players sitting at the table")
	flag.BoolVar(&verbose, "verbose", false, "should output steps")	
	flag.StringVar(&resultsFile, "resultsfile", "", "results file")
	
	// DAVB @TODO - add variables
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
	//log.Printf("[main.go][roundFloat()]")
	ratio := math.Pow(10, float64(precision))
	
	result := math.Round(val*ratio) / ratio
	
	//log.Printf("[main.go][roundFloat][val=%f precision=%d ratio = %v result=%f\n", val, precision, ratio, result) 
	return result
}

func sendGraphite(data map[string]float64) error {
	log.Printf("[main.go][sendGraphite()][entry]")
	graphiteClient := graphite.NewClient("myvps2", 2003, "programming.dev.blackjack_simulator", "tcp")
	// graphiteClient.SendData(data map[string]float64) error - this method expects a map of metrics as an argument
	if err := graphiteClient.SendData(data); err != nil {
		log.Printf("Error sending metrics: %v", err)
		return err
	}
	log.Printf("[main.go][sendGraphite()][exit]")	
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

func printRuntimeVars() {
	log.Printf("[main.go][printRuntimeVars()][entry]")
	log.Printf("[main.go][printRuntimeVars()][version=%s]", version)
	log.Printf("[main.go][printRuntimeVars()][strategyFile=%s]", strategyFile)
	log.Printf("[main.go][printRuntimeVars()][bettingStrategyFile=%s]", bettingStrategyFile)
	log.Printf("[main.go][printRuntimeVars()][bettingStrategyFile2=%s]", bettingStrategyFile2)
	log.Printf("[main.go][printRuntimeVars()][games=%d]", games)
	log.Printf("[main.go][printRuntimeVars()][num_decks=%d]", num_decks)
	log.Printf("[main.go][printRuntimeVars()][num_players=%d]", num_players)
	log.Printf("[main.go][printRuntimeVars()][verbose=%t]", verbose)
	log.Printf("[main.go][printRuntimeVars()][resultsFile=%s]", resultsFile)
	log.Printf("[main.go][printRuntimeVars()][exit]")
}

func main() {
	// figuring out proper way to log output
	configureDlogger()

	fmt.Printf("fmt printf[version %s] Blackjack Simulator\n", version)
	log.Printf("[version %s] Blackjack Simulator\n", version)
	
	// Test dlogger is working for the levels we want to 
	dlog.Debug(fmt.Sprintf("dlog[version %s] Blackjack Simulator\n", version))
	dlog.Info(fmt.Sprintf("dlog[version %s] Blackjack Simulator\n", version))
	dlog.Error(fmt.Sprintf("dlog[version %s] Blackjack Simulator\n", version))
	
	printRuntimeVars()
	
	// @TODO - incorporate
	dlog.Debug("dlogger - HERE - dlog")
	
	log.Printf("[version %s] Playing %d games per round.\n", version, int(games))
	
/*
dev code - take this out, was put in to test if graphite metrics are working	
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
	
	// Track Total # of rounds
	roundsPlayed := 0
	
	// Used to track stats for unique outcomes
	outcomes := make(map[Outcome]int)
	
	// @TODO - add outcome stats tracking per player
	playerOutcomes := make([]map[Outcome]int, num_players)
	for i := range playerOutcomes {
		playerOutcomes[i] = make(map[Outcome]int)
	}
	var playerTotalHands [6]int
	
	
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
	playerWagers := make([]Wager, num_players)
	playerBankRolls := make([]BankRoll, num_players)
	for i:=0; i<num_players; i++ {
		//bankRoll := NewBankRoll(DEFAULT_BANKROLL)
		playerBankRolls[i] = NewBankRoll(DEFAULT_BANKROLL)
		fmt.Printf("[main.go][main][player #%d] Starting bankroll: %s\n", i, playerBankRolls[i].String())
	}
	
	
	for i := 0; i < games; i += 1 {
		//deck := NewMultipleDeck(DEFAULT_DECKS)
		deck := NewMultipleDeck(num_decks)
		
		// DAVB - display the deck before starting
		log.Printf("[main.go][game #%d] Deck before shuffle: %s\n", i, deck.String())

		// This shuffles all decks together, however many there are
		round := NewRound(deck.Shuffle(), num_players)

		// DAVB - display the deck before starting
		log.Printf("[main.go][game #%d] Deck after shuffle: %s\n", i, round.deck.String())
		
		//
		// dxb - seriously curious (why) are these callback function 
		// declarations here and not outside the for loop
		//
		
		// Worked for single-player.  refactor for multi-player.
		
//		strategy := func(round Round) Action {
//			return strategy.GetAction(round.Player, round.Dealer)
//		}
	
		
		// dxb - Nov '23
		// deprecate, this is using the "single hand" code only
		// multi-player code
//		strategy := func(round Round, player_number int) Action {
//			return strategy.GetAction(round.Players[player_number], round.Dealer)
//		}

		// dxb - Nov '23
		// multi-hand code		
		strategy := func(round Round, player_number int, hand_number int) Action {
			//return strategy.GetAction(round.Players[player_number], round.Dealer)
			return strategy.GetAction(round.PlayersObj[player_number].Hands[hand_number], round.Dealer)
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
		// @TODO - make a wager per player
		// deprecated
		//wager := Wager{}
		
		
		
		// dxb - why do I have this here ? Is this to initialize the function with something?
		// why not use bankRoll.streak instead ?
		// deprecated 
		//s := Streak{}
		
		
		
		// This could be any betting strategy function - it's just INITIALIZATION
		for j:=0; j<num_players; j++ {
			fmt.Printf("[main.go][player #%d]INIT wager [i=%d][i mod 2=%d - use strategy ONE - on INIT\n", j, i, (i %2))
			playerWagers[j] = playerWagers[j].NewWager(OUTCOME_INIT, Streak{}, bettingStrategy1fn)	
		}
		
		// Is this needed ?- or remove completely.  Decision is not made here.
		//wagerAction := BETTINGACTION_RESET

		for {

			//outcome := round.Play(strategy)
			//outcome := round.PlayMultiPlayer(strategy)
			roundOutcomes, total_hands_played_this_round := round.PlayMultiPlayer(strategy)
			
			dlog.Info("AFTER ROUND, Round looks like:")
			round.toString()
			
			//totalHands += 1
			// @TODO - splits1 - adapt for multiple hands per player
			//totalHands += num_players
			totalHands += num_players
			totalHandsRecalc += total_hands_played_this_round


			dlog.Debug("[main.go][round over][total_hands_played_this_round==%d, total_hands==%d, totalHandsRecalc==%d]", 
				total_hands_played_this_round,
				totalHands,
				totalHandsRecalc)


			
			endOfRound := false
			for j:=0; j<num_players; j++ {
				log.Printf("[main.go][player #%d][Examining all player outcomes][outcome=%s]", j, outcomeToString(roundOutcomes[j]))
				if roundOutcomes[j] == OUTCOME_ABORT {
				
					// code-sanity check!
					if len(round.deck) > MINIMUM_SHOE_SIZE {
						dlog.Error("[main.go][sanity check failed!][There are still enough cards in the deck: %d", len(round.deck))
						dlog.Error("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
					}
					
					log.Printf("[main.go][player #%d][Round is over - we are out of cards.]", j)
					endOfRound = true
					// short-circuit.  If we are hitting this code then: all players will have the same outcome==OUTCOME_ABORT
					break
				} else {
					// Track overall game stats per outcome possibility
					outcomes[roundOutcomes[j]] += 1
					
					// per player stats (@TODO - deprecate single-hand code
					// NOTE: roundOutcomes[j] is of type Outcome (e.g. INIT, WIN, LOSS, etc.)
					playerOutcomes[j][roundOutcomes[j]] += 1
					
					// multi-hand code
					for k:=0; k<len(round.PlayersObj[j].Outcomes); k++ {
						playerOutcomes[j][round.PlayersObj[j].Outcomes[k]] += 1	
					}
					
					// deprecated, single-hand code
					//playerTotalHands[j] += 1
					
					// refactored for multi-hand code
					playerTotalHands[j] += len(round.PlayersObj[j].Hands)
					
										// re-implement per-player @TODO
					//bankRoll = bankRoll.tallyOutcome(outcome, wager)
					playerBankRolls[j] = playerBankRolls[j].tallyOutcome(roundOutcomes[j], playerWagers[j])
		
					
					// DAVB - @TODO this is where to implement a change in wager logic
					// put this somewhere else, add all functoinality here for now
					
					//streak := bankRoll.streak
					streak := playerBankRolls[j].streak
					
					//fmt.Printf("[main.go] here streak: %s\n", streak.String())
		
					
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
						fmt.Printf("[main.go][player #%d][game_number=%d][i mod 2=%d - use strategy ONE - on PLAY\n", j, i, (i %2))
						playerWagers[j] = playerWagers[j].NewWager(roundOutcomes[j], streak, bettingStrategy1fn)
					} else {
						// Current logic: for games==1 this code will hit
						fmt.Printf("[main.go][player #%d][game_number=%d][i mod 2=%d - use strategy TWO - on PLAY\n", j, i, (i %2))
						playerWagers[j] = playerWagers[j].NewWager(roundOutcomes[j], streak, bettingStrategy2fn)	
					}
				}
			}
						
			if endOfRound {
				roundsPlayed += 1
				break
			}
		}
	}


	log.Printf("Total Rounds\t%d", roundsPlayed)
	log.Printf("Total Hands\t\t%d", totalHands)
	log.Printf("Total Round Wins\t\t%d\t(%0.03f%%)", outcomes[OUTCOME_WIN], pct(outcomes[OUTCOME_WIN], totalHands))
	log.Printf("Total Round Losses\t%d\t(%0.03f%%)", outcomes[OUTCOME_LOSS], pct(outcomes[OUTCOME_LOSS], totalHands))
	log.Printf("Total Round Pushes\t%d\t(%0.03f%%)", outcomes[OUTCOME_PUSH], pct(outcomes[OUTCOME_PUSH], totalHands))


	// Sanity checking
	log.Printf("---- recalculating win/loss percentage --- sanity checking ----")
	validateTotalHands := outcomes[OUTCOME_WIN] + outcomes[OUTCOME_LOSS] + outcomes[OUTCOME_PUSH]
	if totalHandsRecalc != validateTotalHands {
		dlog.Debug("[main.go][sanity check!!!!!!!!!!!][ totalHandsRecalc (%d) != validateTotalHands (%d} ]",
			totalHandsRecalc,
			validateTotalHands)
			
		// @TODO - handle this and check program code
		log.Fatal("[main.go][program logic is not sane.]")
	}

	log.Printf("Total Rounds\t%d", roundsPlayed)
	log.Printf("Total Hands recalculated\t\t%d", totalHandsRecalc)
	log.Printf("Total Round Wins recalculated\t\t%d\t(%0.03f%%)", outcomes[OUTCOME_WIN], pct(outcomes[OUTCOME_WIN], totalHandsRecalc))
	log.Printf("Total Round Losses recalculated\t%d\t(%0.03f%%)", outcomes[OUTCOME_LOSS], pct(outcomes[OUTCOME_LOSS], totalHandsRecalc))
	log.Printf("Total Round Pushes recalculated\t%d\t(%0.03f%%)", outcomes[OUTCOME_PUSH], pct(outcomes[OUTCOME_PUSH], totalHandsRecalc))


	///// end sanity checking



	// Send data to remote ---------------------------
	// metrics map
	metricsMap := map[string]float64{
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.rounds_played", num_players, num_decks, games) : float64(roundsPlayed),
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.total_hands", num_players, num_decks, games) : float64(totalHands),
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.total_wins", num_players, num_decks, games) : float64(outcomes[OUTCOME_WIN]),
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.total_losses", num_players, num_decks, games) : float64(outcomes[OUTCOME_LOSS]),
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.total_pushes", num_players, num_decks, games) : float64(outcomes[OUTCOME_PUSH]),
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.percent_wins", num_players, num_decks, games) : roundFloat(pct(outcomes[OUTCOME_WIN], totalHands), 2),
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.percent_losses", num_players, num_decks, games) : roundFloat(pct(outcomes[OUTCOME_LOSS], totalHands), 2),
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.percent_pushes", num_players, num_decks, games) : roundFloat(pct(outcomes[OUTCOME_PUSH], totalHands), 2),
		
	}
	log.Printf("Send MetricsMap 1: %v\n", metricsMap)
	sendGraphite(metricsMap)
	
	// Send data to remote ---------------------------
	// metrics map
	metricsMap = map[string]float64{
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.total_hands_recalc", num_players, num_decks, games) : float64(totalHandsRecalc),
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.percent_wins_recalc", num_players, num_decks, games) : roundFloat(pct(outcomes[OUTCOME_WIN], totalHandsRecalc), 2),
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.percent_losses_recalc", num_players, num_decks, games) : roundFloat(pct(outcomes[OUTCOME_LOSS], totalHandsRecalc), 2),
		fmt.Sprintf("%dplayers.%ddecks.%dgames.round.outcome.percent_pushes_recalc", num_players, num_decks, games) : roundFloat(pct(outcomes[OUTCOME_PUSH], totalHandsRecalc), 2),
		
	}
	log.Printf("Send MetricsMap 2 recalculated: %v\n", metricsMap)
	sendGraphite(metricsMap)	

	
	// output player stats
	for j:=0; j<num_players; j++ {
		log.Printf("Total player %d Hands\t\t%d", j, playerTotalHands[j])
		log.Printf("Total player %d Wins\t\t%d\t(%0.03f%%)", j, playerOutcomes[j][OUTCOME_WIN], pct(playerOutcomes[j][OUTCOME_WIN], playerTotalHands[j]))
		log.Printf("Total player %d Losses\t%d\t(%0.03f%%)", j, playerOutcomes[j][OUTCOME_LOSS], pct(playerOutcomes[j][OUTCOME_LOSS], playerTotalHands[j]))
		log.Printf("Total player %d Pushes\t%d\t(%0.03f%%)", j, playerOutcomes[j][OUTCOME_PUSH], pct(playerOutcomes[j][OUTCOME_PUSH], playerTotalHands[j]))
		
		metricsMap := map[string]float64{
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.outcome.total_hands", num_players, num_decks, games, j+1) : float64(playerTotalHands[j]),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.outcome.total_wins", num_players, num_decks, games, j+1) : float64(playerOutcomes[j][OUTCOME_WIN]),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.outcome.total_losses", num_players, num_decks, games, j+1) : float64(playerOutcomes[j][OUTCOME_LOSS]),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.outcome.total_pushes", num_players, num_decks, games, j+1) : float64(playerOutcomes[j][OUTCOME_PUSH]),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.outcome.percent_wins", num_players, num_decks, games, j+1) : roundFloat(pct(playerOutcomes[j][OUTCOME_WIN], playerTotalHands[j]), 2),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.outcome.percent_losses", num_players, num_decks, games, j+1) : roundFloat(pct(playerOutcomes[j][OUTCOME_LOSS], playerTotalHands[j]), 2),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.outcome.percent_pushes", num_players, num_decks, games, j+1) : roundFloat(pct(playerOutcomes[j][OUTCOME_PUSH], playerTotalHands[j]), 2),
			
		}
		log.Printf("Send MetricsMap for player %d : %v\n", j, metricsMap)
		sendGraphite(metricsMap)

	}
	

	// output bankroll stats
	for j:=0; j<num_players; j++ {
		log.Printf("[player #%d] Bank Roll\t%v", j, playerBankRolls[j])
		metricsMap = map[string]float64{
		
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.starting_amount", num_players, num_decks, games, j+1) : float64(DEFAULT_BANKROLL),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.amount", num_players, num_decks, games, j+1) : float64(playerBankRolls[j].Amount),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.min", num_players, num_decks, games,j+1) : float64(playerBankRolls[j].Min),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.max", num_players, num_decks, games,j+1) : float64(playerBankRolls[j].Max), 
		}
		log.Printf("Send MetricsMap bankroll for player %d:\n%v\n", j, metricsMap)	
		sendGraphite(metricsMap)
		
		/**
		 * @TODO - idea: implement sharpe ratio tracking (is this possible?  average over ? games ?)
		 */
		
	
		// output streak stats
		metricsMap = map[string]float64{
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.streak.Wins", num_players, num_decks, games, j+1) : float64(playerBankRolls[j].streak.Wins),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.streak.Losses", num_players, num_decks, games, j+1) : float64(playerBankRolls[j].streak.Losses),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.streak.ConsecutiveWins", num_players, num_decks, games, j+1) : float64(playerBankRolls[j].streak.ConsecutiveWins), 
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.streak.ConsecutiveLosses", num_players, num_decks, games, j+1) : float64(playerBankRolls[j].streak.ConsecutiveLosses),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.streak.MaxConsecutiveWins", num_players, num_decks, games, j+1) : float64(playerBankRolls[j].streak.MaxConsecutiveWins),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.streak.MaxConsecutiveLosses", num_players, num_decks, games, j+1) : float64(playerBankRolls[j].streak.MaxConsecutiveLosses),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.streak.MaxWagerWon", num_players, num_decks, games, j+1) : float64(playerBankRolls[j].streak.MaxWagerWon),
			fmt.Sprintf("%dplayers.%ddecks.%dgames.players.player%d.bankroll.streak.MaxWagerLost", num_players, num_decks, games, j+1) : float64(playerBankRolls[j].streak.MaxWagerLost),
		}
		log.Printf("Send MetricsMap streak for player %d:\n%v\n", j, metricsMap)	
		sendGraphite(metricsMap)	
	}
	

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






/****
 * October '23
 * Preserving code here - this is unused in the multi-player version of the code
 */
func mainPreserveSingleplayerCode() {
	// figuring out proper way to log output
	fmt.Printf("fmt printf[version %s] Blackjack Simulator\n", version)
	log.Printf("[version %s] Blackjack Simulator\n", version)
	//dlog(fmt.Sprintf("dlog[version %s] Blackjack Simulator\n", version))
	dlog.Debug("dlog[version %s] Blackjack Simulator\n", version)
	
	printRuntimeVars()
	
	log.Printf("[version %s] Playing %d games per round.\n", version, int(games))
	
/*
dev code - take this out, was put in to test if graphite metrics are working	
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
	
	// Used to track stats for unique outcomes
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
		round := NewRound(deck.Shuffle(), num_players)

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
