package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	dlog "bitbucket.org/thespacecowboys45/dlogger"

)

type BettingStrategy interface {
	// Gets the action that we want to perform.
	GetBettingAction(consecutiveLosses int, consecutiveWins int) BettingAction
}

type internalBettingStrategy struct {
	streakStrategies map[string]map[string]BettingAction
	nhandsStrategies map[string]map[string]BettingAction
}

// DAVB description:
// Determines if the player will HIT or STAND or DOUBLE based on the value in the strategy matrix
// and:
// a) is the player holding a "soft hand" - use soft strategy
// b) is the player holding two or more cards and strategy calls to "double"
//
//func (self *internalBettingStrategy) GetBettingAction(player, dealer Hand) BettingAction {
func (self *internalBettingStrategy) GetBettingAction(consecutiveLosses int, consecutiveWins int ) BettingAction {
	// TODO: We'll need a smarter way to look up actions from our strategies than
	// this...
	dlog.LogEvent("[bettingstrategies.go][GetBettingAction()][entry]\n", "trace")
	
	// -----> need to figure out how to grab this value for a bet
	// This is going to be how far down we are in the "consecutive Losses" stack
	// and or the WINS stack - have to figure this out
	
	// This is going to be "how far down" in the column to look for action
	lossesKey := fmt.Sprintf("L%d", consecutiveLosses)
	winsKey := fmt.Sprintf("W%d", consecutiveWins)

	var action BettingAction
	
	// action is of type BettingAction
	action = self.streakStrategies[lossesKey][winsKey]
	dlog.LogEvent(fmt.Sprintf("[bettingstrategies.go][GetBettingAction()]Determine action lossesKey: %s winsKey: %s ==> %s\n", lossesKey, winsKey, bettingActionToString(action)), "bettingstrategy")

/*	
	for k, v := range self.streakStrategies[lossesKey] {
		fmt.Printf("--> %s , %d --> key: %d\n", k, v, self.streakStrategies[lossesKey][k])	
	}
*/	

	// validate action - possible bug point if too many consecutive losses/wins and is beyond scope of the maxtrix found in the betting strategy file
	dlog.LogEvent("[bettingstrategies.go][GetBettingAction()][exit]", "trace")
	return action
}

func translateBettingAction(action string) BettingAction {
	action = strings.ToLower(action)

	if action == "r" {
		return BETTINGACTION_RESET
	} else if action == "i" {
		return BETTINGACTION_INCREASE
	} else if action == "d" {
		return BETTINGACTION_DECREASE
	}
	
	// covers 's'

	// TODO: What is the default action??
	return BETTINGACTION_STAND
}

/*
 dup in rounds.go
 
func bettingActionToString(action BettingAction) string {
	switch (action) {
		case BETTINGACTION_RESET:
			return "BETTINGACTION_RESET"
			break
		case BETTINGACTION_INCREASE:
			return "BETTINGACTION_INCREASE"
			break
		case BETTINGACTION_DECREASE:
			return "BETTINGACTION_DECREASE"
			break
		case BETTINGACTION_STAND:
			return "BETTINGACTION_STAND"
			break
		default:
			return fmt.Sprintf("unknown bettion action: %v", action)
			break
	}
	return ""
}

*/

func loadBettingStrategy(reader *bufio.Reader) map[string]map[string]BettingAction {
	// For holding the progression of # of wins in a row...
	winStreak := make([]string, 0)
	bettingstrategy := make(map[string]map[string]BettingAction)

	for {
		line, err := reader.ReadString('\n')
		msg := fmt.Sprintf("line: %s\n", line)
		dlog.LogEvent(msg, "bettingstrategy")

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		line = strings.TrimSpace(line)

		if len(winStreak) == 0 {
			// We need to load up the win streak lines
			toks := strings.Split(line, " ")

			for _, tok := range toks {
				msg := fmt.Sprintf("tok: %s\n", tok)
				dlog.LogEvent(msg, "bettingstrategy")
				winStreak = append(winStreak, tok)
			}
			
			dlog.LogEvent(fmt.Sprintf("winStreak final: %v\n", winStreak), "bettingstrategy")
		} else if line == "" || strings.HasPrefix(line, "#") {
			break
		} else {
			// This line describes a strategy, so let's pull it
			// apart. First token is going to be the scenario.
			toks := strings.Split(line, " ")
			
			// AH HA - the scenario is the "1st column" value 
			scenario, actions := toks[0], toks[1:len(toks)-1]
			dlog.LogEvent(fmt.Sprintf("scenario: %s\n", scenario), "bettingstrategy")
			
			dlog.LogEvent(fmt.Sprintf("Parsing strategy line.\n"), "bettingstrategy")
			dlog.LogEvent(fmt.Sprintf("Toks here: %s\n", toks), "bettingstrategy")

			// We'll need a new map here...
			data := make(map[string]BettingAction)

			// To keep of how many we've seen.
			idx := 0

			// ...and now let's load 'er up.
			for _, action := range actions {
				// Skip blank tokens...
				if strings.TrimSpace(action) == "" {
					continue
				}

				dlog.LogEvent(fmt.Sprintf("[%d] Add action: %s\n", idx, action), "bettingstrategy")
				data[winStreak[idx]] = translateBettingAction(action)
				dlog.LogEvent(fmt.Sprintf("winStreak[idx] = %v\n", winStreak[idx]), "bettingstrategy")

				dlog.LogEvent(fmt.Sprintf("Data now: %v\n", data), "bettingstrategy")

				// Gotta keep track of this outselves because we can't trust i here.
				idx += 1
			}

			bettingstrategy[scenario] = data
			dlog.LogEvent(fmt.Sprintf("for scenario %s betting strategy: %v\n", scenario, bettingstrategy[scenario]), "bettingstrategy")
		}
	}

	return bettingstrategy
}

// Loads the relevant strategy in from memory.
func LoadBettingStrategy(path string) BettingStrategy {
	log.Printf("Loading betting strategy %s", path)

	// Let's see if we can read the file.
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	// We got it, so let's get goin'
	defer file.Close()

	bettingstrategy := new(internalBettingStrategy)

	reader := bufio.NewReader(file)

	// Read the whole damn thing in.
	for {
		// Start by getting the headers.
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			dlog.LogEvent(fmt.Sprintf("End of file reached\n"), "bettingstrategy")
			break
		} else if err != nil {
			panic(err)
		}
		
		dlog.LogEvent(fmt.Sprintf("line: %s\n", line), "bettingstrategy")
		

		// If the line starts with a # it's a comment.
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "#") {
			continue
		} else if line == "" {
			// Empty line, nothing to see here.
			continue
		} else if line == "[streakvariant]" {
			dlog.LogEvent(fmt.Sprintf("Foudn streakvariant betting strategy\n"), "bettingstrategy")
			bettingstrategy.streakStrategies = loadBettingStrategy(reader)
		} else if line == "[nhandsvariant]" {
			dlog.LogEvent(fmt.Sprintf("Foudn nhandsvariant betting strategy\n"), "bettingstrategy")
			bettingstrategy.nhandsStrategies = loadBettingStrategy(reader)
		}
	}

	return bettingstrategy
}
