package large_scale_model_service

import (
	"GoRoLingG/global"
	"errors"
)

type LargeScaleModelInterface interface {
	Send(content string) (msgChan chan string, err error)
}

func Send(sessionID uint, content string) (msgChan chan string, err error) {
	var service LargeScaleModelInterface
	switch global.Config.LargeScaleModel.ModelSetting.Name {
	case "qwen":
		service = QWenModel{SessionID: sessionID}
	//case "wenxin":
	//	service = WenXinModel{}
	//case "xinghuo":
	//	service = XinHuoModel{}
	default:
		return nil, errors.New("不支持该大模型")
	}
	return service.Send(content)
}
