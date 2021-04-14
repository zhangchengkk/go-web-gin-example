package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"time"
)

/*
	request 请求格式化日志输出
*/
func LogOutPutRequest(c *gin.Context, info string) {
	output := time.Now().Format("2006-01-02 15:04:05") + " "
	output += info + " "
	output += c.Request.Method + " "
	output += c.Request.Host + " "
	url, _ := json.Marshal(c.Request.URL)
	output += string(url) + " "
	b, _ := ioutil.ReadAll(c.Request.Body)
	// 将读取的数据重新放入body，否则api处理中就无法读取
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	output += string(b)
	fmt.Fprintln(gin.DefaultWriter, output)
}

