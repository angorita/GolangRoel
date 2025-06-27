package main

import (
	"fmt"
	"log"
)

func divide(dividendo, divisor int) {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
			log.SetPrefix("main()")
			log.Fatal("Se fue a la mierda")
		}
	}()

	fmt.Println(dividendo / divisor)
}

func main() {
	divide(12, 13)
	divide(12, 2)
	divide(12, 3)
	divide(12, 1)
	divide(12, 0)
	divide(12, 1)
	divide(12, 1)
	divide(12, 1)
	divide(12, 1)

}
