package parser

import (
	"testing"
)

func TestJudgeStatementKind(t *testing.T) {
	rawQuery := "select 1"
	wantCnt := 7
	wantKind := read
	wantParsedBlock := "select"

	token, err := Tokenize(rawQuery)
	if err != nil {
		t.Fatal(err)
	}

	i, kind, parsed, err := judgeStatementKind(token)
	if err != nil {
		t.Error(err)
	}
	if i != wantCnt {
		t.Errorf("failed cnt\ngot: %d, want: %d", i, wantCnt)
	}
	if kind != wantKind {
		t.Errorf("failed kind\ngot: %d, want: %d", kind, wantKind)
	}
	if parsed.block != wantParsedBlock {
		t.Errorf("failed parsed block\ngot: %s, want: %s", parsed.block, wantParsedBlock)
	}
}
