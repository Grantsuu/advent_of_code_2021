package day_2

import (
	"fmt"
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

	res, err = part_2(input)
	if err != nil {
		return err
	}

	log.Println("Part 2 Answer:", res)

	return nil
}

func part_1(input []string) (int, error) {

	horizontal := 0
	depth := 0

	for _, k := range input {
		s := strings.Split(k, " ")
		if s[0] == "forward" {
			i, err := strconv.Atoi(s[1])
			if err != nil {
				return 0, err
			}
			horizontal += i
		} else if s[0] == "down" {
			i, err := strconv.Atoi(s[1])
			if err != nil {
				return 0, err
			}
			depth += i
		} else if s[0] == "up" {
			i, err := strconv.Atoi(s[1])
			if err != nil {
				return 0, err
			}
			depth -= i
		} else {
			return 0, fmt.Errorf("invalid instruction: %s", s[0])
		}
	}

	return horizontal * depth, nil
}

func part_2(input []string) (int, error) {
	horizontal := 0
	depth := 0
	aim := 0

	for _, k := range input {
		s := strings.Split(k, " ")
		if s[0] == "forward" {
			i, err := strconv.Atoi(s[1])
			if err != nil {
				return 0, err
			}
			horizontal += i
			depth += (aim * i)
		} else if s[0] == "down" {
			i, err := strconv.Atoi(s[1])
			if err != nil {
				return 0, err
			}
			aim += i
		} else if s[0] == "up" {
			i, err := strconv.Atoi(s[1])
			if err != nil {
				return 0, err
			}
			aim -= i
		} else {
			return 0, fmt.Errorf("invalid instruction: %s", s[0])
		}
	}

	return horizontal * depth, nil
}
