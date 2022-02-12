package parser

type Statement struct {
	Kind    statementKind
	Parseds []*Parsed
}

type statementKind uint

const (
	Undefined statementKind = iota
	Create
	Read
	Update
	Delete
	CreateTable
	DropTable
	AlterTable
	// etc...
)

type Parsed struct {
	Block string
	Kind  blockKind
}

type blockKind uint

const (
	CommandKind blockKind = iota // FIXME: "select"とか"insert"などはこれ。命名直す
	ColumnKind
	FromPhrase // from句を構成するblock
	WherePhrase
)
