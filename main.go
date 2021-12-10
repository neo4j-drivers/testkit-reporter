package main

import (
	"bufio"
	"fmt"
	"github.com/fbiville/testkit-reporter/pkg/parser"
	"os"
)

func main() {
	file, err := os.Open("build.log")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()

	var skippedTests []parser.SkippedTest
	reader := bufio.NewReader(file)
	logParser := parser.NewLogParser()
	for {
		line, _, err := reader.ReadLine()
		if line != nil && err == nil {
			skippedTest := logParser.Parse(line)
			if skippedTest != nil {
				skippedTests = append(skippedTests, *skippedTest)
			}
			continue
		}
		break
	}
	fmt.Printf("%s", skippedTests)
}
