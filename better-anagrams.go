package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strings"
)

var sortedDictionary map[string]string

func main() {
    sortedDictionary = make(map[string]string)
    loadDictionary("words.txt")
    
    words := [3]string{"BASHLACK", "MUSTRANT", "NOSYSHIP"}
    
    for _, word := range words {
        printIfWord(word)
    }
}

func loadDictionary(path string) {
    file, _ := os.Open(path)
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        sortedDictionary[normalize(scanner.Text())] = scanner.Text()
    }
}

func printIfWord(maybeWord string) {
    if word, found := sortedDictionary[normalize(maybeWord)]; found {
        fmt.Println(word)
    }
}

func normalize(word string) string {
    chars := sort.StringSlice(strings.Split(word, ""))
    return strings.ToLower(strings.Join(chars, ""))
}