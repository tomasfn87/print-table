package main

import (
	"fmt"
	"strconv"

	table "code.repo/table/src"
)

func Menu() {
	number4 := 1023.69514865
	numberMap1 := []string{
		"hours",
		fmt.Sprintf("%v", number4),
		fmt.Sprintf("%.7f", number4),
		fmt.Sprintf("%.6f", number4),
		fmt.Sprintf("%.5f", number4),
		fmt.Sprintf("%.4f", number4),
		fmt.Sprintf("%.2f", number4),
		fmt.Sprintf("%.3f", number4),
		fmt.Sprintf("%.1f", number4),
		fmt.Sprintf("%.f", number4),
	}
	NumberMap1 := table.NewStrArr(table.StrArr{Arr: numberMap1})

	number5 := 20.58120105
	numberMap2 := []string{
		"$/hour",
		fmt.Sprintf("%v", number5),
		fmt.Sprintf("%.7f", number5),
		fmt.Sprintf("%.6f", number5),
		fmt.Sprintf("%.5f", number5),
		fmt.Sprintf("%.4f", number5),
		fmt.Sprintf("%.3f", number5),
		fmt.Sprintf("%.2f", number5),
		fmt.Sprintf("%.1f", number5),
		fmt.Sprintf("%.f", number5),
	}
	NumberMap2 := table.NewStrArr(table.StrArr{Arr: numberMap2})

	number6 := 1.14340984
	numberMap3 := []string{
		"interest rate",
		fmt.Sprintf("%v", number6),
		fmt.Sprintf("%.7f", number6),
		fmt.Sprintf("%.6f", number6),
		fmt.Sprintf("%.5f", number6),
		fmt.Sprintf("%.4f", number6),
		fmt.Sprintf("%.3f", number6),
		fmt.Sprintf("%.2f", number6),
		fmt.Sprintf("%.1f", number6),
		fmt.Sprintf("%.f", number6),
	}
	NumberMap3 := table.NewStrArr(table.StrArr{Arr: numberMap3})

	NumberMapArray := table.NewStrTable(table.StrTable{
		ArrArr: []table.StrArr{
			*NumberMap1,
			*NumberMap2,
			*NumberMap3,
		},
	})

	fmt.Println("Let's format and print a table: ")
	fmt.Println()

	intervalStartInput := ""
	fmt.Println("(from 0 to 8 for last value) or ENTER for default 0")
	fmt.Printf("Start to print from element: ")
	fmt.Scanf("%s", &intervalStartInput)
	intervalStart, _ := strconv.Atoi(intervalStartInput)
	intervalStart++

	if intervalStart > 9 {
		intervalStart = 9
	} else if intervalStart < 1 {
		intervalStart = 1
	}

	fmt.Println()
	intervalEndInput := ""
	fmt.Println("(from 1 to 9 for last value) or ENTER for default 9")
	fmt.Printf("Finish printing at element: ")
	fmt.Scanf("%s", &intervalEndInput)
	intervalEnd, _ := strconv.Atoi(intervalEndInput)
	intervalEnd++

	if intervalEnd == 1 && intervalEnd == intervalStart {
		intervalEnd = 10
	}

	if intervalEnd == intervalStart {
		intervalEnd++
	}

	if intervalEnd <= 0 || intervalEnd > 10 || intervalEnd < intervalStart {
		intervalEnd = 10
	}

	fmt.Println()
	alignInput := ""
	fmt.Println("(R or C) {right, center} or ENTER for default Left alignment")
	fmt.Printf("Alignment: ")
	fmt.Scanf("%s", &alignInput)

	fmt.Println()
	gapInput := ""
	fmt.Println("Single character between columns or ENTER for empty space")
	fmt.Printf("Gap character: ")
	fmt.Scanf("%s", &gapInput)

	fmt.Println()
	gapLengthInput := ""
	fmt.Println("Length of gap between columns (1, 2 or 3) or ENTER for 1")
	fmt.Printf("Gap length: ")
	fmt.Scanf("%s", &gapLengthInput)
	gapLength, _ := strconv.Atoi(gapLengthInput)

	fmt.Println()
	titleInput := ""
	fmt.Println("(Y or N) {yes, no} or ENTER for N")
	fmt.Printf("Print title: ")
	fmt.Scanf("%s", &titleInput)

	fmt.Println()
	markInput := ""
	fmt.Println("Add a line marker every x lines or ENTER for default 0")
	fmt.Printf("x: ")
	fmt.Scanf("%s", &markInput)
	mark, _ := strconv.Atoi(markInput)

	if mark < 0 {
		mark = 0
	}

	fmt.Println()
	if titleInput == "Y" {
		NumberMapArray.PrintTitledStrTable(
			table.Interval{Start: intervalStart, End: intervalEnd},
			table.Align{Position: alignInput},
			table.Gap{Char: gapInput, Length: gapLength},
			table.Mark{Lines: mark},
		)
		return
	}
	NumberMapArray.PrintStrTable(
		table.Interval{Start: intervalStart, End: intervalEnd},
		table.Align{Position: alignInput},
		table.Gap{Char: gapInput, Length: gapLength},
		table.Mark{Lines: mark},
	)
}

func main() {
	for {
		Menu()
		fmt.Println()
	}
}
