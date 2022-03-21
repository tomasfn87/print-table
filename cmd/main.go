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

	NumberMapArray := table.NewStrTable(table.StrTable{
		ArrArr: []table.StrArr{
			*NumberMap1,
			*NumberMap2,
			*NumberMap3,
		},
	})

	// #1 - passing blank fields
	fmt.Printf("#1 PrintStrTable: no arguments\n")
	fmt.Println("  *  Interval{},")
	fmt.Println("  *  Align{},")
	fmt.Println("  *  Gap{}")
	fmt.Println("  *  Mark{}")
	fmt.Println()
	NumberMapArray.PrintStrTable(
		table.Interval{},
		table.Align{},
		table.Gap{},
		table.Mark{},
	)
	// #2 - passing the same values and if all fields are blank
	fmt.Printf("\n#2 PrintStrTable: arguments cloning #1's behavior:\n")
	fmt.Println("  *  Interval{ Start: 0, End: 0 },")
	fmt.Println("  *  Align{ Position: \"\" },")
	fmt.Println("  *  Gap{ Char: \"\", Length: 0 }")
	fmt.Println("  *  Mark{ Lines: 0}")
	fmt.Println()
	NumberMapArray.PrintStrTable(
		table.Interval{Start: 0, End: 0},
		table.Align{Position: ""},
		table.Gap{Char: "|", Length: 0},
		table.Mark{Lines: 0},
	)
	// #1 and #2 produce the same output

	// #3 - passing other values
	fmt.Printf("\n#3 PrintStrTable: other arguments:\n")
	fmt.Println("  *  Interval{ Start: 1, End: 8 },")
	fmt.Println("  *  Align{ Position: \"R\" },")
	fmt.Println("  *  Gap{ Char: \"|\", Length: 3 },")
	fmt.Println("  *  Mark{ Lines: 2}")
	fmt.Println()
	NumberMapArray.PrintStrTable(
		table.Interval{Start: 1, End: 8},
		table.Align{Position: "R"},
		table.Gap{Char: "|", Length: 3},
		table.Mark{Lines: 2},
	)
	// #3 is a lil' bit different
}
