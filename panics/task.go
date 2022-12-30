package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func checkString(str string) bool {
	return len(str) > 0
}

func main() {
	inFile, err := os.OpenFile("in.txt", os.O_RDONLY, os.ModeAppend)
	rowCounter := 1
	if err != nil {
		panic("Error opening an in file")
	}
	defer inFile.Close()
	outFile, err := os.OpenFile("out.txt", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		panic("Error opening an out file")
	}
	defer outFile.Close()
	defer func() {
		if err := recover(); err != nil {
			errorMessage := fmt.Errorf("%v", err)
			switch {
			case strings.Contains(errorMessage.Error(), "is not full"):
				fmt.Println(errorMessage.Error())
			default:
				panic("FATAL ERROR")
			}
			checkFile, fileErr := os.OpenFile("out.txt", os.O_RDONLY, os.ModeAppend)
			if fileErr != nil {
				panic("FATAL ERROR")
			}
			defer checkFile.Close()
			checkFileScanner := bufio.NewScanner(checkFile)
			for checkFileScanner.Scan() {
				fmt.Println(checkFileScanner.Text())
			}
		}
	}()

	scanner := bufio.NewScanner(inFile)
	writer := bufio.NewWriter(outFile)
	for scanner.Scan() {
		currentData := strings.Split(scanner.Text(), "|")
		currentRow := rowCounter
		currentName := currentData[0]
		currentAddress := currentData[1]
		currentCity := currentData[2]
		if !checkString(currentName) || !checkString(currentAddress) || !checkString(currentCity) {
			panic("Row number " + strconv.Itoa(currentRow) + " is not full")
		}
		validString := fmt.Sprintf("Row: %d\nName: %s\nAddress: %s\nCity: %s\n\n\n", currentRow, currentName, currentAddress, currentCity)
		writer.WriteString(validString)
		rowCounter++
		writer.Flush()
	}
}
