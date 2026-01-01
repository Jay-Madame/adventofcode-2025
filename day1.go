package main
import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

func goLeft(amountToTurn int, currentPosition int) int{
	for i := 0; i < amountToTurn; i++ {
		if currentPosition != 0 {
			currentPosition--
		} else {
			currentPosition = 99
		}
	}
	return currentPosition
}

func goRight(amountToTurn int, currentPosition int) int{
	for i := 0; i < amountToTurn; i++ {
		if currentPosition != 99 {
			currentPosition++
		} else {
			currentPosition = 0
		}
	}
	return currentPosition
}

func processLine(LineOfText string, currentPosition int) int{
	direction := LineOfText[:1]
	amountToTurn, _ := strconv.Atoi(LineOfText[1:])
	if direction == "L" {
		currentPosition = goLeft(amountToTurn, currentPosition)
	} else {
		currentPosition = goRight(amountToTurn, currentPosition)
	}
	return currentPosition
}

func main() {
	file, err := os.Open("input_day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var currentPosition int = 50
	var totalAmountOfZeros int = 0

	for scanner.Scan() {
		line := scanner.Text()
		currentPosition = processLine(line, currentPosition)
		if currentPosition == 0 {
			totalAmountOfZeros++
		}
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Password: %d\n", totalAmountOfZeros)

}