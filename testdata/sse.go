package main

import (
	"github.com/gin-gonic/gin"
	"io"
	"time"
)

func SSEDemoView(c *gin.Context) {
	var msgChan = make(chan int, 1)
	go func() {
		for i := 0; i < 10; i++ {
			msgChan <- i
			time.Sleep(time.Second)
		}
		close(msgChan)
	}()

	c.Stream(func(w io.Writer) bool {
		if s, ok := <-msgChan; ok {
			c.SSEvent("", s)
			return true
		}
		return false
	})

	////直接只用这里，会发现数据不是流式传输过来的，SSE的效果不行，得自己在前面加个流式传输
	//for i := 0; i < 10; i++ {
	//	c.SSEvent("sse", i)
	//	time.Sleep(time.Second)
	//}
}

func main() {
	r := gin.Default()
	//sse尽量用GET方式进行传参，后端好传，前端好拿
	r.GET("/sse", SSEDemoView)
	r.Run(":8081")
}
