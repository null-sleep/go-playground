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
	for i := 1; i <= 10; i++ {
		go func(i int) {
			wg.Add(1)
			fmt.Println(i)
			wg.Done()
		}(i)
	}
	wg.Wait()
}
