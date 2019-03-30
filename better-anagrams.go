package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

var sortedDictionary map[string][]string

func main() {
    sortedDictionary = make(map[string][]string)
    loadDictionary("words.txt")
    
    words := os.Args[1:]
    
    for _, word := range words {
        printIfWord(word)
    }
}

func loadDictionary(path string) {
    file, _ := os.Open(path)
    defer file.Close()
    
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        word := scanner.Text()
        sortedDictionary[normalize(word)] = append(sortedDictionary[normalize(word)], word)
    }
}

func printIfWord(maybeWord string) {
    if words, found := sortedDictionary[normalize(maybeWord)]; found {
        for _, word := range words {
            fmt.Println(word)
        }
    }
}

func normalize(word string) string {
    chars := strings.Split(word, "")
    sort.Strings(chars)
    return strings.ToLower(strings.Join(chars, ""))
}