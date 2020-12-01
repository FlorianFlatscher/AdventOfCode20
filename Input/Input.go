package Input

import (
	"io/ioutil"
	"log"
	"os"
)

func ReadInputFile(day int) string {
	file, err := os.Open("./Input/input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}
