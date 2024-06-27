package main

import "fmt"

type Fighter struct {
	Name            string
	Health          int
	damagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	if fighter1.Name == firstAttacker {
		for fighter1.Health >= 0 || fighter2.Health >= 0 {
			fighter2.Health -= fighter1.damagePerAttack
			if fighter2.Health <= 0 {
				return fighter1.Name
			}
			fighter1.Health -= fighter2.damagePerAttack
			if fighter1.Health <= 0 {
				return fighter2.Name
			}
		}
	} else {
		for fighter1.Health >= 0 || fighter2.Health >= 0 {
			fighter1.Health -= fighter2.damagePerAttack
			if fighter1.Health <= 0 {
				return fighter2.Name
			}
			fighter2.Health -= fighter1.damagePerAttack
			if fighter2.Health <= 0 {
				return fighter1.Name
			}
		}
	}
	return ""
}

func main() {
	fmt.Println(DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Harry"))
}
