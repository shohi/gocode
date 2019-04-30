package md

import (
	"os"
	"path"
	"strings"

	"github.com/tealeg/xlsx"
)

func XlsxToMd(xlsxPath string) string {
	xlFile, err := xlsx.OpenFile(xlsxPath)
	if err != nil {
		panic(err)
	}
	if len(xlFile.Sheets) == 0 {
		panic("no sheets")
	}

	tbl := Parse(xlFile.Sheets[0])
	mdFilePath := FilenameWithoutExtension(xlsxPath) + ".md"

	file, err := os.Create(mdFilePath)
	if err != nil {
		panic(err)
	}

	file.WriteString(tbl.Dump())
	file.Close()

	return mdFilePath
}

func FilenameWithoutExtension(filename string) string {
	return strings.TrimSuffix(filename, path.Ext(filename))
}
