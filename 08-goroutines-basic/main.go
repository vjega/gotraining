package main

import (
	"fmt"
	"sync"
	"time"
)

func asyncfunc(wg *sync.WaitGroup, name string, timeinms time.Duration) {
	for i := 0; i < 10; i++ {
		time.Sleep(timeinms)
		fmt.Println("Processing", name, ":", i)
	}
	wg.Done()
}

func main() {
	fmt.Println("Welcome to concurrency...")
	var wg sync.WaitGroup
	myslice := []string{"marriages", "childbirth", "divorcestat", "literacyrate"}
	timenow := time.Now()
	for idx, value := range myslice {
		wg.Add(1)
		go asyncfunc(&wg, value, time.Millisecond*time.Duration((idx+1)*500))
	}
	wg.Wait()
	fmt.Println(time.Since(timenow))
}
