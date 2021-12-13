package parser_test

import (
	"reflect"
	"strings"
	"testing"

	"github.com/fbiville/testkit-reporter/pkg/entity"
	"github.com/fbiville/testkit-reporter/pkg/parser"
)

const simpleTeamcityLog = `
[06:58:56]W:		 [Integration tests 3.5] test_summary_counters_case_2 (tests.neo4j.test_summary.TestSummary) ... skipped 'YOLO'
`
const teamcityLogWithFeatures = `
[06:58:56]W:		 [Integration tests 3.5] test_summary_counters_case_2 (tests.neo4j.test_summary.TestSummary) ... skipped 'Needs support for Feature.TMP_FULL_SUMMARY, Feature.TMP_RESULT_KEYS'
`
const expectedClass = "tests.neo4j.test_summary.TestSummary"
const expectedMethod = "test_summary_counters_case_2"

func TestLogParser(st *testing.T) {
	logParser := parser.NewLogParser()

	st.Run("parses simple Teamcity log line", func(t *testing.T) {
		expected := &entity.SkippedTest{
			Class:  expectedClass,
			Method: expectedMethod,
			Reason: "YOLO",
		}

		actual := logParser.Parse([]byte(strings.Trim(simpleTeamcityLog, "\n")))

		AssertSkippedTest(t, actual, expected)
	})

	st.Run("parses Teamcity log line with feature flags", func(t *testing.T) {
		expected := &entity.SkippedTest{
			Class:        expectedClass,
			Method:       expectedMethod,
			Reason:       "Needs support for Feature.TMP_FULL_SUMMARY, Feature.TMP_RESULT_KEYS",
			FeatureFlags: []string{"TMP_FULL_SUMMARY", "TMP_RESULT_KEYS"},
		}

		actual := logParser.Parse([]byte(strings.Trim(teamcityLogWithFeatures, "\n")))

		AssertSkippedTest(t, actual, expected)
	})
}

func AssertSkippedTest(t *testing.T, expected *entity.SkippedTest, actual *entity.SkippedTest) {
	if expected.Class != actual.Class {
		t.Errorf("Expected class %q, got %q", expected.Class, actual.Class)
	}
	if expected.Method != actual.Method {
		t.Errorf("Expected method %q, got %q", expected.Method, actual.Method)
	}
	if expected.Reason != actual.Reason {
		t.Errorf("Expected reason %q, got %q", expected.Reason, actual.Reason)
	}
	if !reflect.DeepEqual(expected.FeatureFlags, actual.FeatureFlags) {
		t.Errorf("Expected feature flags %s, got %s", expected.FeatureFlags, actual.FeatureFlags)
	}
}
