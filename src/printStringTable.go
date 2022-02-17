package table

import (
	"fmt"
)

type StrArr struct {
	Arr []string
}

type StrTable struct {
	ArrArr []StrArr
}

func NewStrArr(a StrArr) *StrArr {
	return &StrArr{Arr: a.Arr}
}

func NewStrTable(t StrTable) *StrTable {
	return &StrTable{ArrArr: t.ArrArr}
}

func (a StrArr) getLongestString() int {
	longestStrLength := 0
	for _, v := range a.Arr {
		if len(v) > longestStrLength {
			longestStrLength = len(v)
		}
	}
	return longestStrLength
}

func (a StrArr) getEachItemLength() []int {
	lengths := make([]int, len(a.Arr), len(a.Arr))
	for k, v := range a.Arr {
		lengths[k] = len(v)
	}
	return lengths
}

func (a StrArr) getArrayLength() int {
	return len(a.Arr)
}

func alignRight(text string, size int) string {
	prefix := ""
	for i := 0; i < size; i++ {
		prefix += " "
	}
	return prefix + text
}

func alignLeft(text string, size int) string {
	suffix := ""
	for i := 0; i < size; i++ {
		suffix += " "
	}
	return text + suffix
}

func fillWithTimes(char string, size int) string {
	fill := ""
	for i := 0; i < size; i++ {
		fill += fmt.Sprintf("%c", char[0])
	}
	return fill
}

func getFirstNotEmptyChar(text string) string {
	for _, v := range text {
		if fmt.Sprintf("%c", v) != " " {
			return fmt.Sprintf("%c", v)
		}
	}
	return " "
}

func (a StrArr) PrintStrArr(align string) {
	margin := a.getLongestString()
	lengths := a.getEachItemLength()
	for k, v := range a.Arr {
		switch align {
		case "R":
			fmt.Printf("%s\n", alignRight(v, margin-lengths[k]))
		default:
			fmt.Printf("%s\n", alignLeft(v, margin-lengths[k]))
		}
	}
}

func (t StrTable) PrintStrTable(align string, gapChar string, gapLength int) {
	gapChar = getFirstNotEmptyChar(gapChar)

	if gapLength < 1 {
		gapLength = 1
	} else if gapLength > 3 {
		gapLength = 3
	}

	gap := ""
	switch gapLength {
	case 1:
		gap = " "
	case 2:
		gap = gapChar + " "
	default:
		gap = " " + gapChar + " "
	}

	marginArr := make([]int, len(t.ArrArr))
	for k := range marginArr {
		marginArr[k] = t.ArrArr[k].getLongestString()
	}

	lengthsArr := make([][]int, len(t.ArrArr))
	for k := range lengthsArr {
		lengthsArr[k] = t.ArrArr[k].getEachItemLength()
	}

	numItems := t.ArrArr[:][0].getArrayLength()

	firstLine := ""
	betweenColumns := ""
	switch len(gap) {
	case 1:
		betweenColumns += "+"
	case 2:
		betweenColumns += "+-"
	case 3:
		betweenColumns += "-+-"
	}
	for k := range marginArr {
		firstLine += fmt.Sprint(fillWithTimes("-", marginArr[k]))
		if k != len(t.ArrArr)-1 {
			firstLine += betweenColumns
		}
	}
	fmt.Println(firstLine)

	for i := 0; i < numItems; i++ {
		line := ""
		for k, v := range t.ArrArr {
			if align == "R" {
				line += alignRight(v.Arr[i], marginArr[k]-lengthsArr[k][i])
			} else {
				line += alignLeft(v.Arr[i], marginArr[k]-lengthsArr[k][i])
			}
			if k != len(t.ArrArr)-1 {
				line += gap
			}
		}
		fmt.Println(line)
	}
}
