package parser

type token struct {
	kind kind
	num  int    // 数字
	str  string // 文字
	next *token // 次のトークン
}

type kind uint

// NOTE: 数字と文字だけではなく、例えば、スペースやセミコロンも一種の種類として定義した方が扱いやすいかも
const (
	numberToken kind = iota
	charToken
	spaceToken
	semicolonToken
)
