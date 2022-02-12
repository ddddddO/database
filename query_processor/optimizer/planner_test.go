package optimizer

import (
	"testing"

	"github.com/ddddddO/database/model"
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

	wantPlan := &model.Plan{
		NodeType:     model.SeqScan, // 仮置き
		RelationName: "test_table",
		Alias:        "test_table", // 仮置き
		StartupCost:  0.00,         // 仮置き
		TotalCost:    439.99,       // 仮置き
		PlanRows:     20,           // 仮置き
		PlanWidth:    4,            // 仮置き
	}

	plan, err := buildPlan(statement)
	if err != nil {
		t.Fatal(err)
	}
	if *plan != *wantPlan {
		t.Errorf("failed plan\ngot: %v, want: %v", plan, wantPlan)
	}
}
