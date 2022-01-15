package minicel

import (
	"encoding/csv"
	"os"
)

type Minicel struct {
}

func NewMinicel() *Minicel {
	return &Minicel{}
}

func (m *Minicel) Evaluate(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
