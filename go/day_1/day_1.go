package day_1

import (
	"log"

	"github.com/Grantsuu/advent_of_code_2021/util"
)

func Day_1_solution() error {
	log.Println("Day 1 Solution")

	input, err := util.ParseInputFile("day_1/day_1_input.txt")
	if err != nil {
		return err
	}

	log.Println(input[1])
	return nil
}
