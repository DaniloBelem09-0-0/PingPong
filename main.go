package main

import (
	"fmt"
	ping "goEstudo/ping"
	pong "goEstudo/pong"
	"time"
)

func main() {
	for {
		c1 := make(chan string)
		c2 := make(chan string)
		go ping.Ping(c1)
		go pong.Pong(c2)

		for i := 0; i < 2; i++  {
			select {
			case c := <-c1:
				fmt.Println("Recebido:", c)
				time.Sleep(1 * time.Second)
			case c := <-c2:
				fmt.Println("Recebido:", c)
				time.Sleep(1 * time.Second)
			}
		}
	}
}
