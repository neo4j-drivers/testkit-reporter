package aggregation_test

import (
	"github.com/fbiville/testkit-reporter/pkg/aggregation"
	"github.com/fbiville/testkit-reporter/pkg/entity"
	"reflect"
	"testing"
)

func TestAggregation(st *testing.T) {
	test1 := entity.SkippedTest{Method: "method 1", Class: "class 1", Reason: "because reasons", FeatureFlags: []string{"f1", "f2"}}
	test2 := entity.SkippedTest{Method: "method 2", Class: "class 2", Reason: "because OTHER reasons", FeatureFlags: []string{"f2", "f3"}}
	test3 := entity.SkippedTest{Method: "method 3", Class: "class 3", Reason: "because reasons", FeatureFlags: []string{"f3"}}

	st.Run("groups by reason", func(t *testing.T) {
		aggregator := aggregation.New()
		aggregator.Aggregate(&test1)
		aggregator.Aggregate(&test2)
		aggregator.Aggregate(&test3)

		actual := aggregator.ByReason()

		expected := map[string][]entity.SkippedTest{
			"because reasons":       {test1, test3},
			"because OTHER reasons": {test2},
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

	st.Run("groups by features", func(t *testing.T) {
		aggregator := aggregation.New()
		aggregator.Aggregate(&test1)
		aggregator.Aggregate(&test2)
		aggregator.Aggregate(&test3)

		actual := aggregator.ByFeatureFlags()

		expected := map[string][]entity.SkippedTest{
			"f1": {test1},
			"f2": {test1, test2},
			"f3": {test2, test3},
		}
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("Expected %v, got %v", expected, actual)
		}
	})

}
