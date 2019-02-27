package main

import (
	"regexp"
	"testing"
)

var testRegexp = `^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]+$`

func BenchmarkRegexpString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, err := regexp.MatchString(testRegexp, "jsmith@example.com")
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkRegexpCompiledString(b *testing.B) {
	r, err := regexp.Compile(testRegexp)
	if err != nil {
		panic(err)
	}
	for n := 0; n < b.N; n++ {
		_ = r.MatchString("jsmith@example.com")
	}
}
