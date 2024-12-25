package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

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

// 待写
func Send(qwReq QWenRequest) (msgChan chan string, err error) {
	msgChan = make(chan string, 0)
	baseUrl := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"

	byteData, _ := json.Marshal(qwReq)
	buffer := bytes.NewBuffer(byteData)

	request, err := http.NewRequest("POST", baseUrl, buffer)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 构建请求头
	request.Header.Add("Authorization", "Bearer sk-1b516ed0975d4c43aff1cecae74dedfc")
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

func main() {
	//baseUrl := "https://dashscope.aliyuncs.com/api/v1/services/aigc/text-generation/generation"
	//
	//byteData, _ := json.Marshal(QWenRequest{
	//	Model: "qwen-turbo",
	//	Input: Input{
	//		Messages: []Messages{
	//			{
	//				Role:    "system",
	//				Content: "You are a helpful assistant.",
	//			},
	//			{
	//				Role:    "user",
	//				Content: "你是谁？",
	//			},
	//		},
	//	},
	//	//Parameters: Parameter{
	//	//	ResultFormat: "message",
	//	//},
	//})
	//buffer := bytes.NewBuffer(byteData)
	//
	//request, err := http.NewRequest("POST", baseUrl, buffer)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//// 构建请求头
	//request.Header.Add("Authorization", "Bearer sk-1b516ed0975d4c43aff1cecae74dedfc")
	//request.Header.Add("Content-Type", "application/json")
	////request.Header.Set("X-DashScope-SSE", "enable")
	//
	//response, err := http.DefaultClient.Do(request)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//defer response.Body.Close()
	//
	//var qwResponse QWenResponse
	//scan := bufio.NewScanner(response.Body) //分片读取
	//scan.Split(bufio.ScanLines)             //按行读取
	//for scan.Scan() {
	//	text := scan.Text()
	//	err = json.Unmarshal([]byte(text), &qwResponse)
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Printf("Text: \n%s\n", qwResponse.Output.Text)
	//}

	//err = json.NewDecoder(response.Body).Decode(&qwResponse)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Printf("Text: \n%s\n", qwResponse.Output.Text)

	request := QWenRequest{
		Model: "qwen-turbo",
		Input: Input{
			Messages: []Messages{
				{
					Role:    "system",
					Content: "You are a helpful assistant.",
				},
				{
					Role:    "user",
					Content: "你是谁？",
				},
			},
		},
		//Parameters: Parameter{
		//	ResultFormat: "message",
		//},
	}
	msgChan, err := Send(request)
	if err != nil {
		fmt.Println(err)
		return
	}
	for msg := range msgChan {
		fmt.Println(msg)
	}
}
