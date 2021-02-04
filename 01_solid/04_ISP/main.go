package main

import "fmt"

type Weaponer interface {
	String() string
	Sharpen()
	MakeBlunt()
	Drop()
}

//Sword is a type of weapons
type Sword struct {
	name string // Important tip for RPG players: always name your swords!
}

// Damage returns the damage dealt by this sword.
func (Sword) Damage() int {
	return 2
}

// String implements fmt.Stringer for the Sword type.
func (s Sword) String() string {
	return fmt.Sprintf("%s is a sword that can deal %d points to oppnents", s.name, s.Damage())
}

func (Sword) Sharpen()   {}
func (Sword) MakeBlunt() {}
func (Sword) Drop()      {}

type Damager interface {
	Damage() int
}

type DamageReceiver interface {
	ApplyDamage(int)
}

func Attack(dr DamageReceiver, d Damager) {
	//...

}

type Monster struct{}
