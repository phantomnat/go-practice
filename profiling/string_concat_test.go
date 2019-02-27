package main

import (
	"bytes"
	"strings"
	"testing"
)

var strLen = 100

func BenchmarkConcatString(b *testing.B) {
	var str string
	i := 0
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		str += "x"
		i++
		if i >= strLen {
			i = 0
			str = ""
		}
	}
}

func BenchmarkConcatBuffer(b *testing.B) {
	var buffer bytes.Buffer
	i := 0
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
		i++
		if i >= strLen {
			i = 0
			buffer = bytes.Buffer{}
		}
	}
}

func BenchmarkConcatBufferWithReset(b *testing.B) {
	var buffer bytes.Buffer
	i := 0
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		buffer.WriteString("x")
		i++
		if i >= strLen {
			i = 0
			buffer.Reset()
		}
	}
}

func BenchmarkConcatBuilder(b *testing.B) {
	var str strings.Builder
	var i = 0
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		str.WriteString("x")
		i++
		if i >= strLen {
			i = 0
			str = strings.Builder{}
		}
	}
}
