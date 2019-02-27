package main

import (
	"strconv"
	"testing"
)

func BenchmarkStrconvBool(b *testing.B) {
	//b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseBool("true")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkStrconvInt(b *testing.B) {
	//b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseInt("7182818284", 10, 64)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkStrconvUint(b *testing.B) {
	//b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseUint("7182818284", 10, 64)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkStrconvFloat(b *testing.B) {
	//b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, err := strconv.ParseFloat("3.1415926535", 64)
		if err != nil {
			panic(err)
		}
	}
}
