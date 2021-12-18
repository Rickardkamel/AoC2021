package main

import (
	"fmt"

	"github.com/rickardkamel/AoC2021/one"
)

func main() {
	result := one.CheckPreviousMesurement()
	fmt.Println("Total higher than previous: ", result)
}
