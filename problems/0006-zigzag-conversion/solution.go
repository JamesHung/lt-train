package zigzagconversion

import (
	"fmt"
	"strings"
)

const debugZigzag = true

// convert rearranges s into its zigzag representation with numRows rows and reads row by row.
func convert(s string, numRows int) string {
	if debugZigzag {
		fmt.Printf("convert input s=%q numRows=%d\n", s, numRows)
	}

	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([]strings.Builder, numRows)
	for i := range rows {
		rows[i].Grow(len(s) / numRows)
	}

	curRow := 0
	goingDown := false
	for i := 0; i < len(s); i++ {
		rows[curRow].WriteByte(s[i])
		if debugZigzag {
			fmt.Printf("char=%c row=%d\n", s[i], curRow)
		}
		if curRow == 0 || curRow == numRows-1 {
			goingDown = !goingDown
		}
		if goingDown {
			curRow++
		} else {
			curRow--
		}
	}

	var result strings.Builder
	result.Grow(len(s))
	for i := range rows {
		result.WriteString(rows[i].String())
	}

	output := result.String()
	if debugZigzag {
		fmt.Printf("convert output=%q\n", output)
	}
	return output
}
