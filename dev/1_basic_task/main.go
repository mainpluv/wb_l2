package main

import (
	"fmt"
	"os"
	"time"

	"github.com/beevik/ntp"
)

func main() {
	// получаем локальное время
	localTime := time.Now()
	// получаем точное время с сервера
	ntpTime, err := ntp.Time("pool.ntp.org")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting exact time: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Local time: ", localTime)
	fmt.Println("Server time: ", ntpTime)
}
