package aggregation

import (
	"fmt"
	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
	. "github.com/neo4j-drivers/testkit-reporter/pkg/entity"
	"strconv"
)

func FormatTableByTestCount(title, column1, column2 string, data map[string][]SkippedTest) string {
	table := make([][]string, len(data))
	row := 0
	for k, v := range data {
		table[row] = []string{k, fmt.Sprintf("%d", len(v))}
		row++
	}
	formattedTable, err := markdown.NewTableFormatterBuilder().
		WithCustomSort(sortIntDesc(1), markdown.ASCENDING_ORDER.StringCompare(0)).
		WithPrettyPrint().
		Build(column1, column2).
		Format(table)

	if err != nil {
		panic(err)
	}
	return title + ":\n\n" + formattedTable

}

func sortIntDesc(column int) markdown.SortFunction {
	return markdown.SortFunction{
		Fn:     sortIntDescFn,
		Column: column,
	}
}

func sortIntDescFn(s1, s2 string) int {
	i1, _ := strconv.Atoi(s1)
	i2, _ := strconv.Atoi(s2)
	return i2 - i1
}
