package main

import (
	"fmt"

	table "code.repo/table/src"
)

func main() {
	number1 := 420.69514865
	numberMap1 := []string{
		fmt.Sprintf("%.5f", number1),
		fmt.Sprintf("%.2f", number1),
		fmt.Sprintf("%v", number1),
		fmt.Sprintf("%.4f", number1),
		fmt.Sprintf("%.6f", number1),
		fmt.Sprintf("%.3f", number1),
		fmt.Sprintf("%.f", number1),
		fmt.Sprintf("%.7f", number1),
		fmt.Sprintf("%.1f", number1),
	}
	NumberMap1 := table.NewStrArr(table.StrArr{Arr: numberMap1})

	number2 := 1.37822145
	numberMap2 := []string{
		fmt.Sprintf("%.6f", number2),
		fmt.Sprintf("%.2f", number2),
		fmt.Sprintf("%.4f", number2),
		fmt.Sprintf("%.1f", number2),
		fmt.Sprintf("%v", number2),
		fmt.Sprintf("%.5f", number2),
		fmt.Sprintf("%.f", number2),
		fmt.Sprintf("%.7f", number2),
		fmt.Sprintf("%.3f", number2),
	}
	NumberMap2 := table.NewStrArr(table.StrArr{Arr: numberMap2})

	number3 := 7.58120105
	numberMap3 := []string{
		fmt.Sprintf("%.6f", number3),
		fmt.Sprintf("%.1f", number3),
		fmt.Sprintf("%v", number3),
		fmt.Sprintf("%.4f", number3),
		fmt.Sprintf("%.5f", number3),
		fmt.Sprintf("%.2f", number3),
		fmt.Sprintf("%.3f", number3),
		fmt.Sprintf("%.f", number3),
		fmt.Sprintf("%.7f", number3),
	}
	NumberMap3 := table.NewStrArr(table.StrArr{Arr: numberMap3})

	// NumberMap1.PrintStrArr()
	// NumberMap2.PrintStrArr()
	// NumberMap3.PrintStrArr()
	numberMapArray := table.NewStrTable(table.StrTable{
		ArrArr: []table.StrArr{
			*NumberMap1,
			*NumberMap2,
			*NumberMap3,
		},
	})
	numberMapArray.PrintStrTable('L', " | ")
	fmt.Println()
	numberMapArray.PrintStrTable('R', " | ")
}
