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
	printHeader()
	fmt.Print("Please input a dayIndex (1-24) to execute:")
	dayIndex := askForNumber()
	timeStart := time.Now()

	days := []Days.Day1{Days.Day1{}}
	var output string
	if dayIndex > 0 && dayIndex <= len(days) {
		output = days[dayIndex-1].Calc()
	} else {
		log.Fatal("NO CODE FOUND?! I GUESS NO CHRISTMAS THIS TIME :(")
	}

	fmt.Println(output)
	fmt.Printf("(in %s)", time.Since(timeStart))
}

func printHeader() {
	file, err := os.Open("header.txt")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(file)
	fmt.Print(string(b), "\n")
}

func askForNumber() int {
	var input int
	for {
		_, err := fmt.Scanf("%d", &input)
		if err != nil {
			log.Fatal(err)
		} else {
			return input
		}
	}
}
