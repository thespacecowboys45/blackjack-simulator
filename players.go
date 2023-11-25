package main
/*
 * @date Oct '23
 * @author dxb The Space Cowboy
 *
 * DESCRIPTION:
 *   Implements individual players playing the game
 *
 */
 
// The number of players sitting at the table
const MIN_PLAYERS = 1
// Max number of players sitting at any one time
const MAX_PLAYERS = 6

// in dev, not used right now
type Player struct {
	// phase 1 - get this working
	Hand Hand
	Outcome Outcome
	
	// phase 2 - add hands and outcome for hands to handle splits
	Hands []Hand
	HandOutcomes []Outcome
	
	activeHand int
	// Outcomes []Outcome
	
	
	// Statistics tracking
	Wager Wager
	BankRoll BankRoll
	Streak Streak
}

// Adds an empty hand to the players current set of hands
// This is so a player can play multiple hands
func (player Player) AddHand(hand Hand) []Hand {
	return append(player.Hands, hand)
}

//func (player Player) NewPlayer() *Player {
func (player Player) NewPlayer() {
	//player := new(Player)
	
	// initial code deals with one-hand only
	player.Hand = Hand{}
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