package main

import (
	"fmt"

	"github.com/ariebrainware/go-getting-started/models"
	"github.com/ariebrainware/go-getting-started/utils"
)

func main() {
	// var greetings = "Hello"
	name := "Arie"
	age := 21
	address := "Kavling Lama"
	gender := "male"
	friends := []string{"Jeremy", "Kevin", "Andrian"}

	person1 := models.Human{
		Name:    name,
		Age:     age,
		Address: address,
		Gender:  gender,
		Friends: friends,
	}

	if person1.Gender == "male" {
		fmt.Println("Hello Mr.", person1.Name)
	} else {
		fmt.Println("Hello Mrs.", person1.Name)
	}
	var tipe string
	switch person1.Age {
	case 21:
		tipe = "youngblood"
	case 31:
		tipe = "mature"
	case 41:
		tipe = "old"
	}
	fmt.Println(tipe)

	res := utils.Calculate(4, 4)
	fmt.Println(res)

	msg, err := utils.PingServer("192.168.1.1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(msg)
}
