package main

import (
	"testing"
)

func BenchmarkReverse(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		array := GenArray(1000)
		for pb.Next() {
			Reverse(array, 10)
		}
	})
}

func BenchmarkReverse2(b *testing.B) {
	b.StopTimer()
	b.StartTimer()
	b.RunParallel(func(pb *testing.PB) {
		array := GenArray(1000)
		for pb.Next() {
			Reverse2(array, 10)
		}
	})
}
