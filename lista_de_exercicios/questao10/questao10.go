package main

func Somar(a int, b int, c ...int) int {
    total := a + b
    for _, num := range c {
        total += num
    }
    return total
}
