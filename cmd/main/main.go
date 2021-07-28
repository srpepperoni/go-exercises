package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	quizGame()
}

// Exercise #1: Quiz Game
func quizGame() {
	recordFile, err := os.Open("./resources/problems.csv")
	count := 0
	total := 0

	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	reader := csv.NewReader(recordFile)
	records, _ := reader.ReadAll()

	inputReader := bufio.NewReader(os.Stdin)

	for _, record := range records {
		fmt.Println(record[0])
		answ, _ := inputReader.ReadString('\n')

		if runtime.GOOS == "windows" {
			answ = strings.TrimRight(answ, "\r\n")
		} else {
			answ = strings.TrimRight(answ, "\n")
		}

		if strings.Compare(answ, record[1]) == 0 {
			count++
		}

		total++
	}

	fmt.Println("Correct Answers: " + strconv.Itoa(count) + "/" + strconv.Itoa(total))
}
