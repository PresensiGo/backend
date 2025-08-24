package utils

import "fmt"

func ColumnToName(col int) (string, error) {
	if col <= 0 {
		return "", fmt.Errorf("column index must be a positive integer")
	}
	name := ""
	for col > 0 {
		col--
		name = string(rune('A'+col%26)) + name
		col /= 26
	}
	return name, nil
}
