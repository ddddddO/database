package parser

type token struct {
	kind kind
	num  int    // 数字
	str  string // 文字
	next *token // 次のトークン
}

type kind uint

const (
	numberToken kind = iota
	stringToken
)
