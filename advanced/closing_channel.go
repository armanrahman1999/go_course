package advanced

import "fmt"

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go producer(ch1)
	go filter(ch1, ch2)

	for i := range ch2 {
		fmt.Println(i)
	}

}

func producer(ch chan<- int) {
	for i := range 5 {
		ch <- i
	}
	close(ch)

}

func filter(in <-chan int, out chan<- int) {
	for val := range in {
		if val%2 == 0 {
			out <- val
		}
	}
	close(out)
}

// func main() {

// 	ch := make(chan int)

// 	close(ch)

// 	val, ok := <-ch

// 	if !ok {
// 		fmt.Println("The channel is closed...", val)
// 		return
// 	}
// 	fmt.Println("Channel is not clsoed")

// }