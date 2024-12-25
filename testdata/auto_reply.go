package main

import (
	"GoRoLingG/core"
	"GoRoLingG/global"
	"GoRoLingG/models"
	"fmt"
)

func main() {
	core.InitConfig()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()

	reply := models.AutoReplyModel{}
	model := reply.AutoReplyValidView("nihaonihao123132")
	fmt.Println(model)
}
