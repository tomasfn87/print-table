package main

import (
	"fmt"

	table "code.repo/table/src"
)

func main() {
	number4 := 1023.69514865
	numberMap4 := []string{
		"hours",
		fmt.Sprintf("%.5f", number4),
		fmt.Sprintf("%.2f", number4),
		fmt.Sprintf("%v", number4),
		fmt.Sprintf("%.4f", number4),
		fmt.Sprintf("%.6f", number4),
		fmt.Sprintf("%.3f", number4),
		fmt.Sprintf("%.f", number4),
		fmt.Sprintf("%.7f", number4),
		fmt.Sprintf("%.1f", number4),
	}
	NumberMap4 := table.NewStrArr(table.StrArr{Arr: numberMap4})

	number5 := 20.58120105
	numberMap5 := []string{
		"$/hour",
		fmt.Sprintf("%.6f", number5),
		fmt.Sprintf("%.2f", number5),
		fmt.Sprintf("%.4f", number5),
		fmt.Sprintf("%.1f", number5),
		fmt.Sprintf("%v", number5),
		fmt.Sprintf("%.5f", number5),
		fmt.Sprintf("%.f", number5),
		fmt.Sprintf("%.7f", number5),
		fmt.Sprintf("%.3f", number5),
	}
	NumberMap5 := table.NewStrArr(table.StrArr{Arr: numberMap5})

	number6 := 1.14340984
	numberMap6 := []string{
		"interest rate",
		fmt.Sprintf("%.6f", number6),
		fmt.Sprintf("%.1f", number6),
		fmt.Sprintf("%v", number6),
		fmt.Sprintf("%.4f", number6),
		fmt.Sprintf("%.5f", number6),
		fmt.Sprintf("%.2f", number6),
		fmt.Sprintf("%.3f", number6),
		fmt.Sprintf("%.f", number6),
		fmt.Sprintf("%.7f", number6),
	}
	NumberMap6 := table.NewStrArr(table.StrArr{Arr: numberMap6})

	NumberMapArray2 := table.NewStrTable(table.StrTable{
		ArrArr: []table.StrArr{
			*NumberMap4,
			*NumberMap5,
			*NumberMap6,
		},
	})

	fmt.Printf("#4 PrintTitledStrTable:\n")
	fmt.Println("  *  Interval{},")
	fmt.Println("  *  Align{},")
	fmt.Println("  *  Gap{}")
	fmt.Println()
	NumberMapArray2.PrintTitledStrTable(
		table.Interval{},
		table.Align{},
		table.Gap{},
		table.Mark{},
	)

	fmt.Printf("\n#5 PrintTitledStrTable:\n")
	fmt.Println("  *  Interval{Start: 0, End: 0},")
	fmt.Println("  *  Align{ Position: \"R\" },")
	fmt.Println("  *  Gap{Char: \"|\", Length: 3 }")
	fmt.Println("  *  Mark{Lines: 2}")
	fmt.Println()
	NumberMapArray2.PrintTitledStrTable(
		table.Interval{Start: 0, End: 0},
		table.Align{Position: "R"},
		table.Gap{Char: "|", Length: 3},
		table.Mark{Lines: 2},
	)

	fmt.Printf("\n#6 PrintTitledStrTable:\n")
	fmt.Println("  *  Interval{ Start: 1, End: 5 },")
	fmt.Println("  *  Align{ Position: \"C\" },")
	fmt.Println("  *  Gap{ Char: \"|\", Length: 2 }")
	fmt.Println()
	NumberMapArray2.PrintTitledStrTable(
		table.Interval{Start: 1, End: 5},
		table.Align{Position: "C"},
		table.Gap{Char: "|", Length: 2},
		table.Mark{Lines: 1},
	)

}
