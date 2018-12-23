package main

import (
	"bufio"
	"fmt"
	"github.com/dgryski/dgohash"
	"github.com/paulrosania/go-charset/charset"
	_ "github.com/paulrosania/go-charset/data"
	"github.com/spaolacci/murmur3"
	"hash"
	"hash/fnv"
	"io/ioutil"
	"os"
	"strings"
)

var (
	Words = AllRussianWords("litf-win.txt")
)

func convertToUtf8(encoding string, str string) string {
	r, err := charset.NewReader(encoding, strings.NewReader(str))
	if err != nil {
		panic(err)
	}
	result, err := ioutil.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return string(result)
}

func AllRussianWords(filename string) []string {
	var words []string

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := convertToUtf8("windows-1251", scanner.Text())
		word := strings.Split(line, " ")[0]
		words = append(words, word)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return words
}

func findCollisions(hasher hash.Hash32) {
	hashToWords := make(map[uint32][]string)
	for _, word := range Words {
		if _, err := hasher.Write([]byte(word)); err != nil {
			panic(err)
		}
		wordHash := hasher.Sum32()
		hasher.Reset()

		currentWords := hashToWords[wordHash]
		currentWords = append(currentWords, word)
		hashToWords[wordHash] = currentWords
		hasher.Reset()
	}

	for wordHash, words := range hashToWords {
		if len(words) != 1 {
			fmt.Printf("%d: %v\n", wordHash, words)
		}
	}
}

func printWordsAndHashes() {
	for _, word := range Words {
		fmt.Printf("%s: %d\n", word, murmur3.Sum32([]byte(word)))
	}
}

func main() {
	fmt.Println("murmur3 collisions:")
	findCollisions(murmur3.New32())
	fmt.Println("fnv1 collisions:")
	findCollisions(fnv.New32())
	fmt.Println("fnv1a collisions:")
	findCollisions(fnv.New32a())
	fmt.Println("SuperFastHash collisions:")
	findCollisions(dgohash.NewSuperFastHash())

}

func Murmur3() {
}

func Env1() {
}

func Env1a() {
}

func SuperFastHash() {
}
