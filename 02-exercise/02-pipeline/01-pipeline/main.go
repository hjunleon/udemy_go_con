package main

import (
	"fmt"
	// "sync"
)

// TODO: Build a Pipeline
// generator() -> square() -> print



// generator - convertes a list of integers to a channel
func generator(nums []int, out chan<- int) {
	for _,num := range nums {
		// fmt.Printf("Squaring: %v\n", num)
		out <- num
	}
	close(out)
}

// square - receive on inbound channel
// square the number
// output on outbound channel
func square(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func main() {
	// set up the pipeline
	gen_ch := make(chan int)
	sq_ch := make(chan int)
	num_list := []int{2, 3, 5, 7, 11, 13}
	go generator(num_list,gen_ch)
	go square(gen_ch, sq_ch)
	// run the last stage of pipeline
	// receive the values from square stage
	// print each one, until channel is closed.
	for sq := range sq_ch {
		fmt.Println("Ans: ", sq)
	}
	// close(gen_ch)
	
}
