package main

import "fmt"

const ebulicaoK = 373.2

func main() {
	K := ebulicaoK
	C := int(K - 273)

	fmt.Printf("Em kelvin a temperatura de ebulição da água é de %.1f e em celsius é de %d\n", K, C)
}
