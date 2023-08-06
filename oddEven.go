package main

func odd(channel chan int) {
	channel <- 1
	channel <- 3
}

func even(channel chan int) {
	channel <- 2
	channel <- 4
}

func oddEven() {
	channel := make(chan int)
	go odd(channel)
	go even(channel)
	println(<-channel)
	println(<-channel)
	println(<-channel)
	println(<-channel)
}
