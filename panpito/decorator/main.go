package main

import "log"

type InterfaceSolider interface {
	Attack() int
	Defense() int
}

type BasicSoldier struct{}

func (b BasicSoldier) Attack() int  { return 1 }
func (b BasicSoldier) Defense() int { return 1 }
func DisplayState(solider InterfaceSolider) {
	log.Printf("Soldier state: attack %d, defense %d",
		solider.Attack(), solider.Defense())
}

type SoldierWithSword struct {
	solider InterfaceSolider
}

// Attack implements InterfaceSolider.
func (s SoldierWithSword) Attack() int {
	return s.solider.Attack() + 1
}

// Defense implements InterfaceSolider.
func (s SoldierWithSword) Defense() int {
	return s.solider.Defense() + 1
}

type SoldierWithShield struct {
	solider InterfaceSolider
}

// Attack implements InterfaceSolider.
func (s SoldierWithShield) Attack() int {
	return s.solider.Attack() - 6
}

// Defense implements InterfaceSolider.
func (s SoldierWithShield) Defense() int {
	return s.solider.Defense() + 20
}

func main() {
	basicSoldier := BasicSoldier{}
	DisplayState(basicSoldier)

	soldierWithSword := SoldierWithSword{solider: basicSoldier}
	DisplayState(soldierWithSword)

	soldierWithShield := SoldierWithShield{solider: basicSoldier}
	DisplayState(soldierWithShield)
}
