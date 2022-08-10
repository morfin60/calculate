package helpers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type row struct {
	values []string
}

type Table struct {
	columns    int
	header     *row
	rows       []row
	columnSize []int
}

func NewTable(columns int) *Table {
	return &Table{
		columns:    columns,
		columnSize: make([]int, columns),
		rows:       make([]row, 0),
	}
}

func (t *Table) AddHeader(values []string) {
	t.header = &row{values}

	for col, value := range values {
		if len(value) > t.columnSize[col] {
			t.columnSize[col] = len(value)
		}
	}
}

func (t *Table) AddRow(values []string) error {
	if len(values) != t.columns {
		return errors.New("Invalid number of values")
	}

	t.rows = append(t.rows, row{values})

	for col, value := range values {
		if len(value) > t.columnSize[col] {
			t.columnSize[col] = len(value)
		}
	}

	return nil
}

func (t *Table) ToString() string {
	lines := make([]string, 0, len(t.rows)+2)

	if t.header != nil {
		headerLine := ""
		delimiterLine := ""

		for col, value := range t.header.values {
			colSize := strconv.Itoa(t.columnSize[col] + 5)
			headerLine += fmt.Sprintf("%-"+colSize+"s", value)
			delimiterLine += fmt.Sprintf("%-"+colSize+"s", strings.Repeat("=", len(value)))
		}

		lines = append(lines, headerLine)
		lines = append(lines, delimiterLine)
	}

	for i := range t.rows {
		line := ""

		for col, value := range t.rows[i].values {
			colSize := strconv.Itoa(t.columnSize[col] + 5)

			line += fmt.Sprintf("%-"+colSize+"s", value)
		}

		lines = append(lines, line)
	}

	lines = append(lines, "")

	return strings.Join(lines, "\n")
}
