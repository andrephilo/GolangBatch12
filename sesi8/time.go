package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("scheduler run")
		time.Sleep(time.Second * 15)
	}
}
