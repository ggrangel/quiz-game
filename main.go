package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	csvFilename := flag.String(
		"csv",
		"problems.csv",
		"a csv file in the format of 'question,answer' (default 'problems.csv')",
	)

	limitFlag := flag.Int(
		"limit", 30, "the time limit for the quiz in seconds (default 30)",
	)

	flag.Parse()

	file, err := os.Open(*csvFilename)
	if err != nil {
		log.Fatal(err)
	}
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	problems := parseLines(lines)
	nCorrect := 0

	timer := time.NewTimer(time.Duration(*limitFlag) * time.Second)
	go func() {
		<-timer.C
		exit(file, len(problems), nCorrect)
	}()

	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i, p.question)
		var answer string
		fmt.Scanf("%s", &answer) // Scanf already trims the spaces
		if answer == p.answer {
			nCorrect++
		}
	}

	exit(file, len(problems), nCorrect)
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			question: line[0],
			answer:   strings.TrimSpace(line[1]),
		}
	}
	return ret
}

type problem struct {
	question string
	answer   string
}

func exit(file *os.File, nAsked, Ncorrect int) {
	fmt.Printf("\nYou scored %d out of %d.\n", Ncorrect, nAsked)
	file.Close()
	os.Exit(0)
}
