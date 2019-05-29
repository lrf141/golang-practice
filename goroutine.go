package main

import (
    "fmt"
    "sync"
)

func main() {

    count := 0

    var wg sync.WaitGroup
    var mutex sync.Mutex

    for i := 0; i < 10; i++ {
        wg.Add(1)
        mutex.Lock()
        go func() {
            defer wg.Done()
            defer mutex.Unlock()
            count++
        }()

        wg.Add(1)
        mutex.Lock()
        go func() {
            defer wg.Done()
            defer mutex.Unlock()
            fmt.Println(count)
        }()
        wg.Wait()
    }

}