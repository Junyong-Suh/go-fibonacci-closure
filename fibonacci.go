package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
    previous := 0
    current := 0
    t := 0
    return func() int {
        if current == 0 {
            previous = 0
            current = 1
        } else {
            t = previous
            previous = current
            current = t + current
        }
        return current
    }
}

func main() {
    f := fibonacci()
    for i := 0; i < 10; i++ {
        fmt.Println(f())
    }
}
