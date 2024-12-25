package large_scale_model_service

import (
	"GoRoLingG/global"
	"GoRoLingG/models"
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type QWenModel struct {
	SessionID uint
}

type Messages struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Input struct {
	Messages []Messages `json:"messages"`
}

//type Parameter struct {
//	ResultFormat string `json:"result_format"`
//}

type QWenRequest struct {
	Model string `json:"model"`
	Input Input  `json:"input"`
	//Parameters Parameter `json:"parameters"`
}

type QWenResponse struct {
	Output struct {
		FinishReason string `json:"finish_reason"`
		Text         string `json:"text"`
	} `json:"output"`
	Usage struct {
		TotalTokens  int `json:"total_tokens"`
		OutputTokens int `json:"output_tokens"`
		InputTokens  int `json:"input_tokens"`
	} `json:"usage"`
	RequestID string `json:"request_id"`
}

func (qwen QWenModel) Send(content string) (msgChan chan string, err error) {
	qwSetting := global.Config.LargeScaleModel.ModelSetting
	msgChan = make(chan string, 0)
	baseUrl := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"
	qwReq := QWenRequest{
		Model: "qwen-turbo",
		Input: Input{},
		//Parameters: Parameter{
		//	ResultFormat: "message",
		//},
	}
	// 查询当前上一次的会话ID是什么，以便查询会话历史记录
	if qwen.SessionID != 0 {
		// 查询会话表，去查询角色对应的设定词
		var sessionModel models.LargeScaleModelSessionModel
		err = global.DB.Preload("RoleModel").Take(&sessionModel, qwen.SessionID).Error
		if err != nil {
			return nil, errors.New("不存在该会话")
		}
		qwReq.Input.Messages = append(qwReq.Input.Messages, Messages{
			Role:    "system",
			Content: sessionModel.RoleModel.Prompt,
		})

		// 加历史记录
		var chatList []models.LargeScaleModelChatModel
		global.DB.Find(&chatList, "session_id = ?", qwen.SessionID)
		for _, chat := range chatList {
			qwReq.Input.Messages = append(qwReq.Input.Messages,
				Messages{
					Role:    "user",
					Content: chat.UserContent,
				},
				Messages{
					Role:    "assistant",
					Content: chat.AIContent,
				})
		}
	}

	var firstChat models.LargeScaleModelChatModel
	global.DB.Where("session_id = ? AND status = ?", qwen.SessionID, true).Order("create_at ASC").First(&firstChat)
	if firstChat.ID != 0 {
		// 检查是否找到了记录
		aiContent := firstChat.AIContent
		global.DB.Model(&models.LargeScaleModelSessionModel{}).Where("id = ?", qwen.SessionID).Update("session_name", aiContent)
	}

	// 加用户记录
	qwReq.Input.Messages = append(qwReq.Input.Messages, Messages{
		Role:    "user",
		Content: content,
	})

	byteData, _ := json.Marshal(qwReq)
	buffer := bytes.NewBuffer(byteData)

	request, err := http.NewRequest("POST", baseUrl, buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 构建请求头
	request.Header.Add("Authorization", "Bearer "+qwSetting.ApiKey)
	request.Header.Add("Content-Type", "application/json")
	//request.Header.Set("X-DashScope-SSE", "enable")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()

	var qwResponse QWenResponse
	scan := bufio.NewScanner(response.Body) //分片读取
	scan.Split(bufio.ScanLines)             //按行读取
	go func() {
		for scan.Scan() {
			text := scan.Text()
			err = json.Unmarshal([]byte(text), &qwResponse)
			if err != nil {
				fmt.Println(err)
				continue
			}
			msgChan <- qwResponse.Output.Text
			// 这个判断新街口可能不用写
			if qwResponse.Output.FinishReason == "stop" {
				close(msgChan)
			}
		}
	}()
	return
}
