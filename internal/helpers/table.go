package helpers

import (
	"errors"
	"fmt"
	"strings"
	"unicode/utf8"
)

type row struct {
	values []string
}

type Table struct {
	columns    int
	offset     int
	header     *row
	rows       []row
	columnSize []int
	valueSize  []int
}

func NewTable(columns int) *Table {
	return &Table{
		columns:    columns,
		columnSize: make([]int, columns),
		valueSize:  make([]int, columns),
		rows:       make([]row, 0),
	}
}

// Set table offset
func (t *Table) SetOffset(offset int) {
	t.offset = offset
}

// Add header to table
func (t *Table) AddHeader(values []string) {
	t.header = &row{values}

	for col, value := range values {
		length := utf8.RuneCountInString(value)

		if length > t.columnSize[col] {
			t.columnSize[col] = length
		}
	}
}

// Add row to table
func (t *Table) AddRow(values []string) error {
	if len(values) != t.columns {
		return errors.New("Invalid number of values")
	}

	t.rows = append(t.rows, row{values})

	for col, value := range values {
		length := utf8.RuneCountInString(value)

		if length > t.columnSize[col] {
			t.columnSize[col] = length
		}

		if length > t.valueSize[col] {
			t.valueSize[col] = length
		}
	}

	return nil
}

// Return table with formatting as string
func (t *Table) ToString() string {
	lines := make([]string, 0, len(t.rows)+2)

	if t.header != nil {
		headerLine := ""
		delimiterLine := ""

		for col, value := range t.header.values {
			length := utf8.RuneCountInString(value)
			offset := 0

			switch col {
			case 0:
				offset = t.offset + t.columnSize[col]
			case 1:
				offset = length + t.columnSize[col]
			case 2:
				offset = length + t.columnSize[col]
			}

			formatString := fmt.Sprintf("%%%ds", offset)
			headerLine += fmt.Sprintf(formatString, value)
			delimiterLine += fmt.Sprintf(formatString, strings.Repeat("=", length))
		}

		lines = append(lines, headerLine)
		lines = append(lines, delimiterLine)
	}

	for i := range t.rows {
		line := ""

		for col, value := range t.rows[i].values {
			offset := 0
			length := utf8.RuneCountInString(value)

			switch col {
			case 0:
				offset = t.offset + t.columnSize[col]
			case 1:
				offset = t.columnSize[col] + length
			case 2:
				prevLength := utf8.RuneCountInString(t.rows[i].values[col-1])
				offset = t.columnSize[col] + t.columnSize[col-1] + length + (t.columnSize[col] - length) - prevLength
			}
			formatString := fmt.Sprintf("%%%ds", offset)

			line += fmt.Sprintf(formatString, value)
		}

		lines = append(lines, line)
	}

	lines = append(lines, "")

	return strings.Join(lines, "\n")
}
