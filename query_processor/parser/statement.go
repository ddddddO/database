package parser

type Statement struct {
	Kind    statementKind
	Parseds []*parsed
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

type parsed struct {
	Block string
	Kind  blockKind
}

type blockKind uint

const (
// TODO: ここから
)
