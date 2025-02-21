package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

func main() {
	currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error occurred while connecting to the server")
		os.Exit(1)
	} else {
		fmt.Println(currentTime)
	}
}
