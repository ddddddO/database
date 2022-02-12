package parser

import (
	"testing"
)

func TestJudgeStatementKind(t *testing.T) {
	rawQuery := "select 1"
	wantNextToken := &token{kind: numberToken, num: 1}
	wantStatementKind := read
	wantParsedBlock := "select"

	token, err := Tokenize(rawQuery)
	if err != nil {
		t.Fatal(err)
	}

	nextToken, statementKind, parsed, err := judgeStatementKind(token)
	if err != nil {
		t.Error(err)
	}
	if *nextToken != *wantNextToken {
		t.Errorf("failed next token\ngot: %v, want: %v", nextToken, wantNextToken)
	}
	if statementKind != wantStatementKind {
		t.Errorf("failed statement kind\ngot: %d, want: %d", statementKind, wantStatementKind)
	}
	if parsed.block != wantParsedBlock {
		t.Errorf("failed parsed block\ngot: %s, want: %s", parsed.block, wantParsedBlock)
	}
}
