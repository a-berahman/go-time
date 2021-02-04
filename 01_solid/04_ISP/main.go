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

// Sharpen increases the damage dealt by this sword using a whetstone.
func (Sword) Sharpen() {}

// MakeBlunt decreases the damage dealt by this sword due to constant use.
func (Sword) MakeBlunt() {}

// Drop places the sword on the ground allowing others to pick it up.
func (Sword) Drop() {}

// Damager is implemented by objects that can be used as weapons.
type Damager interface {
	Damage() int
}

// DamageReceiver is implemented by objects that can receive weapon damage.
type DamageReceiver interface {
	ApplyDamage(int)
}
type Monster struct{}

// func Attack(m *Monster, s *Sword)
// Attack deals damage to a monster using a sword.
func Attack(dr DamageReceiver, d Damager) {
	//...

}
