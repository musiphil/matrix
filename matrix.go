package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type key struct{ Row, Col string }

type matrix struct {
	rows    []string
	rowsMap map[string]bool
	cols    []string
	colsMap map[string]bool
	body    map[key]string
}

func newMatrix() *matrix {
	return &matrix{
		rowsMap: make(map[string]bool),
		colsMap: make(map[string]bool),
		body:    make(map[key]string),
	}
}

func (m *matrix) Add(r, c, b string) {
	if !m.rowsMap[r] {
		m.rows = append(m.rows, r)
		m.rowsMap[r] = true
	}
	if !m.colsMap[c] {
		m.cols = append(m.cols, c)
		m.colsMap[c] = true
	}
	m.body[key{Row: r, Col: c}] = b
}

func (m *matrix) Print(w io.Writer, output_sep, missing string) {
	for _, c := range m.cols {
		fmt.Fprint(w, output_sep, c)
	}
	fmt.Fprintln(w)
	for _, r := range m.rows {
		fmt.Fprint(w, r)
		for _, c := range m.cols {
			b, ok := m.body[key{Row: r, Col: c}]
			if !ok {
				b = missing
			}
			fmt.Fprint(w, output_sep, b)
		}
		fmt.Fprintln(w)
	}
}

var (
	inputSep  = flag.String("input-sep", "\t", "input separator")
	outputSep = flag.String("output-sep", "\t", "output separator")
	missing   = flag.String("missing", "", "missing value indicator")
)

func main() {
	flag.Parse()

	m := newMatrix()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.SplitN(line, *inputSep, 3)
		var r, c, b string
		switch {
		case len(items) <= 1:
			r = items[0]
		case len(items) <= 2:
			r = items[0]
			c = items[1]
		default:
			r = items[0]
			c = items[1]
			b = items[2]
		}
		m.Add(r, c, b)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	m.Print(os.Stdout, *outputSep, *missing)
}
