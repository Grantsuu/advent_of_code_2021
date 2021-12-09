package main

import (
	"log"

	"github.com/Grantsuu/advent_of_code_2021/day_1"
)

func main() {
	log.Println("Advent of Code 2021")

	err := day_1.Day_1_solution()

	if err != nil {
		log.Println(err.Error())
	}
}
