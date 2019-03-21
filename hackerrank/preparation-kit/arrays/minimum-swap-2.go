package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strconv"
	"strings"
)

// Complete the minimumSwaps function below.
func minimumSwaps(arr []int32) int32 {
	// find source
	// find target
	//n := swap(arr, len(arr), 0, 0)
	//return int32(n)
	mem := make(map[int32][]int32)
	//for
}

func swap(originalArray []int32, n, startIndex, depth int) int {
	minSwap := 0

	for i := 0; i < n-1; i++ {

		if i+1 != int(originalArray[i]) {
			arr := append(originalArray[:0:0], originalArray...)

			// need to swap
			for j := 0; j < n; j++ {
				if j == i {
					continue
				}
				if i+1 == int(arr[j]) {
					space := ""
					for x := 0; x < depth; x++ {
						space += "  "
					}
					fmt.Printf("%sswap %d - %d\n", space, arr[i], arr[j])
					arr[i], arr[j] = arr[j], arr[i]
					swap := 1 + swap(arr, n, i+1, depth+1)

					if minSwap == 0 || minSwap > 0 && swap < minSwap {
						minSwap = swap
					}
				}
			}
		}
	}

	return minSwap
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	//
	//stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	//checkError(err)
	//
	//defer stdout.Close()
	//
	writer := bufio.NewWriterSize(os.Stdout, 1024*1024)
	//
	//nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	//checkError(err)
	//n := int32(nTemp)
	//
	//arrTemp := strings.Split(readLine(reader), " ")
	//
	//var arr []int32
	//
	//for i := 0; i < int(n); i++ {
	//	arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
	//	checkError(err)
	//	arrItem := int32(arrItemTemp)
	//	arr = append(arr, arrItem)
	//}
	//arr := []int32{4, 3, 1, 2}
	arr := []int32{7, 1, 3, 2, 4, 5, 6}
	res := minimumSwaps(arr)

	fmt.Fprintf(writer, "%d\n", res)

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
