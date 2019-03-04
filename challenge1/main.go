package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	score := 0
	count := 0

	var filename string
	flag.StringVar(&filename, "file", "problems.csv", "Filename of file containing quiz CSV")
	flag.Parse()

	file, err := os.Open(filename)
	check(err)

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		question := strings.Split(line, ",")[0]
		answer, _ := strconv.Atoi(strings.Split(line, ",")[1])

		print(question + ": ")

		var userAnswer int
		fmt.Scanln(&userAnswer)

		if userAnswer == answer {
			score++
		}
		count++

	}

	fmt.Printf("%d/%d\n", score, count)

}
