package aggregation

import (
	. "github.com/neo4j-drivers/testkit-reporter/pkg/entity"
)

type aggregator struct {
	byFeatureFlags map[string][]SkippedTest
	byReason       map[string][]SkippedTest
}

func New() *aggregator {
	return &aggregator{
		byFeatureFlags: make(map[string][]SkippedTest),
		byReason:       make(map[string][]SkippedTest),
	}
}

func (a *aggregator) ByReason() map[string][]SkippedTest {
	return a.byReason
}

func (a *aggregator) ByFeatureFlags() map[string][]SkippedTest {
	return a.byFeatureFlags
}

func (a *aggregator) Aggregate(skippedTest *SkippedTest) {
	if skippedTest.FeatureFlags != nil {
		for _, feature := range skippedTest.FeatureFlags {
			a.byFeatureFlags = aggregate(a.byFeatureFlags, feature, skippedTest)
		}
	}
	a.byReason = aggregate(a.byReason, skippedTest.Reason, skippedTest)
}

func aggregate(data map[string][]SkippedTest, k string, v *SkippedTest) map[string][]SkippedTest {
	if _, exists := data[k]; !exists {
		data[k] = make([]SkippedTest, 0)
	}
	data[k] = append(data[k], *v)
	return data
}
