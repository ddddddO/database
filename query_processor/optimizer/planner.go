package optimizer

import (
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
	// TODO: https://oss-db.jp/dojo/dojo_info_10 によると、実行計画(plan)はランダムにサンプリングしたデータを元にした推定値であるそう。
	//       なので一回はテーブルにアクセスするということっぽい？
	// FIXME:だとすると、ストレージエンジンの実装をしなくてはならなそう。一旦、簡易的に実行計画structに値を入れて返してその先にいったあとでここに戻ってくる

	plan, err := buildPlan(statement)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func buildPlan(statement *parser.Statement) (*plan, error) {
	r := getRelation(statement.Parseds)

	p := &plan{
		nodeType:     seqScan, // 仮置き
		relationName: r,       // NOTE: 現状、statementからここしかわからない。。
		alias:        r,       // 仮置き
		startupCost:  0.00,    // 仮置き
		totalCost:    439.99,  // 仮置き
		planRows:     20,      // 仮置き
		planWidth:    4,       // 仮置き
	}

	return p, nil
}

func getRelation(parseds []*parser.Parsed) string {
	var fromPhraseElements []*parser.Parsed
	for _, elm := range parseds {
		if elm.Kind == parser.FromPhrase {
			fromPhraseElements = append(fromPhraseElements, elm)
		}
	}

	if len(fromPhraseElements) >= 2 {
		return fromPhraseElements[1].Block
	}
	return "" // こない想定
}
