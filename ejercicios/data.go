package ejercicios

import "fmt"

// encuentra los subíndices de los dos elementos cuya suma es 8 de la matriz [1, 3, 5, 7, 8]
// seran: (0,3 ) Y (1,2).

func FindSum() {
	list := []int{1, 3, 5, 7, 8}
	a := 8
	for i := range list {
		for j := i + 1; j < len(list); j++ {
			if list[i]+list[j] == a {
				fmt.Printf("(%d, %d)", list[i], list[j])
			}
			continue
		}

	}
}
