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

func Parse(token *token) (*statement, error) {
	nextToken, statementKind, p, err := judgeStatementKind(token)
	if err != nil {
		return nil, err
	}

	statement := &statement{
		kind:    statementKind,
		parseds: []*parsed{p},
	}

	var parseds []*parsed
	switch statement.kind {
	case read:
		parseds, err = parseSelectStatement(nextToken)
	default:
		err = errors.New("not yet impl")
	}
	if err != nil {
		return nil, err
	}

	statement.parseds = append(statement.parseds, parseds...)

	return statement, nil
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

// NOTE: まずselectをパースする実装をする。他は後回しでいいかも
func parseSelectStatement(token *token) ([]*parsed, error) {
	parseds := []*parsed{}

	ret := ""
	for {
		// 文の終わり(=tokenがnil or セミコロン)まで進める
		if (token == nil) || (token.kind == semicolonToken) {
			parseds = append(parseds, &parsed{block: ret})
			break
		}
		// スペースが現れたら次のparsedのblockを作るためret初期化
		if token.kind == spaceToken {
			parseds = append(parseds, &parsed{block: ret})

			ret = ""
			token = token.next
			continue
		}

		// FIXME: 一旦数字は考えない
		ret += token.str
		token = token.next
	}

	return parseds, nil
}

// TODO:
func parseInsertStatement(token *token) ([]*parsed, error) {
	return []*parsed{{block: "not yet impl"}}, nil
}

// TODO:
func parseUpdateStatement(token *token) ([]*parsed, error) {
	return []*parsed{{block: "not yet impl"}}, nil
}

// TODO:
func parseDeleteStatement(token *token) ([]*parsed, error) {
	return []*parsed{{block: "not yet impl"}}, nil
}

// etc...
// NOTE: 予感だが、今後のためにcreate tableとinsert文のparseは実装した方がいいかも
