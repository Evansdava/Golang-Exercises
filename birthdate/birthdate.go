package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter date of birth (DD-MM-YYYY): ")
	dateString, _ := reader.ReadString('\n')

	dateSlice := strings.Split(dateString, "-")
	// day, _ := strconv.Atoi(dateSlice[0])
	// month, _ := strconv.Atoi(dateSlice[1])
	fmt.Printf("%T, %v\n", dateSlice[2], dateSlice[2])
	year, _ := strconv.Atoi(dateSlice[2])
	fmt.Printf("%T, %v\n", year, year)
	// fmt.Printf("%v %v %v\n", day, month, year)

	// v := "2000"
	// if s, err := strconv.Atoi(v); err == nil {
	// 	fmt.Printf("%T, %v", s, s)
	// }

	fmt.Print("Enter age: ")
	var age int
	fmt.Scanf("%d", &age)

	fmt.Println(age)
}
