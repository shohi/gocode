package md

import (
	"strings"

	"github.com/tealeg/xlsx"
)

type Cell string
type Row []Cell

// TODO: refactor
func (r Row) Dump() string {
	result := "|"
	for _, c := range r {
		result += string(c) + " | "
	}

	return result
}

func hasMerge(r *xlsx.Row) bool {
	for _, cell := range r.Cells {
		if cell.HMerge > 0 || cell.VMerge > 0 {
			return true
		}
	}
	return false
}

func findHeaderRow(s *xlsx.Sheet) int {
	found := -1
	for i, row := range s.Rows {
		if !hasMerge(row) {
			found = i
			break
		}
	}

	return found
}

func isBlank(r *xlsx.Row) bool {
	if r == nil || len(r.Cells) == 0 {
		return true
	}

	for _, cell := range r.Cells {
		if strings.TrimSpace(cell.Value) != "" {
			return false
		}
	}

	return true
}

func convertRow(r *xlsx.Row, limit int) Row {
	row := make(Row, 0, limit)

	min := min(len(r.Cells), limit)
	for j := 0; j < min; j++ {
		row = append(row, Cell(r.Cells[j].Value))
	}

	for j := min; j < limit; j++ {
		row = append(row, "")
	}

	return row
}

func min(i, j int) int {
	if i < j {
		return i
	}

	return j
}
