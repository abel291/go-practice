package ejercicios

import (
	"fmt"
	"time"
)

//user "github.com/go_pruebas/users"

func Fibonacci(n int) {
	x1 := 1
	x2 := 0

	for i := 0; i < n; i++ {
		xtotal := x1 + x2
		x2 = x1
		x1 = xtotal
		fmt.Printf("\n %d) %d + %d = %d ", i+1, x1, x2, xtotal)
		time.Sleep(200 * time.Millisecond)

	}

}
