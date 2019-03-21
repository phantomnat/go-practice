package main

import (
	"bufio"
	"fmt"
	"io"
	//"os"
	//"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {
	bribes := 0
	n := len(q)
	for i := n - 1; i >= 0; i-- {
		if i-1 >= 0 && int(q[i-1]) == i+1 {
			q[i-1], q[i] = q[i], q[i-1]
			bribes++
		} else if (i-2) >= 0 && int(q[i-2]) == i+1 {
			q[i-2], q[i-1], q[i] = q[i-1], q[i], q[i-2]
			bribes += 2
		} else if int(q[i-1]) != i+1 {
			fmt.Println("Too chaotic")
			return
		}
	}
	fmt.Println(bribes)
}

func main() {
	//reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)
	//
	//tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	//checkError(err)
	//t := int32(tTemp)
	//
	//for tItr := 0; tItr < int(t); tItr++ {
	//	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	//	checkError(err)
	//	n := int32(nTemp)
	//
	//	qTemp := strings.Split(readLine(reader), " ")
	//
	//	var q []int32
	//
	//	for i := 0; i < int(n); i++ {
	//		qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
	//		checkError(err)
	//		qItem := int32(qItemTemp)
	//		q = append(q, qItem)
	//	}
	q := []int32{1, 2, 5, 3, 7, 8, 6, 4}
	minimumBribes(q)
	//}
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
