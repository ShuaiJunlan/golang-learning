package channel

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelClose(t *testing.T) {
	var intChannel = make(chan int)
	go func() {
		intChannel <- 10
		close(intChannel)
		fmt.Println("Closing intChannel")
	}()
	time.Sleep(time.Second * 2)
	fmt.Println(<-intChannel)
}

func TestChannelBuffer(t *testing.T) {
	var intChannel = make(chan int, 1)
	go func() {
		intChannel <- 1
		fmt.Println("adding one element")
		intChannel <- 2
		fmt.Println("adding two element")
		close(intChannel) // must be closing
	}()
	time.Sleep(time.Second * 2)
	for {
		if val, ok := <-intChannel; ok {
			fmt.Println(val)
		} else {
			break
		}
	}
}
func TestChannelBufferZeroAndOne(t *testing.T) {
	var intChannelZero = make(chan int)
	fmt.Println("cap", cap(intChannelZero))
	go func() {
		intChannelZero <- 1
		fmt.Println("length", len(intChannelZero))
		close(intChannelZero)
	}()
	time.Sleep(time.Second * 2)
	for {
		if val, ok := <-intChannelZero; ok {
			fmt.Println(val)
		} else {
			break
		}
	}
}
