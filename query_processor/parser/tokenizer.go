package parser

import (
	"strconv"
)

// TODO: https://notes.eatonphil.com/database-basics.html を参考に実装をする（したい、ゆくゆくは）
// 以下を参考にした
// https://www.sigbus.info/compilerbook#%E3%82%B9%E3%83%86%E3%83%83%E3%83%973%E3%83%88%E3%83%BC%E3%82%AF%E3%83%8A%E3%82%A4%E3%82%B6%E3%82%92%E5%B0%8E%E5%85%A5
func Tokenize(rawQuery string) (*token, error) {
	head := &token{}
	var prev *token
	for i := 0; i < len(rawQuery); i++ {
		if i == 0 {
			prev = head
		}

		atom := string(rawQuery[i])
		current := &token{}

		// 数字の場合
		num, err := strconv.Atoi(atom)
		if err == nil {
			current.kind = numberToken
			current.num = num
			prev.next = current

			prev = current
			continue
		}

		// 文字の場合
		kind := charToken
		switch atom {
		case " ":
			kind = spaceToken
		case ";":
			kind = semicolonToken
		}
		current.kind = kind
		current.str = atom
		prev.next = current

		prev = current
	}

	return head.next, nil
}
