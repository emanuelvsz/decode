package main

import (
	"fmt"
)

func main() {
	texto := "Hi!!!"

	binary := TextToBinary(texto)
	fmt.Println(binary)
}

func TextToBinary(texto string) []string {
	binaries := make([]string, 0)
	for _, char := range texto {
		ascii := int(char)
		binary := ""
		for ascii > 0 {
			rest := ascii % 2
			binary = fmt.Sprintf("%d%s", rest, binary)
			ascii = ascii / 2
		}

		binaries = append(binaries, binary)
	}

	return binaries
}
