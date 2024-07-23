package main

import "fmt"

func main()  {
	poodle := Dog{"Poodle", 10, "Woof!"}
	fmt.Println(poodle)
	// {Poodle 10 Woof!}
	fmt.Printf("%+v\n", poodle)
	// {Breed:Poodle Weight:10 Sound:Woof!}
	fmt.Printf("Breed: %v\nWeight: %v\n", poodle.Breed, poodle.Weight)
	// Breed: Poodle
	// Weight: 10

	poodle.Speak()
	// Woof!
	poodle.Sound = "Arf!"
	poodle.Speak()
	// Arf!
	poodle.SpeakThreeTimes()
	// Arf! Arf! Arf!
	poodle.SpeakThreeTimes()
	// Arf! Arf! Arf!
}

type Dog struct {
	Breed string
	Weight int
	Sound string
}

// Speak is how the dog speaks
func (d Dog) Speak() {
	fmt.Println(d.Sound)
}

// SpeakThreeTimes is how the dog speaks loudly
func (d Dog) SpeakThreeTimes() {
	d.Sound = fmt.Sprintf("%v %v %v", d.Sound, d.Sound, d.Sound)
	fmt.Println(d.Sound)
}