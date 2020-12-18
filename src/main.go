package main

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/solution"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	days := []solution.IDay{
		solution.Day1{},
		solution.Day2{},
		solution.Day3{},
		solution.Day4{},
		solution.Day5{},
		solution.Day6{},
		solution.Day7{},
		solution.Day8{},
		solution.Day9{},
		solution.Day10{},
		solution.Day11{},
		solution.Day12{},
		solution.Day13{},
		solution.Day14{},
		solution.Day15{},
		solution.Day16{},
		solution.Day17{},
		solution.Day18{},
	}
	var output string

	printHeader()
	fmt.Print("Please input a day (1-25):")
	dayIndex, err := askForNumber()
	timeStart := time.Now()

	if err == nil && dayIndex > 0 && dayIndex <= len(days) {
		output = days[dayIndex-1].Calc()
	} else {
		log.Fatal("NO DAY FOUND?! I GUESS NO CHRISTMAS THIS TIME :(")
	}
	fmt.Println(output)
	fmt.Printf("(in %s)\n", time.Since(timeStart))
}

func printHeader() {
	path, _ := filepath.Abs("src/header.txt")
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	b, err := ioutil.ReadAll(file)
	fmt.Print(string(b), "\n")
}

func askForNumber() (int, error) {
	var input int
	for {
		_, err := fmt.Scanf("%d", &input)
		return input, err
	}
}
