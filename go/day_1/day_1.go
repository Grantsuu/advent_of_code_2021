package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func day_1_sol() error {
	fmt.Println("Day 1 Solution")
	input, err := read("day_1_input.txt")
	if err != nil {
		return err
	}

	log.Println(input[1])
	return nil
}

func read(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func main() {
	err := day_1_sol()
	if err != nil {
		log.Println(err.Error())
	}
}
