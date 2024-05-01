
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("file2.txt")
	if err != nil {
		fmt.Println("Faylni ochib bo'lmadi:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Fields(line)

		wordCounts := make(map[string]int)
		for _, word := range words {
			wordCounts[word]++
		}

		for word, count := range wordCounts {
			fmt.Printf("%s: %d marta\n", word, count)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Xatolik:", err)
		return
	}
}
