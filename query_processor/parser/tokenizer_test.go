package parser

import (
	"testing"
)

func TestTokenize(t *testing.T) {
	rawQuery := "select 1"

	token, err := Tokenize(rawQuery)
	if err != nil {
		t.Error(err)
	}

	t.Log("token")
	for {
		switch token.kind {
		case numberToken:
			t.Log(token.num)
		case stringToken:
			t.Log(token.str)
		}

		if token.next == nil {
			break
		}
		token = token.next
	}
}
