//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(stream Stream, c chan *Tweet, wg *sync.WaitGroup) {
	defer close(c)
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			break
		}
		c <- tweet
	}
	wg.Done()
}

func consumer(c chan *Tweet, wg *sync.WaitGroup) {
	for t := range c {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
	wg.Done()
}

func main() {
	start := time.Now()
	stream := GetMockStream()
	c := make(chan *Tweet)

	wg := &sync.WaitGroup{}
	wg.Add(2)
	// Producer
	go producer(stream, c, wg)
	// Consumer
	go consumer(c, wg)

	wg.Wait()
	fmt.Printf("Process took %s\n", time.Since(start))
}
