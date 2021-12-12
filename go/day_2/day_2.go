package day_2

import (
	"errors"
	"log"
	"strconv"
	"strings"

	"github.com/Grantsuu/advent_of_code_2021/util"
)

func Day_2_solutions() error {
	log.Println("Day 2 Solution")

	input, err := util.ParseInputFile("day_2/day_2_input.txt")
	if err != nil {
		return err
	}

	res, err := part_1(input)
	if err != nil {
		return err
	}

	log.Println("Part 1 Answer:", res)

	return nil
}

func part_1(input []string) (int, error) {

	horizontal := 0
	depth := 0

	for _, k := range input {
		s := strings.Split(k, " ")
		if s[0] == "forward" {
			i, _ := strconv.Atoi(s[1])
			horizontal += i
		} else if s[0] == "down" {
			i, _ := strconv.Atoi(s[1])
			depth += i
		} else if s[0] == "up" {
			i, _ := strconv.Atoi(s[1])
			depth -= i
		} else {
			return 0, errors.New("invalid instruction")
		}
	}

	log.Println("horizontal", horizontal, "depth", depth)

	return horizontal * depth, nil
}
