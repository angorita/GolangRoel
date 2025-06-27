package funciones

import "fmt"

// divide por cero y da error con Errorf de fmt

func Divide(di, de int) (int, error) {
	if de == 0 {
		return 0, fmt.Errorf("no se puede dividir por cero")
	}
	return di / de, nil
}
