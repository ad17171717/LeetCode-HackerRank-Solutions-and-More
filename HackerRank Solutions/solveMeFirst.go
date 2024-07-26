package main

import (
        "fmt"
)

func solveMeFirst(a uint32, b uint32) uint32{
        return a + b
}

func main(){
        sum := solveMeFirst(7,3)

        fmt.Print(sum)
}
