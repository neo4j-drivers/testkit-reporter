package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fbiville/testkit-reporter/pkg/aggregator"
	. "github.com/fbiville/testkit-reporter/pkg/entity"
	. "github.com/fbiville/testkit-reporter/pkg/formatter"
	"github.com/fbiville/testkit-reporter/pkg/parser"
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

	var aggregator = aggregator.New()
	var skippedTests []SkippedTest

	reader := bufio.NewReader(file)
	logParser := parser.NewLogParser()
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

	fmt.Println(FormatTableAsCountByKey("Skipped Tests by Feature Flags", "Feature", "Tests", aggregator.ByFeatureFlags()))
	fmt.Println(FormatTableAsCountByKey("Skipped Tests by Reason", "Reason", "Tests", aggregator.ByReason()))
}
