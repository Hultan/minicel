package minicel

type CellType int

const (
	CellTypeNone CellType = iota
	CellTypeNumber
	CellTypeString
	CellTypeFormula
)

type Cell struct {
	Typ     CellType
	Number  int
	String  string
	Formula string
}
