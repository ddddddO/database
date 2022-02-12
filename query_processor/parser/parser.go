package parser

import (
	"strings"

	"github.com/pkg/errors"
)

// NOTE: 完全に雰囲気で実装中
type statement struct {
	kind    statementKind
	parseds []*parsed
}

type statementKind uint

const (
	undefined statementKind = iota
	create
	read
	update
	delete
	createTable
	dropTable
	alterTable
	// etc...
)

type parsed struct {
	block string
}

// NOTE: この関数でやりたいことは何か
func Parse(token *token) error {
	i, kind, parsed, err := judgeStatementKind(token)
	if err != nil {
		return err
	}

	_ = i
	_ = kind
	_ = parsed

	return nil
}

// 第1戻り値はtokenをどこまで読み進めたかを表わす数
// TODO: 読み進めたところまでのtokenを返した方が使う側は便利だと思った
func judgeStatementKind(token *token) (int, statementKind, *parsed, error) {
	if token == nil {
		return 0, undefined, nil, errors.New("nil token")
	}

	// 一旦tokenの先頭からスペースが出るまでまで読み進める
	// TODO: create table, drop tableなどは未対応
	ret := ""
	cnt := 0
	for {
		// tokenを最後まで進めた時
		if token == nil {
			break
		}
		// FIXME:
		if (token.kind != charToken) && (token.kind != spaceToken) && (token.kind != semicolonToken) {
			return 0, undefined, nil, errors.New("invalid token")
		}
		cnt++

		if token.kind == spaceToken {
			break
		}

		ret += token.str
		token = token.next
	}

	var kind statementKind
	ret = strings.ToLower(ret)
	switch ret {
	case "insert":
		kind = create
	case "select":
		kind = read
	case "update":
		kind = update
	case "delete":
		kind = delete
	default:
		return 0, undefined, nil, errors.New("invalid statement")
	}

	p := &parsed{
		block: ret,
	}
	return cnt, kind, p, nil
}
