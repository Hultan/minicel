package minicel

import (
	"fmt"
)

type CellType int

const (
	CellTypeNone CellType = iota
	CellTypeNumber
	CellTypeString
	CellTypeFormula
)

type Cell struct {
	Row     int
	Column  int
	Typ     CellType
	Number  int
	String  string
	Formula string
}

func (c *Cell) Value() string {
	switch c.Typ {
	case CellTypeNumber:
		return fmt.Sprintf("NUM(%s%d): %d", string(c.Row+64), c.Column, c.Number)
	case CellTypeFormula:
		return fmt.Sprintf("FRM(%s%d): %s", string(c.Row+64), c.Column, c.Formula)
	case CellTypeString:
		return fmt.Sprintf("STR(%s%d): %s", string(c.Row+64), c.Column, c.String)
	default:
		return "ERROR"
	}
}

func (c *Cell) Eval() string {
	switch c.Typ {
	case CellTypeNumber:
		return fmt.Sprintf("%d", c.Number)
	case CellTypeFormula:
		return fmt.Sprintf("FRM(%s)", c.Formula)
	case CellTypeString:
		return c.String
	default:
		return "ERROR"
	}
}
