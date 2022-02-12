package optimizer

import (
	"fmt"

	"github.com/ddddddO/database/query_processor/parser"
)

// https://www.postgresql.jp/document/11/html/sql-explain.html
// の「以下は同じ問い合わせをJSON出力形式で出力したものです。」のところを参考に形だけつくる
func Optimize(statement *parser.Statement) {
	fmt.Println("Optimize")
}
