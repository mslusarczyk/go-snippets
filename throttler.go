package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

const limit = 2

func main() {
	log.Println("Starting up..")
	inputArray := []string{"foo", "bar", "baz", "boo", "fee"}
	var wg sync.WaitGroup
	wg.Add(len(inputArray))
	throttler := make(chan int, limit)
	for _, elem := range inputArray {
		throttler <- 1
		go doStuff(elem, &wg, throttler)
	}
	wg.Wait()
	log.Println("Done.")
}

func doStuff(elem string, wg *sync.WaitGroup, throttling chan int) {
	log.Printf("Working with: %s ", elem)
	time.Sleep(time.Duration(rand.Int63n(300)) * time.Second)
	<-throttling
	wg.Done()
	log.Printf("Done with: %s ", elem)
}
