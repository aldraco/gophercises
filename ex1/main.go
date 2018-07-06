package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)


func main() {
	// Load problems from quiz file
	file, err := os.Open("problems.csv")
	if err != nil {
	    fmt.Println(err)
	}
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	if err != nil {
	    fmt.Println(err)
	}
	// Set up questioning
	correct := 0
	timerDuration := 15

	// Time the quiz in the background to exit gracefully 
	go func(correct *int, duration int) {
		timer := time.NewTimer(time.Duration(duration) * time.Second)
		<-timer.C
		finish(*correct)
	}(&correct, timerDuration)

	// Start asking questions
	for _, record := range lines {
		question, answer := parseProblem(record)
		response := ask(question)
		if answer == response {
			fmt.Println("Correct!")
			correct++
		}
	}
	fmt.Println("Congrats, you answered all the questions!")
	finish(correct)
}


func finish(correctAnswers int) {
	fmt.Printf("%d answered correctly\n", correctAnswers)
	os.Exit(0)
}

func ask(question string) (response string) {
    fmt.Println(question)
    fmt.Scanf("%s\n", &response)
    return response
}

func parseProblem(problem []string) (question, answer string) {
    return problem[0], problem[1]
}
