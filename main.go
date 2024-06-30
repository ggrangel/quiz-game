package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := "problems.csv"

	file, err := os.Open(fileName)
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
		answer := scanner.Text()
		if answer == record[1] {
			nCorrect++
		}
	}

	fmt.Printf("You scored %d out of %d.\n", nCorrect, nAsked)
}
