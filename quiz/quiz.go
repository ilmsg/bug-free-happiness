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

const totalQuestions = 5

type Question struct {
	question string
	answer   string
}

func main() {
	filename, timeLimit := readArguments()
	f, err := os.Open(filename)
	if err != nil {
		return
	}

	questions, err := readCSV(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if questions == nil {
		return
	}

	score, err := askQuestion(questions, timeLimit)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Printf("Your Score %d/%d\n", score, totalQuestions)
}

func readArguments() (string, int) {
	filename := flag.String("filename", "problem.csv", "CSV file quiz question")
	timeLimit := flag.Int("limit", 30, "Time limit for each question")
	flag.Parse()
	return *filename, *timeLimit
}

func readCSV(f io.Reader) ([]Question, error) {
	allQuestions, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	numOfQuestions := len(allQuestions)
	if numOfQuestions == 0 {
		return nil, fmt.Errorf("No Question in file")
	}

	var data []Question
	for _, line := range allQuestions {
		ques := Question{}
		ques.question = line[0]
		ques.answer = line[1]
		data = append(data, ques)
	}

	return data, nil
}

func askQuestion(questions []Question, timeLimit int) (int, error) {
	totalScore := 0
	timer := time.NewTimer(time.Duration(timeLimit) * time.Second)
	done := make(chan string)

	go getInput(done)

	for i := range [totalQuestions]int{} {
		ans, err := eachQuestion(questions[i].question, questions[i].answer, timer.C, done)
		if err != nil && ans == -1 {
			return totalScore, nil
		}
		totalScore += ans
	}
	return totalScore, nil
}

func getInput(input chan string) {
	for {
		in := bufio.NewReader(os.Stdin)
		result, err := in.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input <- result
	}
}

func eachQuestion(question string, answer string, timer <-chan time.Time, done <-chan string) (int, error) {
	fmt.Printf("%s: ", question)
	for {
		select {
		case <-timer:
			return -1, fmt.Errorf("time out")
		case ans := <-done:
			score := 0
			if strings.Compare(strings.Trim(strings.ToLower(ans), "\n"), answer) == 0 {
				score = 1
			} else {
				return 0, fmt.Errorf("wrong answer")
			}
			return score, nil
		}
	}
}
