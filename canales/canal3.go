package canales

import (
	"fmt"
	"time"
)

func Canal3() {
	tareas := make(chan int)
	done := make(chan bool)

	go func() {
		for {
			t, ok := <-tareas
			if ok {
				fmt.Println("tarea recibida", t)
			} else {
				fmt.Println("tarea terminadas")
				done <- true
				return
			}
		}
	}()
	for i := 0; i < 15; i++ {
		tareas <- i
		fmt.Println("tarea enviada")
		time.Sleep(time.Millisecond * 400)
	}
	close(tareas)
	fmt.Println("tarea finalizada")
	<-done
}
