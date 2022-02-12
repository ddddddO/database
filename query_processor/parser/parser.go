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
	nextToken, statementKind, p, err := judgeStatementKind(token)
	if err != nil {
		return err
	}

	statement := &statement{
		kind:    statementKind,
		parseds: []*parsed{p},
	}

	_ = statement
	_ = nextToken

	return nil
}

// 第1戻り値は読み進めたtokenの次のtoken
func judgeStatementKind(token *token) (*token, statementKind, *parsed, error) {
	if token == nil {
		return nil, undefined, nil, errors.New("nil token")
	}

	// 一旦tokenの先頭からスペースが出るまでまで読み進める
	// TODO: create table, drop tableなどは未対応
	ret := ""
	for {
		// tokenを最後まで進めた時
		if token == nil {
			break
		}
		// FIXME:
		if (token.kind != charToken) && (token.kind != spaceToken) && (token.kind != semicolonToken) {
			return nil, undefined, nil, errors.New("invalid token")
		}

		if token.kind == spaceToken {
			// スペースのtokenまで読み進める
			token = token.next
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
		return nil, undefined, nil, errors.New("invalid statement")
	}

	p := &parsed{
		block: ret,
	}
	return token, kind, p, nil
}
