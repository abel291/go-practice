package canales

import (
	"fmt"
	"time"
)

func Canal4() {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go func() {

		ch1 <- "one"
	}()
	go func() {

		ch2 <- "two"
	}()

	for {
		time.Sleep(time.Millisecond * 100)
		select {
		case msg1 := <-ch1:
			fmt.Println("mesage recibido ", msg1)
		case msg2 := <-ch2:
			fmt.Println("mesage recibido ", msg2)
		default:
			return
		}
	}
}
