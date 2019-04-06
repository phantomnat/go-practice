package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func solve(caseNo, N int, lydiaPath string) {
	lydiaPosX := 0
	lydiaPosY := 0

	// Check start and endpoint
	lastWalkPos := ((N - 1) * 2) - 1
	canGoRightFromStart := lydiaPath[0] == 'S'
	canGoDownFromStart := lydiaPath[0] == 'E'
	canGoRightToFinished := lydiaPath[lastWalkPos] == 'S'
	canGoDownToFinished := lydiaPath[lastWalkPos] == 'E'

	// easy case
	var ans bytes.Buffer

	if canGoDownFromStart && canGoRightToFinished {
		for i := 0; i < N-1; i++ {
			ans.WriteRune('S')
		}
		for i := 0; i < N-1; i++ {
			ans.WriteRune('E')
		}
	} else if canGoRightFromStart && canGoDownToFinished {
		for i := 0; i < N-1; i++ {
			ans.WriteRune('E')
		}
		for i := 0; i < N-1; i++ {
			ans.WriteRune('S')
		}
	} else if canGoDownFromStart && canGoDownToFinished {
		// find E cross
		// find at least 2 ss
		var firstCross int
		for i := 0; i < len(lydiaPath); i++ {
			if i < len(lydiaPath)-1 {
				if lydiaPath[i] == 'S' && lydiaPath[i+1] == 'S' {
					firstCross = lydiaPosY + 1
					break
				}
			}
			if lydiaPath[i] == 'S' {
				lydiaPosY++
			} else {
				lydiaPosX++
			}
		}

		for i := 0; i < firstCross; i++ {
			ans.WriteRune('S')
		}
		for i := 0; i < N-1; i++ {
			ans.WriteRune('E')
		}
		for i := 0; i < (N-1)-firstCross; i++ {
			ans.WriteRune('S')
		}
	} else {
		// find S cross
		// find at least 2 ee
		var firstCross int
		for i := 0; i < len(lydiaPath); i++ {
			if i < len(lydiaPath)-1 {
				if lydiaPath[i] == 'E' && lydiaPath[i+1] == 'E' {
					firstCross = lydiaPosX + 1
					break
				}
			}
			if lydiaPath[i] == 'S' {
				lydiaPosY++
			} else {
				lydiaPosX++
			}
		}

		for i := 0; i < firstCross; i++ {
			ans.WriteRune('E')
		}
		for i := 0; i < N-1; i++ {
			ans.WriteRune('S')
		}
		for i := 0; i < (N-1)-firstCross; i++ {
			ans.WriteRune('E')
		}
	}

	printf("Case #%d: %s\n", caseNo, ans.String())
}

func main() {
	defer writer.Flush()

	var T int

	scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		var N int
		var lydiaPath string
		scanf("%d\n", &N)
		scanf("%s\n", &lydiaPath)
		solve(i+1, N, lydiaPath)
	}
}
