package blackjack

import "slices"
// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	// panic("Please implement the ParseCard function")
    switch card {
        case "ace":
        return 11
        case "two":
        return 2
        case "three":
        return 3
        case "four":
        return 4 
        case "five":
        return 5 
        case "six":
        return 6 
        case "seven":
        return 7 
        case "eight":
        return 8 
        case "nine":
        return 9
        case "ten", "jack", "queen","king":
        return 10 
        default:
        return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	// panic("Please implement the FirstTurn function")
    switch {
        case card1 == "ace" && card2 == "ace":
        return "P"
        case ParseCard(card1) + ParseCard(card2) == 21 && !slices.Contains([]string{"ace","jack", "queen","king","ten"},dealerCard)  :
        return "W"
        case ParseCard(card1) + ParseCard(card2) == 21 && slices.Contains([]string{"ace","jack", "queen","king","ten"},dealerCard):
        return "S"
        case slices.Contains([]int{17,18,19,20}, ParseCard(card1) + ParseCard(card2)):
        return "S"
        case slices.Contains([]int{12,13,14,15,16}, ParseCard(card1) + ParseCard(card2)) && ParseCard(dealerCard) >= 7:
        return "H"
        case slices.Contains([]int{12,13,14,15,16}, ParseCard(card1) + ParseCard(card2)):
        return "S"
        case ParseCard(card1) + ParseCard(card2) <= 11:
        return "H"
    }
return ""    
}
