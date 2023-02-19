package main

import (
	"github.com/MrHenri/meuPet/configs"
)

func main() {
	_, err := configs.Load()
	if err != nil {
		panic(err)
	}

}
