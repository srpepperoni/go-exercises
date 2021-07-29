package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func main() {
	ch := make(chan bool)
	go funcTimer(ch)
	quizGame(ch)
}

/*
* Main problem in part 1 of exercise 1 is compare strings, because when one of strings cames from standard input code has to verify \r and \n and due to every OS makes differents behaviours
* for end of lines you have to trim that characters in different ways.
 */

// Exercise #1: Quiz Game
func quizGame(fin chan bool) {
	recordFile, err := os.Open("./resources/problems.csv")
	count := 0

	if err != nil {
		fmt.Println("An error encountered ::", err)
	}
	reader := csv.NewReader(recordFile)
	records, _ := reader.ReadAll()

	inputReader := bufio.NewReader(os.Stdin)

registros:
	for _, record := range records {
		fmt.Println(record[0])

		select {
		case end, ok := <-fin:
			if ok {
				if end {
					fmt.Println("FIN")
					break registros
				}
			}
		default:
			answ, _ := inputReader.ReadString('\n')

			if runtime.GOOS == "windows" {
				answ = strings.TrimRight(answ, "\r\n")
			} else {
				answ = strings.TrimRight(answ, "\n")
			}

			if strings.Compare(answ, record[1]) == 0 {
				count++
			}
		}
	}

	fmt.Println("Correct Answers: " + strconv.Itoa(count) + "/" + strconv.Itoa(len(records)))
}

func funcTimer(fin chan bool) {
	time.Sleep(3 * time.Second)
	fin <- true
}
