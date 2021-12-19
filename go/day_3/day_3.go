package day_3

import (
	"log"
	"math"
	"strconv"

	"github.com/Grantsuu/advent_of_code_2021/util"
)

func Day_3_solutions() error {
	log.Println("Day 3 Solution")

	input, err := util.ParseInputFile("day_3/day_3_input.txt")
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
	gammaBin := ""
	epsilonBin := ""

	gammaDec := 0
	epsilonDec := 0

	for i := 0; i < len(input[0]); i++ {
		n := 0
		for _, k := range input {
			x, _ := strconv.Atoi(string(k[i]))
			n += x
		}
		if n > (len(input) / 2) {
			gammaBin = gammaBin + "1"
			epsilonBin = epsilonBin + "0"
		} else {
			gammaBin = gammaBin + "0"
			epsilonBin = epsilonBin + "1"
		}
	}

	gammaDec, err := convert_binary_string_to_decimal(gammaBin)
	if err != nil {
		return 0, err
	}

	epsilonDec, err = convert_binary_string_to_decimal(epsilonBin)
	if err != nil {
		return 0, err
	}

	return gammaDec * epsilonDec, nil
}

func part_2(input []string) (int, error) {
	return 0, nil
}

func convert_binary_string_to_decimal(s string) (int, error) {
	res := 0

	for i := len(s) - 1; i >= 0; i-- {
		k, err := strconv.Atoi(string(s[i]))
		if err != nil {
			return 0, err
		}
		res += int(math.Pow(2, float64(len(s)-i-1)) * float64(k))
	}

	return res, nil
}
