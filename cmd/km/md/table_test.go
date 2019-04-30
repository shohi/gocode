package md

import (
	"log"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tealeg/xlsx"
)

func TestTable(t *testing.T) {
	assert := assert.New(t)

	fp := "testdata/sample.xlsx"
	log.Printf("base ==> %v, ext ==> %v ", filepath.Base(fp), filepath.Ext(fp))

	xlFile, err := xlsx.OpenFile(fp)
	assert.Nil(err)

	assert.True(len(xlFile.Sheets) >= 1)

	sheet := xlFile.Sheets[0]
	log.Printf("Sheet name: %s\n", sheet.Name)

	h, idx := GetHeader(sheet)
	log.Printf("index: %v, headers: %+v\n", idx, h)

	tbl := Parse(sheet)

	log.Printf("table:\n %v", tbl.Dump())

}
