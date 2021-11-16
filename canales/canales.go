package canales

import (
	"fmt"
	"math/rand"
)

func producir(ch chan<- int) {
	for i := 0; i < rand.Intn(100); i++ {
		//time.Sleep(time.Second * 1)
		ch <- rand.Intn(1000)
	}
	//close(ch)

}
func operar(ch <-chan int, b <-chan struct{}) {
	for {
		select {
		case i := <-ch:
			fmt.Print("valor:", i, "\n")

		case <-b:
			return

		}
	}

}
func Canales() {
	mi_canal := make(chan int)
	b := make(chan struct{})
	go operar(mi_canal, b)
	producir(mi_canal)
	b <- struct{}{}

}
