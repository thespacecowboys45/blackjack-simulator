package main
// @date 2021
// @date 2023
// @author The Space Cowboy
//
/////
import(
"fmt"
)

// Players start with a default amount to wager with
const DEFAULT_BANKROLL = 500
var bankRoll int

const DEFAULT_WAGER = 5
var currentWager int


type Streak struct {
	LastOutcome Outcome
	Wins int
	Losses int
	// current streak
	ConsecutiveWins int
	ConsecutiveLosses int
	// @TODO - treat this as a histogram-like variable
	MaxConsecutiveWins int
	MaxConsecutiveLosses int
	
}

func (streak Streak) init() Streak {
	fmt.Printf("Streak - init()\n")
	s := Streak{}
	s.LastOutcome = OUTCOME_INIT
	s.Wins = 0
	s.Losses = 0
	s.ConsecutiveWins = 0
	s.ConsecutiveLosses = 0
	s.MaxConsecutiveWins = 0
	s.MaxConsecutiveLosses = 0
	return s
}

// keeping the code simple/readable
func (streak Streak) addWin() Streak {
	s := Streak{}
	if (streak.LastOutcome == OUTCOME_INIT) {
		s.ConsecutiveWins = 1
	} else if (streak.LastOutcome == OUTCOME_WIN) {
		// win then win streak
		s.ConsecutiveWins = streak.ConsecutiveWins + 1		
	} else if (streak.LastOutcome == OUTCOME_LOSS) {
		// reset
		s.ConsecutiveWins = 1
		
	}
	s.LastOutcome = OUTCOME_WIN
	s.Wins = streak.Wins + 1
	s.Losses = streak.Losses
	
	if s.ConsecutiveWins > streak.ConsecutiveWins {
		// NEW High count
		s.MaxConsecutiveWins = s.ConsecutiveWins
	} else {
		// Count remains the same
		s.MaxConsecutiveWins = streak.MaxConsecutiveWins
	}
	
	// keep the same
	//s.ConsecutiveLosses = streak.ConsecutiveLosses
	// No, reset to 0
	s.ConsecutiveLosses = 0
	s.MaxConsecutiveLosses = streak.MaxConsecutiveLosses
	
	fmt.Printf("addWin() %s\n", s.String())
	return s
}

// keeping the code simple/readable
func (streak Streak) addLoss() Streak {
	s := Streak{}
	fmt.Printf("addLoss() lastOutcome: %d\n", streak.LastOutcome)
	if (streak.LastOutcome == OUTCOME_INIT) {
		s.ConsecutiveLosses = 1
	} else if (streak.LastOutcome == OUTCOME_WIN) {
		// reset
		s.ConsecutiveLosses = 1		
	} else if (streak.LastOutcome == OUTCOME_LOSS) {
		// loss then loss streak
		s.ConsecutiveLosses = streak.ConsecutiveLosses + 1
	}
	s.LastOutcome = OUTCOME_LOSS
	s.Losses = streak.Losses + 1
	s.Wins = streak.Wins
	
	fmt.Printf("addLoss() SCOMPARATOR: %s\n", s.String())
	fmt.Printf("addLoss() SCOMPARATOR2: %s\n", streak.String())	
	if s.ConsecutiveLosses > streak.MaxConsecutiveLosses {
		// NEW High count
		s.MaxConsecutiveLosses = s.ConsecutiveLosses		
	} else {
		// Count remains the same
		s.MaxConsecutiveLosses = streak.MaxConsecutiveLosses
	}
	
	// keep the same
	//s.ConsecutiveWins = streak.ConsecutiveWins
	// No, reset to 0
	s.ConsecutiveWins = 0
	s.MaxConsecutiveWins = streak.MaxConsecutiveWins

	fmt.Printf("addLoss() FINAL     : %s\n", s.String())
	return s
}


func (s Streak) String() string {
	return fmt.Sprintf("LastOutcome: %d Wins: %d Losses: %d CWins: %d CLosses: %d MAXCWins: %d MAXCLosses: %d\n",
		s.LastOutcome, 
		s.Wins,
		s.Losses,
		s.ConsecutiveWins,
		s.ConsecutiveLosses,
		s.MaxConsecutiveWins,
		s.MaxConsecutiveLosses)
}

type BankRoll struct {
	Amount int
	Max int
	Min int
	streak Streak
}

type Wager struct {
	Amount int
	// @TODO - add
	MaxWager int
	MinWager int
	// @TODO - add
	//wagerCount map
}

func (bankRoll BankRoll) String() string {
	return fmt.Sprintf("Amount: %d Min: %d Max: %d Streak: %s", bankRoll.Amount,
	bankRoll.Min,
	bankRoll.Max,
	bankRoll.streak.String())
}

func (wager Wager) String() string {
	return fmt.Sprintf("%d", wager.Amount)
}

func NewBankRoll(amount int) BankRoll {
	br := BankRoll{}
	br.Amount = amount
	br.Min = 0
	br.Max = amount
	br.streak = Streak{}
	br.streak = br.streak.init()
	
	fmt.Printf("NewBankRoll: %v\n", br)
	return br
}


// Somehow this is not right.  It should be more tied closely with getBettingAction
// I think .... 
//
// outcome: the outcome of the last bet.  If it is the 1st time around
//          outcome will be set to OUTCOME_INIT
//
func (wager Wager) NewWager(outcome Outcome, streak Streak, determineBet func(streak Streak) BettingAction) Wager {
	
	wg := Wager{}
	// Initialize from current object
	wg.Amount = wager.Amount
	
	// Implements betting strategy
	if outcome == OUTCOME_INIT {
		wg.Amount = DEFAULT_WAGER	
	} else if outcome == OUTCOME_WIN {
		fmt.Printf("\tNewWager OUTCOME_WIN - check out the streak, how many wins in a row? => %d\n", streak.ConsecutiveWins)
	} else if outcome == OUTCOME_LOSS {
		fmt.Printf("\tNewWager OUTCOME_LOSS - check out the streak, how many losses in a row? => %d\n", streak.ConsecutiveLosses)
	}
			
	nextAction := determineBet(streak)	
	switch (nextAction) {
		default:
			fmt.Printf("UNHANDLED betting strategy action: %d\n", nextAction)
		case BETTINGACTION_RESET:
			break
		case BETTINGACTION_INCREASE:
			fmt.Printf("BETTINGACTION_INCREASE - double bet\n")
			// DEV - basically martingale, always double when loosing
			wg.Amount = wg.Amount * 2
			break
		case BETTINGACTION_DECREASE:
			fmt.Printf("BETTINGACTION_DECREASE - half bet\n")
			wg.Amount = wg.Amount / 2
			break	
		case BETTINGACTION_STAND:
			fmt.Printf("BETTINGACTION_STAND - keep bet the same\n")
			wg.Amount = wager.Amount
			break		
	}
	
	fmt.Printf("Compare: wg.Amount: %d to DEFAULT_WAGER: %d\n", wg.Amount, DEFAULT_WAGER)
	// Can never go below our initial wager amount
	if wg.Amount < DEFAULT_WAGER {
		fmt.Printf("?SF#?R#RLK#JFJELKJFEKL How did we got here??????\n")
		wg.Amount = DEFAULT_WAGER
	}
	
	return wg
}

func (bankRoll BankRoll) tallyOutcome(outcome Outcome, wager Wager) BankRoll {
	nbr := BankRoll{}
	msg := fmt.Sprintf("tallyOutcome entry: %d\t current bankRoll: %s\twager: %d\n", 
		outcome,
		bankRoll.String(),
		wager.Amount)
	dlog(msg)
			
	if outcome == OUTCOME_WIN {
		fmt.Printf("tallyOutcome COUNT WIN\n")
		bankRoll.Amount += wager.Amount
		bankRoll.streak = bankRoll.streak.addWin()
	} else if outcome == OUTCOME_LOSS {
		fmt.Printf("tallyOutcome COUNT LOSS\n")
		bankRoll.Amount -= wager.Amount
		bankRoll.streak = bankRoll.streak.addLoss()
	} else {
		// push (non-event)
		fmt.Printf("talyOutcome - PUSH")
		bankRoll.Amount = bankRoll.Amount
		
		// Keep track record the same (non-event)
		nbr.streak.ConsecutiveWins = bankRoll.streak.ConsecutiveWins
		nbr.streak.ConsecutiveLosses = bankRoll.streak.ConsecutiveLosses
		
	}
	
	fmt.Printf("BANKROLL STREAK: %s\n", bankRoll.streak.String())
	
	// Because we cannot modify the object in here
	nbr.Amount = bankRoll.Amount
	nbr.Min = bankRoll.Min
	nbr.Max = bankRoll.Max
	// @TODO - there has to be a more structured way, it's late, i'm tired
	nbr.streak = Streak{}
	nbr.streak.LastOutcome = bankRoll.streak.LastOutcome
	nbr.streak.Wins = bankRoll.streak.Wins
	nbr.streak.Losses = bankRoll.streak.Losses
	nbr.streak.ConsecutiveWins = bankRoll.streak.ConsecutiveWins
	nbr.streak.ConsecutiveLosses = bankRoll.streak.ConsecutiveLosses
	nbr.streak.MaxConsecutiveWins = bankRoll.streak.MaxConsecutiveWins
	nbr.streak.MaxConsecutiveLosses = bankRoll.streak.MaxConsecutiveLosses
		
	fmt.Printf("DEBUG compare: nbr.Amount: %d bankRoll.Max: %d\n", nbr.Amount, bankRoll.Max)
	if nbr.Amount > bankRoll.Max {
		fmt.Printf("\tNEW MAX\n")
		nbr.Max = nbr.Amount
	}
	
	if nbr.Amount < bankRoll.Min {
		fmt.Printf("\tNEW MIN\n")
		nbr.Min = nbr.Amount
	} 
	
	msg = fmt.Sprintf("tallyOutcome exit: %d\t new bankRoll: %s\twager: %d\n", 
		outcome,
		nbr.String(),
		wager.Amount)
	dlog(msg)
	
	// new bankroll amount
	return nbr
}