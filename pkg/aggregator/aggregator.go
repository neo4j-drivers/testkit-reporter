package aggregator

import (
	. "github.com/fbiville/testkit-reporter/pkg/entity"
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
	} else {
		a.byReason = aggregate(a.byReason, skippedTest.Reason, skippedTest)
	}
}

func aggregate(aMap map[string][]SkippedTest, key string, skipped *SkippedTest) map[string][]SkippedTest {
	_, exists := aMap[key]
	if !exists {
		aMap[key] = make([]SkippedTest, 0)
	}
	aMap[key] = append(aMap[key], *skipped)
	return aMap
}
