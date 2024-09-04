package main

import (
	"fmt"

	"github.com/GoesToEleven/dog"
	"github.com/GoesToEleven/puppy"
)

func main() {

	puppy.From12()
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := dog.WhenGrownUp("uau")

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println(puppy.BigBark())
	fmt.Println(puppy.BigBarks())

	//fmt.Println(dog.Latir("uau"))

	puppy.From13()
}
