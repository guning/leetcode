package main

import "fmt"

var morse = map[rune]string{
    'a': ".-",
    'b': "-...",
    'c': "-.-.",
    'd': "-..",
    'e': ".",
    'f': "..-.",
    'g': "--.",
    'h': "....",
    'i': "..",
    'j': ".---",
    'k': "-.-",
    'l': ".-..",
    'm': "--",
    'n': "-.",
    'o': "---",
    'p': ".--.",
    'q': "--.-",
    'r': ".-.",
    's': "...",
    't': "-",
    'u': "..-",
    'v': "...-",
    'w': ".--",
    'x': "-..-",
    'y': "-.--",
    'z': "--..",
}


func uniqueMorseRepresentations(words []string) int {
    diffTrans := make(map[string]int)
    for _, v := range words {
        diffTrans[transformer(v)] = 1
    }
    fmt.Println("sada", diffTrans)
    return len(diffTrans)
}

func transformer(word string) string {
    morseStr := ""
    for _, v := range []rune(word) {
        morseStr = morseStr + morse[v]
    }
    return morseStr
}

func main() {
    words := []string{"gin", "zen", "gig", "msg"}
    num := uniqueMorseRepresentations(words)
    fmt.Println("sdad", num)
}
