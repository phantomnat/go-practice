package main

import (
	"bufio"
	"fmt"
	"os"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

func reverse(a []byte) {
	n := len(a)
	for i := n/2 - 1; i >= 0; i-- {
		opp := n - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
}

func removeLeadingZero(num []byte) string {
	n := len(num)
	for i := 0; i < n; i++ {
		if num[i] == '0' {
			continue
		}
		return string(num[i:])
	}
	return "0"
}

func solveDigit(num string, n, index, sub, caseNo int, A, B []byte) bool {
	if index < 0 {
		reverse(A)
		reverse(B)
		printf("Case #%d: %s %s\n", caseNo, removeLeadingZero(A), removeLeadingZero(B))
		return true
	}

	no := int(num[index]-'0') - sub
	if no < 0 && index == 0 {
		// cannot sub
		return false
	}

	for i := 9; i >= 0; i-- {
		cannotSub := no-i < 0 && index == 0
		aIs4 := i == 4
		bIs4 := (no-i == 4) || (no-i < 0 && (no+10)-i == 4)
		if aIs4 || cannotSub || bIs4 {
			continue
		}

		b := no - i
		sub = 0
		if no-i < 0 {
			b = (no + 10) - i
			sub = 1
		}
		if solveDigit(num, n, index-1, sub, caseNo, append(A, uint8(i+'0')), append(B, uint8(b+'0'))) {
			return true
		}
	}
	return false
}

func main() {
	defer writer.Flush()

	var T int

	scanf("%d\n", &T)
	for i := 0; i < T; i++ {
		var num string
		scanf("%s\n", &num)
		solveDigit(num, len(num), len(num)-1, 0, i+1, nil, nil)
	}
}
