package main

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)

	go func() {
		for {
			select {
			case ch1 <- 1:
				println("ch1 <- 1")
			case ch2 <- 2:
				println("ch2 <- 2")
			case ch3 <- 3:
				println("ch3 <- 3")
			}
		}
	}()
	go func() {
		for {
			select {
			case ch1 <- 4:
				println("ch1 <- 4")
			case ch2 <- 5:
				println("ch2 <- 5")
			case ch3 <- 6:
				println("ch3 <- 6")
			}
		}
	}()
	go func() {
		for {
			select {
			case ch1 <- 7:
				println("ch1 <- 7")
			case ch2 <- 8:
				println("ch2 <- 8")
			case ch3 <- 9:
				println("ch3 <- 9")
			}
		}
	}()
}
