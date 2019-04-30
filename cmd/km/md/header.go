package md

import (
	"strings"

	"github.com/tealeg/xlsx"
)

type Align int

const (
	AlignLeft Align = iota
	AlignCenter
	AlignRight
)

func NewAlign(style string) Align {
	switch style {
	case "center":
		return AlignCenter
	case "right":
		return AlignRight
	default:
		return AlignLeft
	}
}

type HeaderCell struct {
	Title string
	Align Align
}

type Header []HeaderCell

// TODO: refactor
func (h Header) Dump() string {
	result := h.dumpTitleLine()
	result += "\n" + h.dumpAlignLine()

	return result
}

func (h Header) dumpTitleLine() string {
	var titles []string
	for _, c := range h {
		titles = append(titles, c.Title)
	}

	return "| " + strings.Join(titles, " | ") + " |"
}

func (h Header) dumpAlignLine() string {
	result := "|"
	for _, c := range h {
		switch c.Align {
		case AlignLeft:
			result += ":------ |"
		case AlignCenter:
			result += ":------:|"
		case AlignRight:
			result += " ------:|"
		}
	}
	return result
}

func GetHeader(sheet *xlsx.Sheet) (Header, int) {
	idx := findHeaderRow(sheet)
	if idx == -1 {
		return nil, -1
	}

	cells := sheet.Rows[idx].Cells
	header := make([]HeaderCell, 0, len(cells))
	for _, cell := range cells {
		if cell.Value == "" {
			break
		}
		align := cell.GetStyle().Alignment.Horizontal
		header = append(header, HeaderCell{Align: NewAlign(align), Title: cell.Value})
	}

	return header, idx
}
