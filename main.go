package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type SkippedTest struct {
	Method string
	Class  string
	Reason string

	FeatureFlags []string
}

func (s SkippedTest) String() string {
	return fmt.Sprintf("%s %s.%s: %s", s.FeatureFlags, s.Class, s.Method, s.Reason)
}

func main() {
	//url := "https://live.neo4j-build.io/downloadBuildLog.html?buildId=11949005"
	file, err := os.Open("build2.log")
	if err != nil {
		panic(err)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}()
	var skippedTests []SkippedTest
	reader := bufio.NewReader(file)
	var line []byte
	line, _, err = reader.ReadLine()
	for line != nil && err == nil {
		if skippedTest := matchLine(line); skippedTest != nil {
			skippedTests = append(skippedTests, *skippedTest)
		}
		line, _, err = reader.ReadLine()
	}

	fmt.Printf("%s", skippedTests)
}

func matchLine(line []byte) *SkippedTest {
	lineRegex := regexp.MustCompile(".*\\[.*] ([a-zA-Z0-9_-]*) \\(([a-zA-Z0-9_\\-.]*)\\) \\.\\.\\. skipped '(.*)'")
	match := lineRegex.FindSubmatch(line)
	if len(match) != 4 {
		return nil
	}
	reason := match[3]
	result := SkippedTest{
		Method: string(match[1]),
		Class:  string(match[2]),
		Reason: string(reason),
	}
	flagRegex := regexp.MustCompile("Feature\\.([A-Z0-9_\\-]*)")
	flags := flagRegex.FindAllSubmatch(reason, -1)
	var featureFlags []string
	for _, feature := range flags {
		featureFlags = append(featureFlags, string(feature[1]))
	}
	result.FeatureFlags = featureFlags
	return &result
}
