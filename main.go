package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	csvFlag := flag.String(
		"csv",
		"problems.csv",
		"a csv file in the format of 'question,answer' (default 'problems.csv')",
	)

	flag.Parse()

	file, err := os.Open(*csvFlag)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)

	nAsked, nCorrect := 0, 0

	scanner := bufio.NewScanner(os.Stdin)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		nAsked++
		fmt.Printf("Problem #%d: %s = ", nAsked, record[0])
		scanner.Scan()
		answer := strings.Trim(scanner.Text(), " ")
		result := strings.Trim(record[1], " ")

		if answer == result {
			nCorrect++
		} else {
			log.Println("answer:", answer)
			log.Println("response:", record[1])
		}
	}

	fmt.Printf("You scored %d out of %d.\n", nCorrect, nAsked)
}
