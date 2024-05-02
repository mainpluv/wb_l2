package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	for n := range ch {
		println(n)
	}
}

// сначала последовательно выведутся цифры от 0 до 9, а после этого будет
// deadlock, так как range будет вытаться читать из канала, в который
// ничего не приходит.
