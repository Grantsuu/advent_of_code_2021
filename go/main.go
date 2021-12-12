package main

import (
	"log"

	"github.com/Grantsuu/advent_of_code_2021/day_1"
	"github.com/Grantsuu/advent_of_code_2021/day_2"
)

func main() {
	log.Println("Advent of Code 2021")

	err := day_1.Day_1_solutions()
	if err != nil {
		log.Println(err.Error())
	}

	err = day_2.Day_2_solutions()
	if err != nil {
		log.Println(err.Error())
	}
}
