package canales

import (
	"fmt"
	"time"
)

func asignar_numero(numero chan<- int) {
	for i := 0; ; i++ {
		numero <- i
	}
}
func calcular_cuadrado(numero <-chan int, cuadrado chan<- int) {
	for i := 0; ; i++ {
		x := <-numero
		cuadrado <- x * x
	}
}
func Canales2() {
	numero := make(chan int)
	cuadrado := make(chan int)
	go asignar_numero(numero)
	go calcular_cuadrado(numero, cuadrado)
	for {
		fmt.Println("cuadrado: ", <-cuadrado)
		time.Sleep(time.Millisecond * 200)
	}
}
