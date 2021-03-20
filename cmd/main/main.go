package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	quizGame()
}

// Exercise #1: Quiz Game
func quizGame() {
	recordFile, err := os.Open("./resources/problems.csv")
	var count, total int

	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	reader := csv.NewReader(recordFile)
	records, _ := reader.ReadAll()

	inputReader:= bufio.NewReader(os.Stdin)

	for _ , record := range records {
		fmt.Println(record[0])
		answ, _ := inputReader.ReadString('\n')

		answ = strings.TrimSuffix(answ, "\n")

		resp, _ := strconv.Atoi(answ)
		sol, _ := strconv.Atoi(record[1])

		if resp == sol {
			count++
		}

		total++
	}

	fmt.Println("Correct Answers: " + strconv.Itoa(count) + "/" + strconv.Itoa(total))
}



