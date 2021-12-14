package main

import (
	"fmt"
	"treuzedev/geheim/config"
)

func main() {
	fmt.Println(config.Get())
}
