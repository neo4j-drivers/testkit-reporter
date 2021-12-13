package formatter

import (
	"fmt"

	"github.com/fbiville/markdown-table-formatter/pkg/markdown"
	. "github.com/fbiville/testkit-reporter/pkg/entity"
)

func FormatTableAsCountByKey(title string, column1Title string, column2Title string, m map[string][]SkippedTest) string {
	var table [][]string = make([][]string, len(m))
	var i int = 0
	for k, v := range m {
		table[i] = []string{k, fmt.Sprintf("%d", len(v))}
		i++
	}

	formattedTable, err := markdown.NewTableFormatterBuilder().
		WithPrettyPrint().
		Build(column1Title, column2Title).
		Format(table)
	if err != nil {
		panic(err)
	}
	return title + ":\n\n" + formattedTable

}
