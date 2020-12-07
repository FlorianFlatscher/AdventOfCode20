package Input

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func ReadInputFile(day int) string {
	file, err := os.Open(fmt.Sprintf("src/Input/input%d.txt", day))
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}
