package main

import (
	"bufio"
	"fmt"
	"github.com/fbiville/testkit-reporter/pkg/aggregation"
	. "github.com/fbiville/testkit-reporter/pkg/entity"
	"github.com/fbiville/testkit-reporter/pkg/parsing"
	"os"
)

func main() {
	aggregator := aggregation.New()
	reader := bufio.NewReader(os.Stdin)
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
