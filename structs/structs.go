package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type house struct {
	address string
	city    string
	rooms   int
	price   int
}

func main() {
	var houses []house
	for i := 1; i != 0; {
		fmt.Printf("Enter house %d: \n", i)
		houses = append(houses, enterHouse())

		fmt.Print("Enter another? (y/n) ")
		var cont string
		fmt.Scanf("%s", &cont)

		if cont != "y" {
			i = 0
		} else {
			i += 1
		}
	}

	for i, h := range houses {
		fmt.Printf("%d: %-20v\t%-15v\t%-2d Rooms\t$%d\n", i+1, h.address, h.city, h.rooms, h.price)
	}
}

func enterHouse() house {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter address: ")
	adr, _ := reader.ReadString('\n')

	fmt.Print("Enter city: ")
	city, _ := reader.ReadString('\n')

	fmt.Print("Enter number of rooms: ")
	var rooms int
	fmt.Scanf("%d", &rooms)

	fmt.Print("Enter price: ")
	var price int
	fmt.Scanf("%d", &price)

	return house{address: strings.TrimSpace(adr), city: strings.TrimSpace(city), rooms: rooms, price: price}
}
