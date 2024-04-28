package main

import (
	"fmt"
	"time"
)

func orChannel(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}
	// создаем канал для передачи сигнала о завершении
	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		// если больше 2 каналов используем рекурсию для объединения
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-orChannel(append(channels[3:], orDone)...):
				// если ни один из первых трех каналов не закрыт вызываем orChannel для оставшихся каналов

			}
		}
	}()
	return orDone
}

// возвращает канал, который закрывается после указанной задержки
func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()
	<-orChannel(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v\n", time.Since(start))
}
