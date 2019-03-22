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

	mem := make(map[int32]int32)
	visits := make(map[int32]bool)

	for i, v := range arr {
		mem[v] = int32(i) + 1
	}
	var swapCount int32
	n := len(arr)
	var ok bool

	for i := 1; i <= n; i++ {
		if _, ok = visits[int32(i)]; ok || arr[i-1] == int32(i) {
			continue
		}
		count := 0
		j := int32(i)
		for {
			if visits[j] {
				break
			}
			count++
			visits[j] = true
			j, _ = mem[j]
		}
		swapCount += int32(count - 1)
	}

	return swapCount
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
