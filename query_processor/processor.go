package query_processor

import (
	"fmt"

	"github.com/ddddddO/database/query_processor/optimizer"
	"github.com/ddddddO/database/query_processor/parser"
)

func Process() {
	fmt.Println("Process!")

	parser.Tokenize()
	parser.Parse()

	optimizer.Optimize()
}
