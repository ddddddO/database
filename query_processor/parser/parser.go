package parser

import (
	"strings"

	"github.com/pkg/errors"
)

// NOTE: 完全に雰囲気で実装中
func Parse(token *token) (*Statement, error) {
	nextToken, statementKind, p, err := judgeStatementKind(token)
	if err != nil {
		return nil, err
	}

	statement := &Statement{
		Kind:    statementKind,
		Parseds: []*Parsed{p},
	}

	var parseds []*Parsed
	switch statement.Kind {
	case Read:
		parseds, err = parseSelectStatement(nextToken)
	default:
		err = errors.New("not yet impl")
	}
	if err != nil {
		return nil, err
	}

	statement.Parseds = append(statement.Parseds, parseds...)

	return statement, nil
}

// 第1戻り値は読み進めたtokenの次のtoken
func judgeStatementKind(token *token) (*token, statementKind, *Parsed, error) {
	if token == nil {
		return nil, Undefined, nil, errors.New("nil token")
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
			return nil, Undefined, nil, errors.New("invalid token")
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
		kind = Create
	case "select":
		kind = Read
	case "update":
		kind = Update
	case "delete":
		kind = Delete
	default:
		return nil, Undefined, nil, errors.New("invalid statement")
	}

	p := &Parsed{
		Block: ret,
		Kind:  CommandKind,
	}
	return token, kind, p, nil
}

// NOTE: まずselectをパースする実装をする。他は後回しでいいかも
func parseSelectStatement(token *token) ([]*Parsed, error) {
	parseds := []*Parsed{}

	ret := ""
	for {
		// 文の終わり(=tokenがnil or セミコロン)まで進める
		if (token == nil) || (token.kind == semicolonToken) {
			parseds = append(parseds, &Parsed{Block: ret})
			break
		}
		// スペースが現れたら次のparsedのblockを作るためret初期化
		if token.kind == spaceToken {
			parseds = append(parseds, &Parsed{Block: ret})

			ret = ""
			token = token.next
			continue
		}

		// FIXME: 一旦数字は考えない
		ret += token.str
		token = token.next
	}

	prev := parseds[0]
	// FIXME: bugをはらんでるはず
	for i, p := range parseds {
		isColumn := (i == 0 || prev.Kind == ColumnKind) && p.Block != "from"
		isFromPhrase := p.Block == "from" || prev.Kind == FromPhrase && p.Block != "where"
		isWherePhrase := (prev.Kind != ColumnKind) && (prev.Kind != FromPhrase)

		if isColumn {
			p.Kind = ColumnKind
			prev = p
			continue
		}
		if isFromPhrase {
			p.Kind = FromPhrase
			prev = p
			continue
		}
		if isWherePhrase {
			p.Kind = WherePhrase
			prev = p
			continue
		}
	}

	return parseds, nil
}

// TODO:
func parseInsertStatement(token *token) ([]*Parsed, error) {
	return []*Parsed{{Block: "not yet impl"}}, nil
}

// TODO:
func parseUpdateStatement(token *token) ([]*Parsed, error) {
	return []*Parsed{{Block: "not yet impl"}}, nil
}

// TODO:
func parseDeleteStatement(token *token) ([]*Parsed, error) {
	return []*Parsed{{Block: "not yet impl"}}, nil
}

// etc...
// NOTE: 予感だが、今後のためにcreate tableとinsert文のparseは実装した方がいいかも
