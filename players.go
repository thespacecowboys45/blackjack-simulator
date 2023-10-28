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
const DEFAULT_PLAYERS = 1
// Max number of players sitting at any one time
const MAX_PLAYSER = 6

// in dev, not used right now
type Player struct {
	// phase 1 - get this working
	Hand Hand
	Outcome Outcome
	
	// phase 2 - add hands for splits
	Hands []Hand
	// Outcomes []Outcome
	
	
	// Statistics tracking
	Wager Wager
	BankRoll BankRoll
	Streak Streak
}