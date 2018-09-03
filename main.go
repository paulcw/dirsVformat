package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

const eltFmt = " %12s %s\n"

func main() {
	var curr *elt
	elts := make([]*elt, 0)
	for in := bufio.NewScanner(os.Stdin); in.Scan(); {
		line := in.Text()
		e, err := parseLine(line)
		if err != nil {
			continue
		}
		if curr == nil {
			curr = e
			continue
		}
		elts = append(elts, e)
	}

	sort.Slice(elts, sortFor(curr, elts))

	r := color.New(color.FgRed)
	y := color.New(color.FgYellow)
	c := color.New(color.FgCyan)
	fmt.Printf(eltFmt, r.Sprint("==>"), curr.path)
	for i, e := range elts {
		f := y
		if i%2 == 0 {
			f = c
		}
		fmt.Printf(eltFmt, f.Sprint(e.pos), f.Sprint(e.path))
	}
}

type elt struct {
	pos  int
	path string
}

// sortFor returns a sort function for the
// given current path element and list of other paths.
// the sort function in question just initializes values
// and returns the result of pathCompare.
func sortFor(c *elt, es []*elt) func(i, j int) bool {
	return func(i, j int) bool {
		return pathCompare(es[i].path, es[j].path, c.path)
	}
}

func pathCompare(p1, p2, comp string) bool {
	if comp == "." || comp == "/" || comp == "~" || comp == "" {
		return p1 < p2
	}
	iHas := strings.HasPrefix(p1, comp)
	jHas := strings.HasPrefix(p2, comp)
	if iHas && jHas {
		return p1 < p2
	}
	if iHas && !jHas {
		return true
	}
	if jHas && !iHas {
		return false
	}
	return pathCompare(p1, p2, filepath.Dir(comp))
}

func parseLine(line string) (*elt, error) {
	line = strings.TrimSpace(line)
	parts := strings.SplitN(line, " ", 2)
	pos, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, err
	}
	return &elt{
		pos:  pos,
		path: strings.TrimSpace(parts[1]),
	}, nil
}
