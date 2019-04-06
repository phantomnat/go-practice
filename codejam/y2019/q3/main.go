package q3

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"sort"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

const MaxN = 10000

var primeTable = make([]bool, MaxN)
var primes = make([]int, 0, MaxN)

func generatePrimes() {
	var x, y, n int
	nsqrt := math.Sqrt(MaxN)

	for x = 1; float64(x) <= nsqrt; x++ {
		for y = 1; float64(y) <= nsqrt; y++ {
			n = 4*(x*x) + y*y
			if n <= MaxN && (n%12 == 1 || n%12 == 5) {
				primeTable[n] = !primeTable[n]
			}
			n = 3*(x*x) + y*y
			if n <= MaxN && n%12 == 7 {
				primeTable[n] = !primeTable[n]
			}
			n = 3*(x*x) - y*y
			if x > y && n <= MaxN && n%12 == 11 {
				primeTable[n] = !primeTable[n]
			}
		}
	}

	for n = 5; float64(n) <= nsqrt; n++ {
		if primeTable[n] {
			for y = n * n; y < MaxN; y += n * n {
				primeTable[y] = false
			}
		}
	}

	primeTable[2] = true
	primeTable[3] = true

	for x = 0; x < len(primeTable)-1; x++ {
		if primeTable[x] {
			primes = append(primes, x)
		}
	}
}

type Pair struct {
	A, B int
}

type UsedNums []int

func (p UsedNums) Len() int           { return len(p) }
func (p UsedNums) Less(i, j int) bool { return p[i] < p[j] }
func (p UsedNums) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func divmod(num, d int) (int, int) {
	return num / d, num % d
}

func findPrimeDivisors(num, N int) []Pair {
	result := make([]Pair, 0)

	for i := len(primes) - 1; i >= 0; i-- {
		sq := primes[i]
		if sq > num || sq > N {
			continue
		}

		q, r := divmod(num, sq)
		if r != 0 || q >= N || !primeTable[q] {
			continue
		}
		result = append(result, Pair{sq, q})
	}

	return result
}

func solve(caseNo, N, L int, nums []int) {
	num := nums[0]

	pairs := findPrimeDivisors(num, N)
	for _, pair := range pairs {
		if ans := solvePair(pair.A, pair.B, L, N, nums); ans != "" {
			printf("Case #%d: %s\n", caseNo, ans)
			return
		}
	}

	printf("Case #%d: %s\n", caseNo, "WRONG")
	panic("error")
}

func tryAddPrime(num int, charTable map[int]byte, usedNums []int) []int {
	if _, ok := charTable[num]; !ok {
		// not found
		charTable[num] = '0'
		return append(usedNums, num)
	}
	return usedNums
}

func solvePair(A, B, L, N int, nums []int) string {
	chars := make([]int, 0)
	usedNums := make([]int, 0)
	charTable := make(map[int]byte)

	var num = nums[1]
	var lastNum int
	// calculate second pair

	if qa, ra := divmod(num, A); ra == 0 {
		chars = append(chars, B, A, qa)
		usedNums = tryAddPrime(B, charTable, usedNums)
		usedNums = tryAddPrime(A, charTable, usedNums)
		usedNums = tryAddPrime(qa, charTable, usedNums)
		lastNum = qa
	} else if qb, rb := divmod(num, B); rb == 0 {
		chars = append(chars, A, B, qb)
		usedNums = tryAddPrime(A, charTable, usedNums)
		usedNums = tryAddPrime(B, charTable, usedNums)
		usedNums = tryAddPrime(qb, charTable, usedNums)
		lastNum = qb
	} else {
		return ""
	}

	for i := 2; i < L; i++ {
		num = nums[i]
		if q, r := divmod(num, lastNum); r == 0 {
			if !primeTable[q] {
				return ""
			}
			chars = append(chars, q)
			usedNums = tryAddPrime(q, charTable, usedNums)
			lastNum = q
		} else {
			return ""
		}
	}

	if len(usedNums) != 26 {
		//panic("error")
		return ""
	}

	sort.Sort(UsedNums(usedNums))

	for i, n := range usedNums {
		charTable[n] = 'A' + byte(i)
	}

	var ans bytes.Buffer
	for _, c := range chars {
		ans.WriteByte(charTable[c])
	}

	return ans.String()
}

func main() {
	defer writer.Flush()

	generatePrimes()

	var T int

	scanf("%d\n", &T)

	for caseNo := 0; caseNo < T; caseNo++ {
		var n, l int

		scanf("%d %d\n", &n, &l)

		nums := make([]int, l)

		for j := 0; j < l; j++ {
			scanf("%d\n", &nums[j])
		}

		solve(caseNo+1, n, l, nums)
	}
}
