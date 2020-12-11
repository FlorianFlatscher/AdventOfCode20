package main

import (
	"fmt"
	"github.com/FlorianFlatscher/AdventOfCode/src/days"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
)

func main() {
	days := []days.IDay{
		days.Day1{},
		days.Day2{},
		days.Day3{},
		days.Day4{},
		days.Day5{},
		days.Day6{},
		days.Day7{},
		days.Day8{},
		days.Day9{},
		days.Day10{},
		days.Day11{},
	}
	var output string

	printHeader()
	fmt.Print("Please input a day (1-24) to execute:")
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
