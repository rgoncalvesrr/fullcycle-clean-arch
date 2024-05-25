package main

import (
	"fmt"

	"github.com/rgoncalvesrr/fullcycle-clean-arch/configs"
)

func main() {
	cfg := configs.LoadConfig(".")
	fmt.Println(cfg.DBDriver)
}
