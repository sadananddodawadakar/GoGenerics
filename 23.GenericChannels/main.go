package main

import (
	"fmt"
)

func Produce[T any](ch chan T, data []T) {
	for _, v := range data {
		ch <- v
	}
	close(ch)
}

func main() {
	ch := make(chan int)
	go Produce(ch, []int{1, 2, 3, 4, 5})

	for val := range ch {
		fmt.Println(val)
	}
	chs := make(chan string)
	go Produce(chs, []string{"Hi", "Hello", "Go", "Generics"})
	for val := range chs {
		fmt.Println(val)
	}
}
