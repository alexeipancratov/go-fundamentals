package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func ChannelsBasicDemo() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1) // we have one goroutine to wait for
	go func() {
		// channel operations block until the complementary operation is ready
		ch <- "The message"
	}()

	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()

	// So in the above code Go will know that the second goroutine is waiting for the first one to send a message
	// because they are both using the same channel

	wg.Wait()
}

func ChannelsSelectDemo() {
	// channelsDemo1()
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "message to channel 1"
	}()

	go func() {
		ch2 <- "message to channel 2"
	}()

	// we're giving the task scheduler enough time to recognize there're two
	// different goroutines
	time.Sleep(10 * time.Millisecond)

	// select is blocking unless there's a default option.
	// if both channels are ready then a random one will be selected
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	}
}

func ChannelLoopingThroughMessagesDemo() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)
	}
}
