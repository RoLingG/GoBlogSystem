package large_scale_model_api

import (
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"os"
	"path"
)

func (LargeScaleModelApi) ModelRoleIconsView(c *gin.Context) {
	dir, err := os.ReadDir("upload/role_icon")
	if err != nil {
		logrus.Error(err)
		res.FailWithMsg("目录不存在", c)
		return
	}
	var list []models.Options[string]
	for _, entry := range dir {
		key := "/" + path.Join("uploads/role_icons", entry.Name())
		list = append(list, models.Options[string]{
			Label: key,
			Value: key,
		})
	}

	res.OKWithData(list, c)
}
