package optimizer

import (
	"testing"

	p "github.com/ddddddO/database/query_processor/parser"
)

func TestBuildPlan(t *testing.T) {
	statement := &p.Statement{
		Kind: p.Read,
		Parseds: []*p.Parsed{
			{Block: "select", Kind: p.CommandKind},
			{Block: "*", Kind: p.ColumnKind},
			{Block: "from", Kind: p.FromPhrase},
			{Block: "test_table", Kind: p.FromPhrase},
		},
	}

	wantPlan := &plan{
		nodeType:     seqScan, // 仮置き
		relationName: "test_table",
		alias:        "test_table", // 仮置き
		startupCost:  0.00,         // 仮置き
		totalCost:    439.99,       // 仮置き
		planRows:     20,           // 仮置き
		planWidth:    4,            // 仮置き
	}

	plan, err := buildPlan(statement)
	if err != nil {
		t.Fatal(err)
	}
	if *plan != *wantPlan {
		t.Errorf("failed plan\ngot: %v, want: %v", plan, wantPlan)
	}
}
