package scrabble

import (
    "fmt"
)

func Score(input string) int {
    m := make(map[byte]int)
    one := []byte{'A', 'E', 'I', 'O', 'U', 'L', 'N', 'R', 'S', 'T'}
    two := []byte{'D', 'G'}
    three := []byte{'B', 'C', 'M', 'P'}
    four := []byte{'F', 'H', 'V', 'W', 'Y'}
    five := []byte{'K'}
    eight := []byte{'J', 'X'}
    ten := []byte{'Q', 'Z'}
    for _, i := range one {
        m[i] = 1
    }
    for _, i := range two {
        m[i] = 2
    }
    for _, i := range three {
        m[i] = 3
    }
    for _, i := range four {
        m[i] = 4
    }
    for _, i := range five {
        m[i] = 5
    }
    for _, i := range eight {
        m[i] = 8
    }
    for _, i := range ten {
        m[i] = 10
    }
    ret := 0
    for i := 0; i < len(input); i++ {
        v := input[i]
        if input[i] > 90 {
            v -= 32
        }
        val, ok := m[v]
        if ok {
            ret += val
        } else {
            fmt.Println("Error", v)
        }

    }
    return ret
}
