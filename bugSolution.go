```go
package main

import (

        "fmt"
        "sync"
)

func main() {
        var wg sync.WaitGroup
        ch := make(chan int, 10) // Buffered channel to avoid blocking

        for i := 0; i < 10; i++ {
                wg.Add(1)
                go func(i int) {
                        defer wg.Done()
                        fmt.Printf("goroutine %d: starting\n", i)
                        ch <- i // Send data to the channel
                        fmt.Printf("goroutine %d: finished\n", i)
                }(i)
        }

        go func() {
                for i := 0; i < 10; i++ {
                        val := <-ch
                        fmt.Printf("Main goroutine received: %d\n", val)
                }
        }()

        wg.Wait()
        close(ch)
        fmt.Println("Program finished")
}
```