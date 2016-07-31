package main

import (
    "fmt"
    "math"
)

func sqrt(x float64) float64 {
    guess := float64(1)
    for {
        if math.Abs(guess * guess - x) > 0.0001 {
            guess = ((x / guess) + guess) / 2
        } else {
            return guess
        }
    }
}

func main() {
    fmt.Println(sqrt(2), sqrt(4), sqrt(3), sqrt(30625))
}
