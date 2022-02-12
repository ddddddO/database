package optimizer

import (
	"fmt"

	"github.com/ddddddO/database/query_processor/parser"
)

// NOTE: https://www.postgresql.jp/document/11/html/sql-explain.html
// の「以下は同じ問い合わせをJSON出力形式で出力したものです。」のところを参考に形だけつくる

// 実行計画struct
// https://www.postgresql.jp/document/11/html/using-explain.html
type plan struct {
	nodeType     nodeType
	relationName string // table name
	alias        string // table name
	startupCost  float32
	totalCost    float32
	planRows     uint
	planWidth    uint
}

type nodeType string

const (
	seqScan        = nodeType("seq_scan")
	indexScan      = nodeType("index_scan")
	bitmapHeapScan = nodeType("bitmap_heap_scan")
)

func Optimize(statement *parser.Statement) (*plan, error) {
	p := &plan{}
	fmt.Println("Optimize")
	// FIXME: 渡ってきたstatementの中のparsedにどんな種類のblockか(e.g. from句なのかwhere句なのかとか)のtypeを持たせた方がいいかも

	// TODO: https://oss-db.jp/dojo/dojo_info_10 によると、実行計画(plan)はランダムにサンプリングしたデータを元にした推定値であるそう。
	//       なので一回はテーブルにアクセスするということっぽい？

	return p, nil
}
