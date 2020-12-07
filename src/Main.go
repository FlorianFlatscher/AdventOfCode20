package main

import (
	"./Days"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

func main() {
	days := []Days.IDay{Days.Day1{}, Days.Day2{}, Days.Day3{}, Days.Day4{}, Days.Day5{}, Days.Day6{}, Days.Day7{}}
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
	fmt.Printf("(in %s)", time.Since(timeStart))
}

func printHeader() {
	file, err := os.Open("src/header.txt")
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
