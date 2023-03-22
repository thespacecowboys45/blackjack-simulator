# Blackjack Simulator
 
Forked from: https://github.com/bradhe/go-blackjack









BUSINESS LOGIC:
- a random set of cards is created from a set # of decks (e.g. total cards)
- cards are drawn off the "top of the deck" -> the deck is a fifo queue
- currently, dealer stands on soft 17
- functionally, a "strategy" to "double" will also 
   - hit with that amount if more than 2 cards


THINGS not handled in original code:
Gameplay:
- splits ****** This is one major thing not covered
	concept: add [splits] section
- dealer hits soft-17
- betting
- statistics
- a good random number generator method
	


Priority 1
--------------
Graphite integration



	
Concepts:
Add batting average, and sharpe ratio

Keep track of:
win / loss / push
- find standard deviation (# of pushes in a row, etc)	

// Add flag variable to control # of decks to use

Concept:
Have random strategies for random players and see how
that affects the outcome

Concepts:
- find the best strategy which gives you the best win/loss ratio (batting average)
- find the best betting strategy which gives you the best sharpe ratio ($ won per win vs $ lost per loss)


Concept:
- run through ALL permutations of ALL possible actions 
---- create "in memory" a list of all possible combinations to fill
---- in the matrix of player actions

Concept:
- game plays "differently" with different statistics in a "real world" environment
   - test computer decision making against a real deck of cards
   
   

TRACKING statistics (things I want to track)
- # times hit on a total of 21 (or any other number/total) <- track for a "player mistake"
- if/when vary "stand on all 17's" or "hit soft 17", track
	# of times dealer dealt a soft-17
    # of times dealer busts when hitting 17
    # of times dealer improves hand when hitting 17

- Betting stats:
 - amount won
 - amount lost 
 - # of consecutive wins (per streak)
 - # of consecutive losses (per streak)
 - histogram output of # of losses in a row (e.g. had three (3) streaks with four (4) losses in a row)
 
 - random number generator effect on outcomes
 
Structure: winLossStreak -
	- determined by a 
		win followed by a loss 





TESTING SITUATIONS:
Strategy comparisons()
- betting strategy - flat bet

- betting strategy - martingale

** dev branch CHANGELOG
v1.3_concept - tracking stats
- add graphite interaction (tracking)

v1.2_concept - surgery notes
- strategies.go - mod translateAction (add "split")
- decks.go - override default_decks variable with flag
- add "deck #" to card (or something to track which deck it came from)
- recalculate "minimum_shoe_size" or add as a variable to the AI machine
-- add multi-players (rounds.go)
-- add "auto-strategy" for other players
-- add flag to split / not-split multiple times
-- add logic for "dealer stands on soft 17"
- determine what hands.go sumWithAlternatives, alternatesUsed variable is for
- add table limit (and player action if table limit supposed to be exceeded)



NATURAL GAME VARIABLES (things found in the real-world, hard to simulate):
- Dealer placing the cut card to stop drawing from the deck, random human determinate.
- Player bet wager (typically there is an idea behind how a player bets/alternates betting)
- if player treats split hands both won as two consecutive wins, or one win
- if player treats double down bet won as special (add/reduce)
- if player treats BLACKJACK hand won as special (to vary bet)
- if player bets an "odd" amount, and blackjack or doubledown won, payout is rounded down - track how much margin this produces the house
- if player will "doubledown" on a split
- number of players


END RESULT DESIRED:
- A realistic game, which can simulate all possible real actions of players including:
  hit, stand, double down, splitting
  
- A game which tracks a bankroll for a player given a determinate "betting strategy"

- A game which ends after a deterministic event
	 (# of hands played, 
	  bankroll limit reached,
	  too many drinks consumed,
	  etc.)
	  
- An implemented "betting strategy" which determines how a wager is placed 




(new) BETTING STRATEGIES:
Question: what other ways would you decide to vary your wager, or not?
a) win/loss streak <- 'n' number of wins / losses in a row
b) a "hot" deck <- card counting
c) last amount wagered / won <- say, if win "big bet" then take down or reduce
d) every 'n'th hand <- example: every 3rd hand I double my bet

WAYS TO INCREASE ODDS:
- Figure out when "odds are in your favor" <- if at any time, and take advantage of these.
Things to find out: what percentage of time does player action increase/decrease odds for
	- splits
	- doubledowns
	
	Like, track the number of times you won a doubldown vs lost
	Trake # of times a split was won, per hand
	
	

 


- PREMISE: take advantage of win/loss streak pattern to vary wager

- martingale

Structure:
	- initial bet amount (1st hand)
	- goose-style initial bet amount (1st hand, wild bet, just for fun because you are baller)

Concept:
	Depending on the win/loss outcome of the previous hand, change or don't change wager:
		- stand (keep betting level same)
		- reset (go back to minimum bet
		- decrease(amount or percentage)
		- increase(amount or percentage)
		
Possible reason to increase a bet:
	- lost last bet : 	(reset) - lost a "big bet"
						(reset) - win / loss streak determination
						(increase) - martingale strategy
	- won previous bet : (reset) - won a "big bet"
	                     	
	

CUSTOM BETTING STRATEGY (attempt to recreate martingale up to house limit)
Strategy: after any loss, double bet
  W W W W W W W W W W (number of consecutive wins)
L I I I I I I I I I I
L I
L I
L I
L I
L I
L I
L I
L I
L I
(number of consecutive losses)

S = same
I = increase
D = decrease
R = reset
# = integer multiplier (idea)
x = dont care



CUSTOM #2
- after 1 loss keep bet same
- after 2 losses in a row, double bet
- after 3 or more losses in a row, reset bet to default

I = INIT
W = WIN STREAK COUNT
L = LOSS STREAK COUNT
  I
  W W W W W W
L S I R R R R
L I x x x x x
L R x x x x x
L R x x x x x
L R x x x x x




CUSTOM #3
- after 1 win keep bet same
- after 2 consecutive wins, double bet
- after 3 consecutive wins, reset bet to default


(Need stats on consecutive wins/losses)





FUNCTION HOOKS (surgery planning):
- STEP 1 
+ add bankroll tracking - Okay
+ take 'outcome' and produce a resultant change in bank account - Okay

- STEP 2
create new bettingStrategy struct, similar to how the playerStrategy
is construtured.  Load up the strategy from a file on the computer.

- add ability to determine "when to stop drawing" the deck based on % of total cards
  in the deck.  More accurately, place a "stop" into the deck, just like the 
  dealer does with the card.
  
- need to modify the round so that if the player 'doubles' then bet is changed
	-- pass in the bet amount, which may be affected by the round: split/double/blackjack


- modify OUTCOME constants to add PUSH as result/outcome
	- PUSH
	- DOUBLEDOWN_WON
	
- add Multi-player	
- add card-counter capability
- add "AI" engine <- how the hell does this thing learn / adjust?
	
- FOLLOWUP: hands.go -> figure out IsSoft() function, is correct?


v1.1 - surgery notes
- add 'version #'
- add "printDeck()" to show the deck of cards being played from
- only "seed" the random number generator ONCE
- print hand total in verbose mode

v1.0 - modify to run on winblowz

MASTER CHANGE

I've always been fascinated by Blackjack. Some pros say that, if you follow a
basic strategy, your odds of winning go up significantly. So, that got me
wondering:

> If you follow a Blackjack strategy algorithmically, how will you do?

This tiny app is meant to address that.

## Disclaimers

1. I actually don't know very much at all about Blackjack.
1. There might be bugs, so the numbers may be incorrect.
1. Memorizing strategy (probably) doesn't make up for knowing the game.

## How does it work?

You author a strategy with a pretty straight forward DSL. The app will run this
strategy against a given number of games (default 100) and output how it does.

### Strategy DSL

Here's an example strategy.

```
[soft]
   2 3 4 5 6 7 8 9 10 A
13 H H H H H H H H  H H
14 H H H H H H H H  H H
15 H H H H H H H H  H H
16 H H H H H H H H  H H
17 S S S S S S S S  S S
18 S S S S S S S S  S S
19 S S S S S S S S  S S
20 S S S S S S S S  S S
21 S S S S S S S S  S S

[hard]
   2 3 4 5 6 7 8 9 10 A
 4 H H H H H H H H  H H
 5 H H H H H H H H  H H
 6 H H H H H H H H  H H
 7 H H H H H H H H  H H
 8 H H H H H H H H  H H
 9 H H H H H H H H  H H
10 H H H H H H H H  H H
11 H H H H H H H H  H H
12 H H H H H H H H  H H
13 H H H H H H H H  H H
14 H H H H H H H H  H H
15 H H H H H H H H  H H
16 H H H H H H H H  H H
17 S S S S S S S S  S S
18 S S S S S S S S  S S
19 S S S S S S S S  S S
20 S S S S S S S S  S S
21 S S S S S S S S  S S
```

The `[soft]` section describes soft-hand strategy. The `[hard]` section
describes hard-hand strategy.

The actions are described as follows.

```
H = hit
D = double
S = stand
```

You can run that strategy through the simulator like this.

```
$ ./go-blackjack --strategy=strategies/passive --games=10000
2013/11/09 22:31:07 Loading strategy strategies/passive
2013/11/09 22:31:09 Total Hands         551588
2013/11/09 22:31:09 Total Wins          213924  (38.783%)
2013/11/09 22:31:09 Total Losses        277828  (50.369%)
2013/11/09 22:31:09 Total Pushes        49836   (9.035%)
```

## Does it actually work??

I dunno. So far, I've tried two different strategies and here are my results for each.

### Passive Strategy

This strategy is checked in to the repo.

```
$ ./go-blackjack --strategy=strategies/passive --games=100000
2013/11/09 22:32:12 Loading strategy strategies/passive
2013/11/09 22:32:33 Total Hands         5515165
2013/11/09 22:32:33 Total Wins          2141896 (38.836%)
2013/11/09 22:32:33 Total Losses        2780783 (50.421%)
2013/11/09 22:32:33 Total Pushes        492486  (8.930%)
```

### Wizard of Odds Strategy

This strategy is also checked in to the repo and described on the [Wizard of
Odds](http://wizardofodds.com/games/blackjack/) website.

**NOTE:** One big missing piece that is described in the Wizard of Odds
strategy that is missing here is splitting. This simulator does not support it!

```
$ ./go-blackjack -strategy strategies/wizard_simple -games 100000
2013/11/09 22:25:13 Loading strategy strategies/wizard_simple
2013/11/09 22:25:33 Total Hands       5562401
2013/11/09 22:25:33 Total Wins        2275528 (40.909%)
2013/11/09 22:25:33 Total Losses      2761467 (49.645%)
2013/11/09 22:25:33 Total Pushes      425406  (7.648%)
```

## Assumptions

1. `/dev/urandom` is sufficiently random for our purposes.
1. Not shuffling between hands is OK.
1. Simulates a six-deck shoe by default.

## Contributing

You know what do! Fork and submit a pull request. Strategies are, of course,
welcome as well.
