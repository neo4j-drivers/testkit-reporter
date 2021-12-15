package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fbiville/testkit-reporter/pkg/aggregation"
	. "github.com/fbiville/testkit-reporter/pkg/entity"
	"github.com/fbiville/testkit-reporter/pkg/parsing"
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

	aggregator := aggregation.New()
	reader := bufio.NewReader(file)
	logParser := parsing.NewLogParser()

	var skippedTests []SkippedTest
	for {
		line, _, err := reader.ReadLine()
		if line != nil && err == nil {
			skippedTest := logParser.Parse(line)
			if skippedTest != nil {
				skippedTests = append(skippedTests, *skippedTest)
				aggregator.Aggregate(skippedTest)
			}
			continue
		}
		break
	}

	fmt.Println(aggregation.FormatTableByTestCount("Skipped Tests by Feature Flags", "Feature", "Tests", aggregator.ByFeatureFlags()))
	fmt.Println(aggregation.FormatTableByTestCount("Skipped Tests by Reason", "Reason", "Tests", aggregator.ByReason()))
}
