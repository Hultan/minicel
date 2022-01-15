package minicel

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type Minicel struct {
}

func NewMinicel() *Minicel {
	return &Minicel{}
}

func (m *Minicel) Evaluate(fileName string) ([][]Cell, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	csv, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	cells, err := parseCells(csv)
	return cells, nil
}

func parseCells(csv [][]string) ([][]Cell, error) {
	var result [][]Cell

	for r, innerCsv := range csv {
		var cells []Cell
		for c, s := range innerCsv {
			var cell Cell = Cell{Row: r + 1, Column: c + 1}
			trimmedCell := strings.Trim(s, " \t\n)")
			i, err := strconv.Atoi(trimmedCell)
			if err == nil {
				cell.Typ = CellTypeNumber
				cell.Number = i
				cells = append(cells, cell)
				continue
			}
			if trimmedCell[0] == '=' {
				cell.Typ = CellTypeFormula
				cell.Formula = trimmedCell
				cells = append(cells, cell)
				continue
			}
			cell.Typ = CellTypeString
			cell.String = trimmedCell
			cells = append(cells, cell)
		}
		result = append(result, cells)
	}

	return result, nil
}
