package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("pagina.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	f, e := file.Write([]byte("Hola Oscar"))
	if e != nil {
		fmt.Println(f)
		return
	}

}
