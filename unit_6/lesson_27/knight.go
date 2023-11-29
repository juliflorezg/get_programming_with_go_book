package main

import "fmt"

type item string

type character struct {
	name     string
	leftHand *item
}

func (c *character) pickUp(i *item) {
	if i == nil {
		fmt.Println("there's no item")
		return
	}

	fmt.Printf("character %v has picked up %v\n", c.name, *i)
	c.leftHand = i
}

func (c *character) giveTo(to *character) {
	fmt.Println("character A:", c)
	fmt.Println("character B:", to)
	if c == nil {
		fmt.Println("no character that can perform this action (character A)")
		return
	}

	if c.leftHand == nil {
		fmt.Println("This character has nothing to give")
		return
	}

	if to.leftHand != nil {
		fmt.Printf("Character %v already has an item\n", c.name)
		return
	}

	if to.name == "" || to == nil {
		fmt.Println("no character that this action can be applied to (character B)")
		return
	}

	to.leftHand = c.leftHand

	fmt.Printf("character A (%v) has given this item: '%v' to character B (%v)\n", c.name, *c.leftHand, to.name)

	c.leftHand = nil
}

func knights() {
	var character1 character

	character1.pickUp(character1.leftHand)

	var item1 item = "sword"

	arthur := character{
		name:     "Arthur",
		leftHand: &item1,
	}
	fmt.Println("->"+arthur.name+"<-", arthur.leftHand, *arthur.leftHand)

	arthur.giveTo(&character1) // fails bc character 1 has not been initialized (no name)

	character1.name = "lol"

	arthur.giveTo(&character1)
	fmt.Println("->"+arthur.name+"<-", arthur.leftHand)                               // ->Arthur<- <nil>
	fmt.Println("->"+character1.name+"<-", character1.leftHand, *character1.leftHand) // ->lol<- 0xc00008a0c0 sword

}
