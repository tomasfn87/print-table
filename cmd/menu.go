package main

import (
	"fmt"
	"strconv"

	table "code.repo/table/src"
)

func Menu() {
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

	numberMapArray := table.NewStrTable(table.StrTable{
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

	if intervalStart > 8 {
		intervalStart = 8
	} else if intervalStart < 0 {
		intervalStart = 0
	}

	fmt.Println()
	intervalEndInput := ""
	fmt.Println("(from 1 to 9 for last value) or ENTER for default 9")
	fmt.Printf("Finish printing at element: ")
	fmt.Scanf("%s", &intervalEndInput)
	intervalEnd, _ := strconv.Atoi(intervalEndInput)

	if intervalEnd == intervalStart && intervalStart != 0 {
		intervalEnd++
	}

	if intervalEnd <= 0 || intervalEnd > 9 || intervalEnd < intervalStart {
		intervalEnd = 9
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
	numberMapArray.PrintStrTable(
		table.Interval{Start: intervalStart, End: intervalEnd},
		table.Align{Position: alignInput},
		table.Gap{Char: gapInput, Length: gapLength},
	)
}

func main() {
	for {
		Menu()
		fmt.Println()
	}
}
