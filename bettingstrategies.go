package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type BettingStrategy interface {
	// Gets the action that we want to perform.
	GetBettingAction(streak Streak, outcome Outcome) BettingAction
}

func (self *internalBettingStrategy) Stub() {
	fmt.Printf("BettingStrategy stub()\n")	
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
func (self *internalBettingStrategy) GetBettingAction(streak Streak, outcome Outcome) BettingAction {
	// TODO: We'll need a smarter way to look up actions from our strategies than
	// this...
	
	// -----> need to figure out how to grab this value for a bet
	// This is going to be how far down we are in the "consecutive Losses" stack
	// and or the WINS stack - have to figure this out
	
	// This is going to be "how far down" in the L's column to look for action
	lossesKey := fmt.Sprintf("%d", streak.ConsecutiveLosses)

	// Need some special rules for this one, to deal with bets.
	var winsKey string

	// DAVB - revamp this to fit a "bet" column - do we need this exception ? 
	/*
	if dealer[0].Symbol == CARD_ACE {
		dealerKey = "A"
	} else {
		dealerKey = fmt.Sprintf("%d", dealer[0].Value)
	}
	*/

	var action BettingAction

/*
	if player.IsSoft() {
		if val, ok := self.softStrategies[lossesKey][winsKey]; ok {
			action = val
		} else {
			// No soft strategy available.
			action = self.hardStrategies[lossesKey][winsKey]
		}
	} else {
		action = self.hardStrategies[lossesKey][winsKey]
	}
*/	
	action = self.streakStrategies[lossesKey][winsKey]

/*
	// If the player's hand has more than 2 cards and the action the strategy
	// calls for is double, we'll hit instead.
	if action == ACTION_DOUBLE && len(player) > 2 {
		action = ACTION_HIT
	}
*/
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

func loadBettingStrategy(reader *bufio.Reader) map[string]map[string]BettingAction {
	// For holding the progression of # of wins in a row...
	winStreak := make([]string, 0)
	bettingstrategy := make(map[string]map[string]BettingAction)

	for {
		line, err := reader.ReadString('\n')
		msg := fmt.Sprintf("line: %s\n", line)
		dlog(msg)

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
				dlog(msg)
				winStreak = append(winStreak, tok)
			}
		} else if line == "" || strings.HasPrefix(line, "#") {
			break
		} else {
			// This line describes a strategy, so let's pull it
			// apart. First token is going to be the scenario.
			toks := strings.Split(line, " ")
			scenario, actions := toks[0], toks[1:len(toks)-1]

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

				data[winStreak[idx]] = translateBettingAction(action)

				// Gotta keep track of this outselves because we can't trust i here.
				idx += 1
			}

			bettingstrategy[scenario] = data
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
			break
		} else if err != nil {
			panic(err)
		}

		// If the line starts with a # it's a comment.
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, "#") {
			continue
		} else if line == "" {
			// Empty line, nothing to see here.
			continue
		} else if line == "[streakvariant]" {
			bettingstrategy.streakStrategies = loadBettingStrategy(reader)
		} else if line == "[nhandsvariant]" {
			bettingstrategy.nhandsStrategies = loadBettingStrategy(reader)
		}
	}

	return bettingstrategy
}
