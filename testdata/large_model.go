package main

import (
	"GoRoLingG/core"
	"GoRoLingG/global"
	"fmt"
)

func main() {
	core.InitConfig()

	fmt.Println(global.Config.LargeScaleModel)
}
