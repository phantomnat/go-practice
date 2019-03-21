package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strconv"
	"strings"
)

// Complete the rotLeft function below.
func rotLeft(a []int32, d int32) []int32 {

	var output []int32

	output = append(a[d:], a[:d]...)

	return output
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
	//nd := strings.Split(readLine(reader), " ")

	//nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	//checkError(err)
	//n := int32(nTemp)

	//dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	//checkError(err)
	//d := int32(dTemp)

	//aTemp := strings.Split(readLine(reader), " ")
	var d int32
	var a []int32
	//
	//for i := 0; i < int(n); i++ {
	//	aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
	//	checkError(err)
	//	aItem := int32(aItemTemp)
	//	a = append(a, aItem)
	//}
	a = []int32{1, 2, 3, 4, 5}
	d = 4

	result := rotLeft(a, d)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, " ")
		}
	}

	fmt.Fprintf(writer, "\n")

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
