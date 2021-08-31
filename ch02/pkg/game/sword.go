package game

import "fmt"

type Sword struct {
	name string
}

func (s Sword) Damage() int {
	return 2
}

func (s Sword) String() string {
	return fmt.Sprintf(
		"%s is a sword that can deal %d points of damage to opponents",
		s.name, s.Damage())
}

type EnchantedSword struct {
	Sword
}

func (receiver EnchantedSword) Damage() int {
	return 42
}
