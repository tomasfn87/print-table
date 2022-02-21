package table

import (
	"fmt"
)

type StrArr struct {
	Arr []string
}

type Align struct {
	Position string
}

func NewStrArr(a StrArr) *StrArr {
	return &StrArr{Arr: a.Arr}
}

func (a StrArr) PrintStrArr(al Align) {
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
	lengths := make([]int, len(a.Arr))
	for k, v := range a.Arr {
		lengths[k] = len(v)
	}
	return lengths
}

func (a StrArr) getArrayLength() int {
	return len(a.Arr)
}

func alignRight(text string, size int) string {
	return fillWithTimes(" ", size) + text
}

func alignCenter(text string, size int) string {
	pLen := size / 2
	prefix := fillWithTimes(" ", pLen)
	suffix := fillWithTimes(" ", size-pLen)
	return prefix + text + suffix
}

func alignLeft(text string, size int) string {
	return text + fillWithTimes(" ", size)
}

func fillWithTimes(char string, size int) string {
	fill := ""
	for i := 0; i < size; i++ {
		fill += fmt.Sprintf("%c", char[0])
	}
	return fill
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

type Mark struct {
	Lines int
}

func NewStrTable(t StrTable) *StrTable {
	return &StrTable{ArrArr: t.ArrArr}
}

func (t StrTable) PrintStrTable(i Interval, al Align, g Gap, m Mark) {
	table := t.ArrArr[:]
	tableLength := len(table[0].Arr)

	if i.End < 1 {
		i.End = tableLength
	}

	for k, v := range table {
		table[k].Arr = v.Arr[i.Start:i.End]
	}

	gap := t.getGapBetweenColumns(g)
	t.printTableBody(table, al, gap, m)
}

func (t StrTable) printTableBody(table []StrArr, al Align, gap string, m Mark) {
	numItems := table[0].getArrayLength()
	for i := 0; i < numItems; i++ {
		if m.Lines <= 0 {
			if i == 0 {
				t.printColumnIndicator(table, gap)
			}
		} else {
			if m.Lines == 1 && i != 0 {
				t.printColumnIndicator(table, gap)
			} else if i == 0 || i == m.Lines || i > m.Lines && (i%m.Lines) == 0 {
				t.printColumnIndicator(table, gap)
			}
		}
		t.printLine(table, al, gap, "N", i)
	}
}

func (t StrTable) PrintTitledStrTable(i Interval, al Align, g Gap, m Mark) {
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
	t.printTitledTableBody(table, al, gap, m)
}

func (t StrTable) printTitledTableBody(table []StrArr, al Align, gap string, m Mark) {
	numItems := table[0].getArrayLength()
	for i := 0; i < numItems; i++ {
		if m.Lines <= 0 {
			if i == 1 {
				t.printColumnIndicator(table, gap)
			}
		} else {
			if m.Lines == 1 && i != 0 {
				t.printColumnIndicator(table, gap)
			} else if i == 1 || i == m.Lines+1 || i > m.Lines+1 && (i%m.Lines) == 1 {
				t.printColumnIndicator(table, gap)
			}
		}
		t.printLine(table, al, gap, "Y", i)
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

func getFirstNotEmptyChar(text string) string {
	for _, v := range text {
		if fmt.Sprintf("%c", v) != " " {
			return fmt.Sprintf("%c", v)
		}
	}
	return " "
}

func (t StrTable) printColumnIndicator(table []StrArr, gap string) {
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
	marginsArr := t.getMargins(table)
	for k := range marginsArr {
		columnIndicator += fmt.Sprint(fillWithTimes("-", marginsArr[k]))
		if k != len(t.ArrArr)-1 {
			columnIndicator += betweenColumns
		}
	}
	fmt.Println(columnIndicator)
}

func (t StrTable) printLine(table []StrArr, al Align, gap string, title string, index int) {
	marginsArr := t.getMargins(table)
	lengthsArr := t.getLengths(table)
	line := ""
	for k, v := range table {
		if al.Position == "R" {
			line +=
				alignRight(v.Arr[index], marginsArr[k]-lengthsArr[k][index])
		} else if al.Position == "C" {
			line +=
				alignCenter(v.Arr[index], marginsArr[k]-lengthsArr[k][index])
		} else {
			line +=
				alignLeft(v.Arr[index], marginsArr[k]-lengthsArr[k][index])
		}
		if k != len(t.ArrArr)-1 {
			line += gap
		}
	}
	fmt.Println(line)
}

func (t StrTable) getMargins(table []StrArr) []int {
	marginsArr := make([]int, len(t.ArrArr))
	for k := range table {
		marginsArr[k] = table[k].getLongestString()
	}
	return marginsArr
}

func (t StrTable) getLengths(table []StrArr) [][]int {
	lengthsArr := make([][]int, len(table))
	for k := range lengthsArr {
		lengthsArr[k] = table[k].getEachItemLength()
	}
	return lengthsArr
}
