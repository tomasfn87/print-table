package table

import "fmt"

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

func (a StrArr) PrintStrArr(align rune) {
	margin := a.getLongestString()
	lengths := a.getEachItemLength()
	for k, v := range a.Arr {
		switch align {
		case 'R':
			fmt.Printf("%s\n", alignRight(v, margin-lengths[k]))
		case 'L':
			fmt.Printf("%s\n", alignLeft(v, margin-lengths[k]))
		}
	}
}

func (t StrTable) PrintStrTable(align rune, gap string) {
	marginArr := make([]int, len(t.ArrArr))
	for k := range marginArr {
		marginArr[k] = t.ArrArr[k].getLongestString()
	}

	lengthsArr := make([][]int, len(t.ArrArr))
	for k := range lengthsArr {
		lengthsArr[k] = t.ArrArr[k].getEachItemLength()
	}

	numItems := t.ArrArr[:][0].getArrayLength()

	for i := 0; i < numItems; i++ {
		line := ""
		for k, v := range t.ArrArr {
			switch align {
			case 'R':
				line += alignRight(v.Arr[i], marginArr[k]-lengthsArr[k][i])
			case 'L':
				line += alignLeft(v.Arr[i], marginArr[k]-lengthsArr[k][i])
			}
			if k != len(t.ArrArr)-1 {
				line += gap
			}
		}
		fmt.Println(line)
	}
}
