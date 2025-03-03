package log_api

import (
	"GoRoLingG/res"
	"github.com/gin-gonic/gin"
)

type OptionResponse struct {
	Label string `json:"label"`
	Value int    `json:"value"`
}

func (LogApi) LogLevelListView(c *gin.Context) {
	res.OKWithData([]OptionResponse{
		{"debug", 1},
		{"info", 2},
		{"warning", 3},
		{"error", 4},
	}, c)
}
