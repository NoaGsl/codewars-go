package main

import "fmt"

func main() {
	fmt.Printf("The winner is : %v\n", DeclareWinner(Fighter{"Lew", 10, 2}, Fighter{"Harry", 5, 4}, "Lew"))
}

type Fighter struct {
	Name            string
	Health          int
	DamagePerAttack int
}

func DeclareWinner(fighter1 Fighter, fighter2 Fighter, firstAttacker string) string {
	for fighter1.Health > 0 && fighter2.Health > 0 {
		if fighter1.Name == firstAttacker {
			firstAttacker = fighter2.Name
			fighter2.Health = fighter1.Damage(fighter2)
			fmt.Printf("%s attacks %s; %s now has %d health.\n", fighter1.Name, fighter2.Name, fighter2.Name, fighter2.Health)
		} else {
			firstAttacker = fighter1.Name
			fighter1.Health = fighter2.Damage(fighter1)
			fmt.Printf("%s attacks %s; %s now has %d health.\n", fighter2.Name, fighter1.Name, fighter1.Name, fighter1.Health)
		}
	}
	if fighter1.Health <= 0 {
		return fighter2.Name
	}
	return fighter1.Name
}

func (attacker Fighter) Damage(receiver Fighter) int {
	return receiver.Health - attacker.DamagePerAttack
}
