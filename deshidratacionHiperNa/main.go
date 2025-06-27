package main

import "fmt"

const (
	precision  = 2
	asteriscos = "*************Deshidratacion HiperNa********************************"
	DV         = 7
	SF         = 154
)

// variables
var NaR float64
var PN float64
var PA float64

func main() {
	//entrada
	fmt.Print("Peso Nacimiento: ")
	fmt.Scanln(&PN)
	fmt.Print("Peso Actual: ")
	fmt.Scanln(&PA)
	fmt.Print("Na serico: ")
	fmt.Scanln(&NaR)
	//metodos
	DH := (1 - PA/PN) * 100                          //deshidratacion
	AGL := ((0.5 * 24) * PN * (0.75 - DH/100)) / NaR //agua libre
	NB := (80 + (DV * 10)) * (PN / 1000)             //necesidades basales
	DP := PN - PA                                    //deficit previo
	VOL := DP/2 + NB                                 //volumen diario a infundir
	AGLP := AGL / VOL * 100                          //porcentaje de agua libre
	NaXLt := (NaR * (100 - AGLP)) / 100              //sodio por litro
	NaPHP := (VOL * NaXLt) / 1000                    //sodio en la solucion

	fmt.Println("*************metodo a 150 kilo con 1/3 fisio 2/3 dex 5%*************")
	SFDX := (150 * (PN / 1000)) / 3 // me calcula el tercio de 150 por kilo
	DX := SFDX * 2                  //calcula los 2/3 de dextrosa
	GO := (DX + SFDX) / 24          //el goteo de la suma de 1/3 y 2/3
	NaSF := SFDX * 154 / 1000       //calcula la cantidad de sodio en la solucion 1/3
	fmt.Println("imprime cantidades dextrosa,fisio,fisio+dextrosa goteo en 24 horas, y la cantidad de sodio de la solucion...")

	fmt.Println(DX, SFDX, (DX+SFDX)/24, NaSF) //imprime cantidades dextrosa,fisio,fisio+dextrosa goteo en 24 horas, y la cantidad de sodio de la solucion...
	//**********************************************************************

	//salida
	fmt.Println(asteriscos)
	fmt.Printf("Deshidratacion %.*f %%\n", precision, DH)
	fmt.Printf("Agua libre %.*f ml.\n", precision, AGL)
	fmt.Printf("Necesidades basales %.*f ml.\n", precision, NB)
	fmt.Printf("Deficit previo %.*f ml.\n", precision, DP)
	fmt.Printf("Volumen total %.*f ml.\n", precision, VOL)
	fmt.Printf("Sodio por litro %.*f meq/litro.\n", precision, NaXLt)
	fmt.Printf("Sodio en PHP %.*f meq \n", precision, NaPHP)
	fmt.Println(asteriscos)
	fmt.Printf("Goteo: %.2f ml/hora. \n", GO)

	var c rune
	for c != 'S' && c != 's' {
		fmt.Print("Salir?(s/n): ")
		fmt.Scanf("%c\n", &c)
	}

	/*
		DH:=1-PA/PN
		DP2:=PN-PA/2
		NB:=80*PN
		AC:=0.75/
		   Deshidratacion hipernatremica
		   *****************************
		   si tiene 160 o menos via enteral, excepto este contraindicado(nec,vomitos,snc,etc)
		   si tiene signos de shock expandir con solucion fisiologica que tiene 154 meq litro
		   en caso de tener el sodio serico expandir y sea de 170 expandir con una solucion con 150 meq
		   por litro (20 menos que el sodio serico)
		   la expansion sera a 10 ml kilo

		   la meta es bajar 0,5 meq kg en 12 horas o sea 0,5 x 12= 6 meq

		   ejemplo
		   Paciente...
		   Edad 8 dias de vida
		   PN 3000 grs.
		   PA 2400 2,4/3=0.8 tiene el 80 % de su peso
		   Na 170
		   Sodio a disminuir 12 meq/l (0,5 por 24)
		   AC -> Agua corporal 0,75 - 0,20 = 0,55 se le resta la deshidratacion 20%

		   Para calcular el agua libre
		   Formula de Molteni
		   ------------------
		   12 meq * peso en grs * 0,55 = 116 ml
		   ---------------------------
		   		170
		   Sodio real - 12 meq de Na a bajar en 24 horas 170-12=158

		   Deficit Previo 3000-2400=600/2 = 300 por dia en dos dias
		   Necesidades basales = 80+10 (por dia hasta el 7mo) por kg = 150/3=450 ml
		   Total 300 de deficit previo por dia + 450 de basales
		   Agua libre 116 ml es el 15% de 750 ml (el total de liquidos)

		   Sodio de la solucion administrar

		   170*(100-15)
		   ------------=143 meq/l
		   	100

		   1000 ml 143 meq/l
		   750 ml 107 meq/l
		   1 ml de ClNa al 20% tiene 3,4 meq / l

		   31 ml de ClNa al 20% en agua destilada
		   750 ml /24 hs
		   goteo a 31 ml hora


	*/

}
