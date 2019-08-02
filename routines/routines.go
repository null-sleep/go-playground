package main

import (
	"fmt"
	"sync"
)

// func main() {
// 	for i := 1; i <= 10; i++ {
// 		go fmt.Println(i)
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	for i := 1; i <= 10; i++ {
// 		go func() {
// 			wg.Add(1)
// 			fmt.Println(i)
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// }

func main() {
	var wg sync.WaitGroup
	var mux sync.Mutex
	sum := 0
	for i := 1; i <= 10; i++ {
		go func(i int) {
			wg.Add(1)
			mux.Lock()
			sum += i
			fmt.Printf("i: %d, sum: %d\n", i, sum)
			mux.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
}
