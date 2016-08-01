package main

import (
    "fmt"
)

func abs(x float64) float64 {
    if (x > 0) {
        return x
    } else {
        return -x
    }
}

func square(x float64) float64 {
    return x * x
}

func average(x, y float64) float64 {
    return (x + y) / 2
}

func sqrtImprove(guess, x float64) float64 {
    return average(guess, x / guess)
}

func sqrtGoodEnough(guess, x float64) bool {
    return abs(square(guess) - x) < 0.0001
}

func sqrtIter(guess, x float64) float64 {
    if sqrtGoodEnough(guess, x) {
        return guess
    } else {
        return sqrtIter(sqrtImprove(guess, x), x)
    }
}

func sqrt(x float64) float64 {
    return sqrtIter(1, x)
}

func main() {
    fmt.Println(sqrt(2))
    fmt.Println(sqrt(3))
    fmt.Println(sqrt(4))
    fmt.Println(sqrt(30625))
}
