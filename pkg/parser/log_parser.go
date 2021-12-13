package parser

import (
	"fmt"
	"regexp"

	. "github.com/fbiville/testkit-reporter/pkg/entity"
)

type LogParser interface {
	Parse(line []byte) *SkippedTest
}

func NewLogParser() LogParser {
	// [06:58:56]W:		 [Integration tests 3.5] test_summary_counters_case_2 (tests.neo4j.test_summary.TestSummary) ... skipped 'Needs support for Feature.TMP_FULL_SUMMARY, Feature.TMP_RESULT_KEYS'
	// 											 ^ method					   ^ class                                            ^ reason (including       ^ feature flags)
	methodRegex := `([a-zA-Z0-9_-]+)`
	classRegex := `([a-zA-Z0-9_\-.]+)`
	reasonRegex := `(.*)`
	return &teamcityLogParser{
		lineMatcher: regexp.MustCompile(fmt.Sprintf(
			`.*\[.*] %s \(%s\) \.\.\. skipped '%s'`,
			methodRegex, classRegex, reasonRegex,
		)),
		featureFlagRegex: regexp.MustCompile(`Feature\.([A-Z0-9_\-]+)`),
	}
}

type teamcityLogParser struct {
	lineMatcher      *regexp.Regexp
	featureFlagRegex *regexp.Regexp
}

func (t *teamcityLogParser) Parse(line []byte) *SkippedTest {
	match := t.lineMatcher.FindSubmatch(line)
	if len(match) != 4 {
		return nil
	}
	reason := match[3]
	result := &SkippedTest{
		Class:  string(match[2]),
		Method: string(match[1]),
		Reason: string(reason),
	}

	flags := t.featureFlagRegex.FindAllSubmatch(reason, -1)
	var featureFlags []string
	for _, feature := range flags {
		featureFlags = append(featureFlags, string(feature[1]))
	}
	result.FeatureFlags = featureFlags
	return result
}
