package main
import (
	"fmt"
	"os"
	"log"
	"bufio"
	"strconv"
)

func goLeft(amountToTurn int, currentPosition int) (int, int){
	passedZeroCounter := 0
	for i := 0; i < amountToTurn; i++ {
		if currentPosition != 0 {
			currentPosition--
		} else {
			currentPosition = 99
			passedZeroCounter++
		}
	}
	return currentPosition, passedZeroCounter
}

func goRight(amountToTurn int, currentPosition int) (int, int){
	passedZeroCounter := 0
	for i := 0; i < amountToTurn; i++ {
		if currentPosition != 99 {
			currentPosition++
		} else {
			currentPosition = 0
			passedZeroCounter++
		}
	}
	return currentPosition, passedZeroCounter
}

func processLine(LineOfText string, currentPosition int) (int, int){
	direction := LineOfText[:1]
	amountToTurn, _ := strconv.Atoi(LineOfText[1:])
	if direction == "L" {
		return goLeft(amountToTurn, currentPosition)
	} else {
		return goRight(amountToTurn, currentPosition)
	}
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
	var zerosPassed int = 0

	for scanner.Scan() {
		line := scanner.Text()
		currentPosition, zerosPassed = processLine(line, currentPosition)
		if currentPosition == 0 {
			totalAmountOfZeros++
		}
		totalAmountOfZeros = totalAmountOfZeros+zerosPassed
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	fmt.Printf("Password: %d\n", totalAmountOfZeros)

}