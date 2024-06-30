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
	"time"
)

func main() {
	csvFlag := flag.String(
		"csv",
		"problems.csv",
		"a csv file in the format of 'question,answer' (default 'problems.csv')",
	)

	limitFlag := flag.Int(
		"limit", 30, "the time limit for the quiz in seconds (default 30)",
	)

	flag.Parse()

	file, err := os.Open(*csvFlag)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)

	scanner := bufio.NewScanner(os.Stdin)

	nAsked, nCorrect := 0, 0

	timer := time.NewTimer(time.Duration(*limitFlag) * time.Second)

	go func() {
		<-timer.C
		exit(file, nAsked, nCorrect)
	}()

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
		}
	}

	exit(file, nAsked, nCorrect)
}

func exit(file *os.File, nAsked, Ncorrect int) {
	fmt.Printf("\nYou scored %d out of %d.\n", Ncorrect, nAsked)
	file.Close()
	os.Exit(0)
}
