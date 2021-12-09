package day_1

import (
	"log"
	"strconv"

	"github.com/Grantsuu/advent_of_code_2021/util"
)

func Day_1_solutions() error {
	log.Println("Day 1 Solution")

	input, err := util.ParseInputFile("day_1/day_1_input.txt")
	if err != nil {
		return err
	}

	res, err := part_1(input)
	if err != nil {
		return err
	}

	log.Println("Part 1 Answer:", res)

	res, err = part_2(input)
	if err != nil {
		return err
	}

	log.Println("Part 2 Answer:", res)

	return nil
}

func part_1(input []string) (int, error) {

	res := 0

	for i := 1; i < len(input); i++ {
		a, _ := strconv.Atoi(input[i])
		b, _ := strconv.Atoi(input[i-1])
		if a > b {
			res++
		}
	}

	return res, nil
}

func part_2(input []string) (int, error) {

	res := 0
	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])
	c, _ := strconv.Atoi(input[2])
	prev := a + b + c

	for i := 3; i < len(input); i++ {
		a, _ = strconv.Atoi(input[i])
		b, _ = strconv.Atoi(input[i-1])
		c, _ = strconv.Atoi(input[i-2])
		window := a + b + c

		if window > prev {
			res++
		}
		prev = window
	}

	return res, nil
}
