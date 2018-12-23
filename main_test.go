package main

import (
	"github.com/dgryski/dgohash"
	"github.com/spaolacci/murmur3"
	"hash"
	"hash/fnv"
	"testing"
)

func BenchmarkMurmur3(b *testing.B) {
	benchmarkHashHelper(b, murmur3.New32())
}

func BenchmarkEnv1(b *testing.B) {
	benchmarkHashHelper(b, fnv.New32())
}

func BenchmarkEnv1a(b *testing.B) {
	benchmarkHashHelper(b, fnv.New32a())
}

func BenchmarkSuperFastHash(b *testing.B) {
	benchmarkHashHelper(b, dgohash.NewSuperFastHash())
}

func benchmarkHashHelper(b *testing.B, hasher hash.Hash32) {
	for i := 0; i < b.N; i++ {
		for _, word := range Words {
			hasher.Write([]byte(word))
			hasher.Sum32()
			hasher.Reset()
		}
	}
}
