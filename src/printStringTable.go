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

type Interval struct {
	Start int
	End   int
}

type Gap struct {
	Char   string
	Length int
}

type Align struct {
	Position string
}

func NewStrArr(a StrArr) *StrArr {
	return &StrArr{Arr: a.Arr}
}

func NewStrTable(t StrTable) *StrTable {
	return &StrTable{ArrArr: t.ArrArr}
}

func alignRight(text string, size int) string {
	prefix := fillWithTimes(" ", size)
	return prefix + text
}

func alignLeft(text string, size int) string {
	suffix := fillWithTimes(" ", size)
	return text + suffix
}

func alignCenter(text string, size int) string {
	pLen := size / 2
	prefix := fillWithTimes(" ", pLen)
	suffix := fillWithTimes(" ", size-pLen)
	return prefix + text + suffix
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

func (a StrArr) getLongestString(i Interval) int {
	longestStrLength := 0
	for _, v := range a.Arr[i.Start:i.End] {
		if len(v) > longestStrLength {
			longestStrLength = len(v)
		}
	}
	return longestStrLength
}

func (a StrArr) getEachItemLength(i Interval) []int {
	lengths := make([]int, len(a.Arr[i.Start:i.End]))
	for k, v := range a.Arr[i.Start:i.End] {
		lengths[k] = len(v)
	}
	return lengths
}

func (a StrArr) getArrayLength(i Interval) int {
	return len(a.Arr[i.Start:i.End])
}

func (a StrArr) PrintStrArr(i Interval, al Align) {
	margin := a.getLongestString(i)
	lengths := a.getEachItemLength(i)
	for k, v := range a.Arr {
		switch al.Position {
		case "R":
			fmt.Printf("%s\n", alignRight(v, margin-lengths[k]))
		case "C":
			fmt.Printf("%s\n", alignCenter(v, margin-lengths[k]))
		default:
			fmt.Printf("%s\n", alignLeft(v, margin-lengths[k]))
		}
	}
}

func (t StrTable) getGapBetweenColumns(g Gap) string {
	g.Char = getFirstNotEmptyChar(g.Char)
	if g.Length < 1 {
		g.Length = 1
	} else if g.Length > 3 {
		g.Length = 3
	}

	gap := ""
	switch g.Length {
	case 1:
		gap = " "
	case 2:
		gap = g.Char + " "
	default:
		gap = " " + g.Char + " "
	}
	return gap
}

func (t StrTable) getMargins(i Interval) []int {
	marginArr := make([]int, len(t.ArrArr))
	for k := range marginArr[:] {
		marginArr[k] = t.ArrArr[k].getLongestString(i)
	}
	return marginArr
}

func (t StrTable) getLengths(i Interval) [][]int {
	lengthsArr := make([][]int, len(t.ArrArr))
	for k := range lengthsArr[:] {
		lengthsArr[k] = t.ArrArr[k].getEachItemLength(i)
	}
	return lengthsArr
}

func (t StrTable) printTableFirstLine(i Interval, gap string) {
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
	marginArr := t.getMargins(i)
	for k := range marginArr {
		firstLine += fmt.Sprint(fillWithTimes("-", marginArr[k]))
		if k != len(t.ArrArr)-1 {
			firstLine += betweenColumns
		}
	}
	fmt.Println(firstLine)
}

func (t StrTable) printTableBody(i Interval, al Align, gap string) {
	numItems := t.ArrArr[:][0].getArrayLength(i)
	marginArr := t.getMargins(i)
	lengthsArr := t.getLengths(i)

	for j := 0; j < numItems; j++ {
		line := ""
		for k, v := range t.ArrArr[:] {
			v.Arr = v.Arr[i.Start:i.End]
			if al.Position == "R" {
				line += alignRight(v.Arr[j], marginArr[k]-lengthsArr[k][j])
			} else if al.Position == "C" {
				line += alignCenter(v.Arr[j], marginArr[k]-lengthsArr[k][j])
			} else {
				line += alignLeft(v.Arr[j], marginArr[k]-lengthsArr[k][j])
			}
			if k != len(t.ArrArr)-1 {
				line += gap
			}
		}
		fmt.Println(line)
	}
}

func (t StrTable) PrintStrTable(i Interval, al Align, g Gap) {
	if i.End == 0 {
		i.End = len(t.ArrArr[:][0].Arr)
	}
	gap := t.getGapBetweenColumns(g)
	t.printTableFirstLine(i, gap)
	t.printTableBody(i, al, gap)
}
