package main

import (
	"fmt"
	"strings"
)

// the interface type is concerned with what a type can do and not the value it holds
var t interface {
	talk() string
}

//* types martian and laser satisfy the t interface

type martian struct{}

func (m martian) talk() string {
	return "nack nack"
}

type laser int

func (l laser) talk() string {
	return strings.Repeat("ups, failure", int(1))
	// return strings.Repeat("lol ", 1)
}

// as a convention an interface is declared with a suffix -er, meaning that a talker is anything that talks
type talker interface {
	talk() string
}

func shout(t talker) {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

type crater struct{}

type rover string

func (r rover) talk() string {
	return "whir whir"
}

//? using interfaces with struct embedding 
type starship struct {
	laser
}

func main() {
	t = martian{}
	fmt.Println(t.talk())
	t = laser(3)
	fmt.Println(t.talk())

	// t variable can change shape (from a martian to a laser).
	// Computer scientists say that interfaces provide polymorphism, which means “many shapes.”

	// using the shout function with types that satisfy the talker interface
	shout(martian{})
	shout(laser(2))
	// shout(crater{}) //! error: crater does not implement the talker interface (missing method talk)

	r := rover("whir whir")
	shout(r)

	//? embedding 
	s:= starship{
		laser(3),
	}

	//> We can call the talk method directly from starship type, thanks to struct embedding::
	fmt.Println(s.talk())
	shout(s)

	getDates()
	printLocationMessage()
}
