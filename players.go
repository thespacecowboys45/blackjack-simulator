package main
/*
 * @date Oct '23
 * @author dxb The Space Cowboy
 *
 * DESCRIPTION:
 *   Implements individual players playing the game
 *
 */

import(
	dlog "bitbucket.org/thespacecowboys45/dlogger"
)


// The number of players sitting at the table
const MIN_PLAYERS = 1
// Max number of players sitting at any one time
const MAX_PLAYERS = 6

// in dev, not used right now
type Player struct {
	// phase 1 - get this working
	// @TODO - deprecate
	//Hand Hand
	//Outcome Outcome
	
	// phase 2 - add hands and outcome for hands to handle splits
	splitsPlayed int
	activeHand int
	Hands []Hand
	Outcomes []Outcome
	
	// Statistics tracking
	Wager Wager
	BankRoll BankRoll
	Streak Streak
}

// Adds an empty hand to the players current set of hands
func (player Player) AddHand(hand Hand) []Hand {
	return append(player.Hands, hand)
}

// Adds an empty outcome to the players current set of outcomes
func (player Player) AddOutcome(outcome Outcome) []Outcome {
	return append(player.Outcomes, outcome)
}


func (player Player) toString() {
	dlog.LogEvent("[players.go][toString()][entry]", "trace")
	//dlog.Always("Player hand (deprecate): %v", player.Hand)
	//dlog.Always("Player outcome (deprecate): %s", outcomeToString(player.Outcome))
	dlog.Always("Player splits played: %d", player.splitsPlayed)
	dlog.Always("Player active hand: %d", player.activeHand)
	dlog.Always("Player %d / %d Hands: %v", len(player.Hands), cap(player.Hands), player.Hands)
	for j:=0; j<len(player.Hands); j++ {
		dlog.Always("Player Hand #%d total: %d outcome: %s", j, player.Hands[j].Sum(), outcomeToString(player.Outcomes[j]))
	}
	//dlog.Always("Player %d / %d Outcomes: %v", len(player.Outcomes), cap(player.Outcomes), player.Outcomes)
	
	dlog.LogEvent("[players.go][toString()][exit]", "trace")
	
}

//func (player Player) NewPlayer() *Player {
func (player Player) NewPlayer() {
	//player := new(Player)
	
	// initial code deals with one-hand only
	//player.Hand = Hand{}
	// assume one hand only
	player.activeHand = 0 
	
	// phase 2 - use a slice to handle split possibility (multiple-hands per game)
	// works
	player.Hands = make([]Hand, 1)
	
	// alternative way
	//player.Hands = player.Hands.AddHand(Hand{})
	
	// or (idk if this'll work tho ...)
	//player.Hands[0] = Hand{}
	
//	return player
}