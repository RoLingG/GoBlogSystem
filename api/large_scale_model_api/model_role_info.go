package large_scale_model_api

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type TagResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Color string `json:"color"`
}

type RoleDetailResponse struct {
	models.Model
	Name      string        `json:"name"`
	Abstract  string        `json:"abstract"`
	Tags      []TagResponse `json:"tags"`
	ChatCount int64         `json:"chat_count"`
	Icon      string        `json:"icon"`
}

func (LargeScaleModelApi) ModelRoleInfoView(c *gin.Context) {
	var cr models.IDRequest
	err := c.ShouldBindUri(&cr)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	var roleModel models.LargeScaleModelRoleModel
	err = global.DB.Preload("Tags").Take(&roleModel, cr.ID).Error
	if err != nil {
		res.FailWithMsg("大模型角色不存在", c)
		return
	}

	var tagList = make([]TagResponse, 0)
	for _, tag := range roleModel.Tags {
		tagList = append(tagList, TagResponse{
			ID:    tag.ID,
			Title: tag.RoleTitle,
			Color: tag.Color,
		})
	}
	var response RoleDetailResponse
	response = RoleDetailResponse{
		Model:    roleModel.Model,
		Name:     roleModel.Name,
		Abstract: roleModel.Abstract,
		Icon:     roleModel.Icon,
		Tags:     tagList,
	}
	global.DB.Model(models.LargeScaleModelChatModel{}).Where("role_id = ?", cr.ID).
		Select("count(id)").Scan(&response.ChatCount)

	res.OKWithData(response, c)
}
