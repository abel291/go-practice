package canales

import (
	"fmt"
	"sync"
	"time"
)

func work(i int, wg *sync.WaitGroup) {
	fmt.Println("tarea recibida")
	time.Sleep(time.Second)
	fmt.Println("tarea terminada")
	defer wg.Done()

}

func Canal5() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go work(i, &wg)
	}
	wg.Wait()
}
