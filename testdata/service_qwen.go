package main

import (
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/service/large_scale_model_service"
	"fmt"
)

func main() {
	core.InitConfig()
	global.Log = core.InitLogger()
	msgChan, err := large_scale_model_service.Send("qwen", "给我讲个笑话")

	if err != nil {
		fmt.Println(err)
		return
	}
	for msg := range msgChan {
		fmt.Println(msg)
	}
}
