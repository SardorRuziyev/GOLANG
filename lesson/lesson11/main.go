package main

import (
	"bufio"
	"fmt"
	"os"
)

func printLastNLines(lines []string, num int) []string {
	//var printLastNLines []string
	//for i := len(lines) - num; i < len(lines); i++ {
	//	printLastNLines = append(printLastNLines, lines[i])
	//}
	return lines[len(lines)-num:]
	//return printLastNLines
}

func main() {
	filepath := "data.txt"
	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// print the last 10 lines of the file
	printLastNLines := printLastNLines(lines, 5)
	for _, line := range printLastNLines {
		fmt.Println(line)
		fmt.Println("________")
	}
}
