package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func closeWriteFileCon(counter int, totalBytes int) {
	fmt.Println("Was written " + strconv.Itoa(counter) + " rows")
	fmt.Println("Was written " + strconv.Itoa(totalBytes) + " bytes")
}

func main() {
	start := time.Now()
	var end time.Time
	defer func() {
		end = time.Now()
		elapsed := end.Sub(start)
		strExecTine := elapsed.String()
		fmt.Println("Execution time equals to: " + strExecTine)
	}()
	counter := 0
	totalBytes := 0
	inFile, err := os.OpenFile("in.txt", os.O_RDONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer inFile.Close()
	outFile, err := os.OpenFile("out.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		closeWriteFileCon(counter, totalBytes)
		outFile.Close()
	}()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		currentLine := strconv.Itoa(counter) + " " + scanner.Text() + "\n"
		outFile.WriteString(currentLine)
		strLen := len(currentLine)
		byteToInt := int(strLen)
		totalBytes += byteToInt
		counter++
	}
}
