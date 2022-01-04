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

	o2_value := 0
	co2_value := 0

	o2_values := input
	co2_values := input

	for i := range o2_values[0] {
		o2_values, err := create_bit_slice(o2_values, i, true)
		if err != nil {
			return 0, err
		}
		if len(o2_values) == 1 {
			o2_value, _ = convert_binary_string_to_decimal(o2_values[0])
			break
		}
	}

	for i := range co2_values[0] {
		co2_values, err := create_bit_slice(co2_values, i, false)
		if err != nil {
			return 0, err
		}
		if len(co2_values) == 1 {
			co2_value, _ = convert_binary_string_to_decimal(co2_values[0])
			break
		}
	}

	return o2_value * co2_value, nil
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

// create slice containing most or least common values of bits at position n
// s slice of values, n bit position, b most common or least
func create_bit_slice(s []string, n int, b bool) ([]string, error) {
	var ret []string

	x := 0

	// find most common value at bit n
	for _, k := range s {
		y, err := strconv.Atoi(string(k[n]))
		if err != nil {
			return nil, err
		}
		x += y
	}

	// determine which values to keep based on o2 or co2 criteria
	if b {
		if x >= int(math.Ceil(float64(float64(len(s))/2))) {
			x = 1
		} else {
			x = 0
		}
	} else {
		if x >= int(math.Ceil(float64(float64(len(s))/2))) {
			x = 1
		} else {
			x = 0
		}
	}

	// append over determined criteria
	for _, k := range s {
		y, err := strconv.Atoi(string(k[n]))
		if err != nil {
			return nil, err
		}

		if b {
			if y == x {
				ret = append(ret, k)
			}
		} else {
			if y != x {
				ret = append(ret, k)
			}
		}
	}

	return ret, nil
}
