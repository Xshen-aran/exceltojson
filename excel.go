package exceltojson

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func parseFile(filepath string) {
	// open file
	// parse file
	// return json
	xlsx, err := excelize.OpenFile(filepath)
	if err != nil {
		panic(err.Error())
	}
	sheets := xlsx.GetSheetList()
	for _, sheet := range sheets {
		rows, err := xlsx.GetRows(sheet)
		if err != nil {
			return
		}
		key := rows[0]
		fmt.Println(key)
		for _, row := range rows[1:] {
			for _, cell := range row {
				fmt.Println(cell)
			}
		}
	}
}
