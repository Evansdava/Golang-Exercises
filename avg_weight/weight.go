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
	weightString, _ := reader.ReadString('\n')
	weights := strings.Split(weightString, " ")
	avg := mean(weights)
	fmt.Print(avg)
}

func mean(nums []string) float64 {
	var total float64 = 0.0
	for _, v := range nums {
		float, _ := strconv.ParseFloat(v, 32)
		total += float
	}
	return total / float64(len(nums))
}
