package md

import (
	"strings"

	"github.com/tealeg/xlsx"
)

type Table struct {
	Header
	Body []Row
}

// Dump outputs markdown contents with align info
func (b *Table) Dump() string {
	d := &strings.Builder{}
	d.WriteString(b.Header.Dump())
	d.WriteString("\n")
	for _, r := range b.Body {
		d.WriteString(r.Dump())
		d.WriteString("\n")
	}

	return d.String()
}

func Parse(s *xlsx.Sheet) *Table {
	h, idx := GetHeader(s)
	if idx == -1 {
		return nil
	}

	tbl := Table{Header: h}

	for k := idx + 1; k < len(s.Rows); k++ {
		r := s.Rows[k]
		if isBlank(r) {
			continue
		}
		tbl.Body = append(tbl.Body, convertRow(r, len(h)))
	}

	return &tbl
}
