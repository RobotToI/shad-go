//go:build !solution

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: wordcount <file1> <file2> ...")
	}

	var fileNames []string = os.Args[1:]
	checkedFiles := fileChecker(fileNames)
	if len(checkedFiles) == 0 {
		log.Fatal("Correct file names doesn't provided")
	}

	var result []string
	for _, fileName := range checkedFiles {
		res, err := readInts(fileName)
		if err != nil {
			continue
		}
		result = append(result, res...)
	}
	countedInts := countInts(result)
	for k, v := range countedInts {
		if v > 1 {
			fmt.Printf("%d\t%s\n", v, k)
		}
	}
}

func fileChecker(fileNames []string) []string {
	correctFiles := []string{}
	for _, fileName := range fileNames {
		if _, err := os.Stat(fileName); err == nil {
			correctFiles = append(correctFiles, fileName)
		}
	}

	return correctFiles
}

func readInts(f string) ([]string, error) {
	file, err := os.Open(f)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", f, err)
	}
	defer file.Close()

	result := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		result = append(result, line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", f, err)
	}
	return result, nil
}

func countInts(intsSlice []string) map[string]int {
	counter := make(map[string]int)
	for _, number := range intsSlice {
		counter[number]++
	}

	return counter
}
