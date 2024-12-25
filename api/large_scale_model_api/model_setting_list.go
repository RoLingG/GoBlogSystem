package large_scale_model_api

import (
	"GoRoLingG/config"
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"GoRoLingG/utils/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
	"path"
)

const docsPath = "upload/docs"

type ModelSetting struct {
	config.ModelSetting
	Help string `json:"help"`
}

// ModelSettingView 获取大模型配置
func (LargeScaleModelApi) ModelSettingView(c *gin.Context) {
	token := c.GetHeader("token")
	var roleID int
	customClaims, err := jwt.ParseToken(token)
	if err == nil && customClaims != nil {
		roleID = customClaims.Role
	}
	if roleID == models.AdminRole {
		// 判断用户是不是管理员，管理员就展示所有信息
		modelSetting := ModelSetting{
			ModelSetting: global.Config.LargeScaleModel.ModelSetting,
		}
		//这里通过后端获取md文件主要还是为了通过读取后端读取文件内容，上传数据提示给前端
		if modelSetting.Name != "" {
			filePath := path.Join(docsPath, fmt.Sprintf("%s.md", modelSetting.Name))
			fileData, err := os.ReadFile(filePath)
			if err == nil {
				modelSetting.Help = string(fileData)
			}
		}

		res.OKWithData(modelSetting, c)
		return
	}

	res.OKWithData(ModelSetting{
		ModelSetting: config.ModelSetting{
			Name:   global.Config.LargeScaleModel.ModelSetting.Name,
			Enable: global.Config.LargeScaleModel.ModelSetting.Enable,
			Title:  global.Config.LargeScaleModel.ModelSetting.Title,
			Slogan: global.Config.LargeScaleModel.ModelSetting.Slogan,
		},
		Help: "",
	}, c)
	return
}
