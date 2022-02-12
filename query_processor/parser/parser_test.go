package parser

import (
	"testing"
)

func TestJudgeStatementKind(t *testing.T) {
	rawQuery := "select 1"
	wantNextToken := &token{kind: numberToken, num: 1}
	wantStatementKind := Read
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
	if parsed.Block != wantParsedBlock {
		t.Errorf("failed parsed block\ngot: %s, want: %s", parsed.Block, wantParsedBlock)
	}
}

func TestParseSelectStatement(t *testing.T) {
	rawQuery := "select * from test_table;"
	wantParseds := []*parsed{
		{Block: "*"},
		{Block: "from"},
		{Block: "test_table"},
	}

	token, err := Tokenize(rawQuery)
	if err != nil {
		t.Fatal(err)
	}
	nextToken, _, _, err := judgeStatementKind(token)
	if err != nil {
		t.Error(err)
	}
	parsed, err := parseSelectStatement(nextToken)
	if err != nil {
		t.Error(err)
	}

	for i, p := range parsed {
		if *p != *wantParseds[i] {
			t.Errorf("failed parsed\ngot: %v, want: %v", p, wantParseds[i])
		}
	}
}
