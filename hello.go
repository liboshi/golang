package main

import "fmt"

func Sqrt(x float64) float64 {
        z := 1.0
        for i := 0; i < 1000; i++ {
                z -= (z * z - x) / (2 * z)
        }
        return z
}

func main() {
        var a string = "Hello"
        var b string = "world.\n"
        s := a + " " + b
        fmt.Printf("Hello world.\n")
        fmt.Printf(s)
        for i := 0; i < 10; i++ {
                if i > 5 {
                        break
                }
                println(i)
        }
        list := [] int {1, 2, 3, 4, 5}
        for k, v := range list {
                println(k)
                println(v)
        }
        j := 1
        switch j {
                case 0:
                case 1:
                        println("Hey, j is here...")
                default:
                        break
        }
        for i := 0; i < 5; i++ {
                defer fmt.Printf("%d ", i)
        }
        test := func() {
        
        func callback(y int, f func(int)) {
                f(y)
        }
                println("Anonymous function...")
        }
        test()
}
