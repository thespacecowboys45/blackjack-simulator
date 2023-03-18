package main

import(
"fmt"
)

// Players start with a default amount to wager with
const DEFAULT_BANKROLL = 500
var bankRoll int

const DEFAULT_WAGER = 5
var currentWager int


type BankRoll struct {
	Amount int
	Max int
	Min int
}

type Wager struct {
	Amount int
}

func (bankRoll BankRoll) String() string {
	return fmt.Sprintf("Amount: %d Min: %d Max: %d", bankRoll.Amount,
	bankRoll.Min,
	bankRoll.Max)
}

func (wager Wager) String() string {
	return fmt.Sprintf("%d", wager.Amount)
}

func NewBankRoll(amount int) BankRoll {
	br := BankRoll{}
	br.Amount = amount
	br.Min = 0
	br.Max = amount
	
	fmt.Printf("NewBankRoll: %v\n", br)
	return br
}

func (wager Wager) NewWager(outcome Outcome) Wager {
	
	wg := Wager{}
	// Initialize from current object
	wg.Amount = wager.Amount
	
	// Implements betting strategy
	if outcome == OUTCOME_INIT {
		wg.Amount = DEFAULT_WAGER	
	} else if outcome == OUTCOME_WIN {
		// @TODO - do some logic here
		fmt.Printf("\tNewWager() OUTCOME_WIN - reset bet to default\n")
		wg.Amount = DEFAULT_WAGER
	} else if outcome == OUTCOME_LOSS {
		fmt.Printf("\tNewWager() OUTCOME_LOSS - double bet\n")
		// DEV - basically martingale, always double when loosing
		wg.Amount = wg.Amount * 2
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
		bankRoll.Amount += wager.Amount
	} else if outcome == OUTCOME_LOSS {
		bankRoll.Amount -= wager.Amount
	} else {
		// push (non-event)
		bankRoll.Amount = bankRoll.Amount
	}
	
	// Because we cannot modify the object in here
	nbr.Amount = bankRoll.Amount
	nbr.Min = bankRoll.Min
	nbr.Max = bankRoll.Max
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