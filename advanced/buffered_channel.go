package main

import "fmt"

// // buffer is a storage

func main() {

	 ch := make(chan int, 4)

    go func() {
        for i := 1; i <= 5; i++ {
            ch <- i
            fmt.Println("Sent:", i)
        }
        close(ch)
    }()

    for v := range ch {
        fmt.Println("Received:", v)
    }


}