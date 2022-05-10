package main

import (
	"fmt"
)

type Sword struct {
	name string
}

type Bow struct {
	name string
}

type Item interface {
	cost() int
}
type Weapon interface {
	detail() string
	Item
}

func attack(w Weapon, name string) {
	fmt.Printf("weapon %s: %s, cost: %d\n", name, w.detail(), w.cost())
}

func (w Sword) detail() string {
	return "detail sword"
}

func (w Sword) cost() int {
	return 100
}

func (w Bow) detail() string {
	return "detail bow"
}

func (w Bow) cost() int {
	return 50
}

func main() {
	sword := Sword{name: "sword"}
	bow := Bow{name: "bow"}

	attack(sword, sword.name)
	attack(bow, bow.name)
	// or
	// var w = sword
	// attack(w)
}
