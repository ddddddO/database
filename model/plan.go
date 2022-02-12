package model

// 実行計画struct
// https://www.postgresql.jp/document/11/html/using-explain.html
type Plan struct {
	NodeType     nodeType
	RelationName string // table name
	Alias        string // table name
	StartupCost  float32
	TotalCost    float32
	PlanRows     uint
	PlanWidth    uint
}

type nodeType string

const (
	SeqScan        = nodeType("seq_scan")
	IndexScan      = nodeType("index_scan")
	BitmapHeapScan = nodeType("bitmap_heap_scan")
)
