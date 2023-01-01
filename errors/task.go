package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func main() {
	counter := 0
	defer func() {
		if err := recover(); err != nil {
			errMessage := fmt.Errorf("Error: %v", err)
			fmt.Println(errMessage.Error())
			fmt.Printf("Total strings: %d\n", counter)
		}
	}()
	inFile, err := os.OpenFile("in.txt", os.O_RDONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	defer func() {
		inFile.Close()
	}()
	scanner := bufio.NewScanner(inFile)
	for {
		if flag := scanner.Scan(); !flag {
			err := errors.New("End of the file")
			panic(err)
		}
		counter++
	}
}
