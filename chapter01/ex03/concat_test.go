package main

import (
	"strings"
	"testing"
)


func BenchmarkConcatJoin(b *testing.B) {
	data := [] string{"abc","test","hello","world","123","positive","negative","aiueo"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		strings.Join(data, " ")
	}
}

func BenchmarkConcatFor(b *testing.B) {
	data := [] string{"abc","test","hello","world","123","positive","negative","aiueo"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var newStr string
		for _, str := range data {
			newStr += str + " "
		}
	}
}
