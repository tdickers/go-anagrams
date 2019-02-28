package main

import (
    "github.com/cznic/mathutil"
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

var dictionary map[string]bool

func main() {
    dictionary = make(map[string]bool)
    loadDictionary("words.txt")
    
    words := [3]string{"BASHLACK", "MUSTRANT", "NOSYSHIP"}
    
    for _, word := range words {
        permuteChars(word)
    }
}

func permuteChars(word string) {
    chars := sort.StringSlice(strings.Split(word, ""))
    
    mathutil.PermutationFirst(chars)
    printIfWord(chars)
    
    for mathutil.PermutationNext(chars) {
        printIfWord(chars)
    }
}

func loadDictionary(path string) {
    file, _ := os.Open(path)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        dictionary[scanner.Text()] = true
    }
}

func inDictionary(maybeWord string) bool {
    _, exists := dictionary[strings.ToLower(maybeWord)]
    return exists
}

func printIfWord(permutationChars []string) {
    permutation := strings.ToLower(strings.Join(permutationChars, ""))
    if inDictionary(permutation) {
        fmt.Println(permutation)
    }
}