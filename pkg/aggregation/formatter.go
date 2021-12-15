package aggregation

import (
	"fmt"

	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
	. "github.com/fbiville/testkit-reporter/pkg/entity"
)

func FormatTableByTestCount(title, column1, column2 string, data map[string][]SkippedTest) string {
	table := make([][]string, len(data))
	row := 0
	for k, v := range data {
		table[row] = []string{k, fmt.Sprintf("%d", len(v))}
		row++
	}

	formattedTable, err := markdown.NewTableFormatterBuilder().
		WithPrettyPrint().
		Build(column1, column2).
		Format(table)

	if err != nil {
		panic(err)
	}
	return title + ":\n\n" + formattedTable

}
