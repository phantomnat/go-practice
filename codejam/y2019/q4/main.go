package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) { fmt.Fprintf(writer, f, a...) }
func scanf(f string, a ...interface{})  { fmt.Fscanf(reader, f, a...) }

type Data struct {
	Start, End int
	Char       byte
	Input      []byte
	Output     []byte
	Missing    int
}

type Region struct {
	l, m, r            int
	numB, mStart, mEnd int
	// numM match
	numM int
}

type BrokenMachines []int

func (p BrokenMachines) Len() int           { return len(p) }
func (p BrokenMachines) Less(i, j int) bool { return p[i] < p[j] }
func (p BrokenMachines) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func generateData(region Region, bits []byte) {
	for i := region.l; i < region.r; i++ {
		if i < region.m {
			bits[i] = '1'
		} else {
			bits[i] = '0'
		}
	}
}

func calculateResult(region Region, output string, bits []byte, broken []int) ([]Region, []int) {
	//fmt.Printf("region: %v\n", region)
	regions := make([]Region, 0)
	out := output[region.mStart:]
	maxBroken := region.numB
	maxMatch := region.numM
	maxOneMatch := region.m - region.l
	oneCount := 0
	for ; oneCount < maxOneMatch && oneCount < maxMatch && out[oneCount] == '1'; oneCount++ {
	}
	var rl *Region
	var rr *Region

	if oneCount == maxOneMatch {
		// 1 match with the same len in region
		// search 0R region
		mr := region.m + ((region.r - region.m) / 2)
		rr = &Region{region.m, mr, region.r, maxBroken,
			region.mStart + oneCount, region.mEnd, maxMatch - oneCount}
	} else if oneCount == 0 {
		// 0 is match, but 1s do not
		for i := region.l; i < region.m; i++ {
			broken = append(broken, i)
		}

		// search 0R region
		//ml := region.l + ((region.m - region.l) / 2)
		//rl = &Region{region.l, ml, region.m, maxBroken, region.mStart, region.mEnd, 0}
		mr := region.m + ((region.r - region.m) / 2)
		rr = &Region{region.m, mr, region.r, maxBroken - (region.m - region.l),
			region.mStart, region.mEnd, maxMatch - oneCount}
	} else {
		// search 1L region
		ml := region.l + ((region.m - region.l) / 2)
		rl = &Region{region.l, ml, region.m,
			maxOneMatch - oneCount, region.mStart, region.mStart + oneCount,
			oneCount}
		// search 0R region
		// with
		mr := region.m + ((region.r - region.m) / 2)
		rr = &Region{region.m, mr, region.r,
			maxBroken - (maxOneMatch - oneCount), region.mStart + oneCount, region.mEnd,
			maxMatch - oneCount}
	}

	if rl != nil && rl.numM == 0 {
		for i := rl.l; i < rl.r; i++ {
			broken = append(broken, i)
		}
	} else if rl != nil {
		regions = append(regions, *rl)
	}

	if rr != nil && rr.numM == 0 {
		for i := rr.l; i < rr.r; i++ {
			broken = append(broken, i)
		}
	} else if rr != nil {
		regions = append(regions, *rr)
	}

	return regions, broken
}

func main() {
	defer writer.Flush()

	//for {
	//	//N, B, F := 1024, 1, 10
	//	//skips := []int{0, 1, 3}
	//	//N, B, F := 4, len(skips), 10
	//	skips := []int{0, 2, 3}
	//	N, B, F := 4, len(skips), 10
	//
	//	broken := make([]int, 0, 20)
	//	bits := make([]byte, N)
	//	r := Region{0, N / 2, N, B, 0, N - B, N - B}
	//
	//	generateData(r, bits)
	//	fmt.Println(string(bits))
	//
	//	//answer := string([]byte{bits[0], bits[1], bits[4]})
	//	buf := bytes.Buffer{}
	//	for i := 0; i < N; i++ {
	//		isSkip := false
	//		for _, j := range skips {
	//			if i == j {
	//				isSkip = true
	//				break
	//			}
	//		}
	//		if isSkip {
	//			continue
	//		}
	//		buf.WriteByte(bits[i])
	//	}
	//	answer := buf.String()
	//
	//	fmt.Printf("%v\n", answer)
	//
	//	var regions []Region
	//
	//	regions, broken = calculateResult(r, answer, bits, broken)
	//
	//	fmt.Printf("%v\n", regions)
	//
	//	for j := 1; j < F; j++ {
	//
	//		for _, region := range regions {
	//			generateData(region, bits)
	//		}
	//
	//		// try to send N char
	//		fmt.Println(string(bits))
	//
	//		buf := bytes.Buffer{}
	//		for i := 0; i < N; i++ {
	//			isSkip := false
	//			for _, j := range skips {
	//				if i == j {
	//					isSkip = true
	//					break
	//				}
	//			}
	//			if isSkip {
	//				continue
	//			}
	//			buf.WriteByte(bits[i])
	//		}
	//		answer := buf.String()
	//
	//		newRegions := make([]Region, 0, 100)
	//		for _, region := range regions {
	//			var r []Region
	//			r, broken = calculateResult(region, answer, bits, broken)
	//			newRegions = append(newRegions, r...)
	//		}
	//		regions = newRegions
	//	}
	//
	//	// print broken
	//	sort.Sort(BrokenMachines(broken))
	//	var ans bytes.Buffer
	//
	//	ans.WriteString(fmt.Sprintf("%d", broken[0]))
	//	for i := 1; i < len(broken); i++ {
	//		ans.WriteString(fmt.Sprintf(" %d", broken[i]))
	//	}
	//	fmt.Println(ans.String())
	//
	//	break
	//}

	var T int

	scanf("%d\n", &T)

	for caseNo := 0; caseNo < T; caseNo++ {
		var N, B, F int
		var answer string
		scanf("%d %d %d\n", &N, &B, &F)

		broken := make([]int, 0, 20)
		bits := make([]byte, N)
		//mid := N / 2

		r := Region{0, N / 2, N, B, 0, B, N - B}

		generateData(r, bits)

		fmt.Println(string(bits))
		scanf("%s\n", &answer)

		var regions []Region
		regions, broken = calculateResult(r, answer, bits, broken)

		for j := 1; j < F; j++ {

			for _, region := range regions {
				generateData(region, bits)
			}

			// try to send N char
			fmt.Println(string(bits))

			scanf("%s\n", &answer)

			newRegions := make([]Region, 0)
			for _, region := range regions {
				var r []Region
				r, broken = calculateResult(region, answer, bits, broken)
				newRegions = append(newRegions, r...)
			}
			regions = newRegions
		}

		// print broken
		sort.Sort(BrokenMachines(broken))
		brokenstr := make([]string, len(broken))
		for _, b := range broken {
			brokenstr = append(brokenstr, fmt.Sprintf("%d", b))
		}
		fmt.Println(strings.Join(brokenstr, " "))

		var result int
		scanf("%d\n", &result)
		if result != 1 {
			os.Exit(1)
		}
	}
}
