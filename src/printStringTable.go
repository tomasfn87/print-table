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

func fillWithTimes(char string, size int) string {
	fill := ""
	for i := 0; i < size; i++ {
		fill += fmt.Sprintf("%c", char[0])
	}
	return fill
}

func alignRight(text string, size int) string {
	return fillWithTimes(" ", size) + text
}

func alignLeft(text string, size int) string {
	return text + fillWithTimes(" ", size)
}

func alignCenter(text string, size int) string {
	pLen := size / 2
	prefix := fillWithTimes(" ", pLen)
	suffix := fillWithTimes(" ", size-pLen)
	return prefix + text + suffix
}

func getFirstNotEmptyChar(text string) string {
	for _, v := range text {
		if fmt.Sprintf("%c", v) != " " {
			return fmt.Sprintf("%c", v)
		}
	}
	return " "
}

func (a StrArr) getLongestString() int {
	longestStrLength := 0
	for _, v := range a.Arr[:] {
		if len(v) > longestStrLength {
			longestStrLength = len(v)
		}
	}
	return longestStrLength
}

func (a StrArr) getEachItemLength() []int {
	lengths := make([]int, len(a.Arr[:]))
	for k, v := range a.Arr[:] {
		lengths[k] = len(v)
	}
	return lengths
}

func (a StrArr) getArrayLength() int {
	return len(a.Arr[:])
}

func (a StrArr) PrintStrArr(i Interval, al Align) {
	margin := a.getLongestString()
	lengths := a.getEachItemLength()
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
		if g.Char == " " {
			gap = " " + g.Char
		} else {
			gap = g.Char + " "
		}
	default:
		gap = " " + g.Char + " "
	}
	return gap
}

func (t StrTable) getMargins() []int {
	marginArr := make([]int, len(t.ArrArr))
	for k := range marginArr[:] {
		marginArr[k] = t.ArrArr[k].getLongestString()
	}
	return marginArr
}

func (t StrTable) getLengths() [][]int {
	lengthsArr := make([][]int, len(t.ArrArr))
	for k := range lengthsArr[:] {
		lengthsArr[k] = t.ArrArr[k].getEachItemLength()
	}
	return lengthsArr
}

func (t StrTable) printLine(al Align, gap string, title string, index int) {
	marginArr := t.getMargins()
	lengthsArr := t.getLengths()

	line := ""
	for k, v := range t.ArrArr[:] {
		v.Arr = v.Arr[:]
		if al.Position == "R" {
			line +=
				alignRight(v.Arr[index], marginArr[k]-lengthsArr[k][index])
		} else if al.Position == "C" {
			line +=
				alignCenter(v.Arr[index], marginArr[k]-lengthsArr[k][index])
		} else {
			line +=
				alignLeft(v.Arr[index], marginArr[k]-lengthsArr[k][index])
		}
		if k != len(t.ArrArr)-1 {
			line += gap
		}
	}
	fmt.Println(line)
}

func (t StrTable) printColumnIndicator(gap string) {
	columnIndicator := ""
	betweenColumns := ""
	switch len(gap) {
	case 1:
		betweenColumns += "+"
	case 2:
		if gap == "  " {
			betweenColumns += "-+"
		} else {
			betweenColumns += "+-"
		}
	case 3:
		betweenColumns += "-+-"
	}
	marginArr := t.getMargins()
	for k := range marginArr {
		columnIndicator += fmt.Sprint(fillWithTimes("-", marginArr[k]))
		if k != len(t.ArrArr)-1 {
			columnIndicator += betweenColumns
		}
	}
	fmt.Println(columnIndicator)
}

func (t StrTable) printTableBody(table []StrArr, al Align, gap string) {
	numItems := t.ArrArr[:][0].getArrayLength()
	for j := 0; j < numItems; j++ {
		if j == 0 {
			t.printColumnIndicator(gap)
		}
		t.printLine(al, gap, "N", j)
	}
}

func (t StrTable) PrintStrTable(i Interval, al Align, g Gap) {
	table := t.ArrArr[:]
	tableLength := len(table[0].Arr)

	if i.End < 1 {
		i.End = tableLength
	}

	for k, v := range table {
		table[k].Arr = v.Arr[i.Start:i.End]
	}

	gap := t.getGapBetweenColumns(g)
	t.printTableBody(table, al, gap)
}

func (t StrTable) printTitledTableBody(table []StrArr, al Align, gap string) {
	numItems := t.ArrArr[:][0].getArrayLength()
	for j := 0; j < numItems; j++ {
		if j == 1 {
			t.printColumnIndicator(gap)
		}
		t.printLine(al, gap, "Y", j)
	}
}

func (t StrTable) PrintTitledStrTable(i Interval, al Align, g Gap) {
	table := t.ArrArr[:]
	tableLength := len(table[0].Arr)
	if i.Start == 0 {
		i.Start = 1
	}

	if i.End <= 1 || i.End < i.Start || i.End > tableLength {
		i.End = tableLength
	}

	for k, v := range table {
		table[k].Arr = append(v.Arr[0:1], v.Arr[i.Start:i.End]...)
	}

	gap := t.getGapBetweenColumns(g)
	t.printTitledTableBody(table, al, gap)
}
