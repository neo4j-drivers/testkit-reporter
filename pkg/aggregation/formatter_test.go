package aggregation_test

import (
	"github.com/fbiville/testkit-reporter/pkg/aggregation"
	"github.com/fbiville/testkit-reporter/pkg/entity"
	"testing"
)

func TestFormatting(st *testing.T) {
	input := map[string][]entity.SkippedTest{
		"skipped it": {
			{Method: "method 3", Class: "class", Reason: "skipped it as well"},
		},
		"skipped it as well": {
			{Method: "method 1", Class: "class", Reason: "skipped it"},
			{Method: "method 2", Class: "class", Reason: "skipped it"},
		},
	}
	expected := `skipped tests by reason:

| reason             | test count |
| ------------------ | ---------- |
| skipped it         | 1          |
| skipped it as well | 2          |
`

	st.Run("arranges data by test count", func(t *testing.T) {
		actual := aggregation.FormatTableByTestCount(
			"skipped tests by reason",
			"reason",
			"test count",
			input)

		if actual != expected {
			t.Errorf("Expected %q but got %q", expected, actual)
		}
	})
}
