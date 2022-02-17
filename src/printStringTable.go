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

func addMarginLeft(text string, size int) string {
	prefix := ""
	for i := 0; i < size; i++ {
		prefix += " "
	}
	return prefix + text
}

func addMarginRight(text string, size int) string {
	suffix := ""
	for i := 0; i < size; i++ {
		suffix += " "
	}
	return text + suffix
}

func (a StrArr) PrintStrArr() {
	margin := a.getLongestString()
	lengths := a.getEachItemLength()
	for k, v := range a.Arr {
		fmt.Printf("%s\n", addMarginLeft(v, margin-lengths[k]))
	}
}

func (t StrTable) PrintStrTable(gap string, align rune) {
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
				line += fmt.Sprintf("%s%s", addMarginRight(
					v.Arr[i], marginArr[k]-lengthsArr[k][i]), gap)
			case 'L':
				line += fmt.Sprintf("%s%s", addMarginLeft(
					v.Arr[i], marginArr[k]-lengthsArr[k][i]), gap)
			}
		}
		fmt.Println(line)
	}
}
