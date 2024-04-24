package other

import (
	"bytes"
    "encoding/json"
    "errors"
    "io/ioutil"
    "net/http"
    "time"
    "fmt"
)
// RequestOptions 定义了请求的额外选项
type RequestOptions struct {
    Method  string            // 请求方法，如 GET、POST 等
    Headers map[string]string // 自定义头部
    Body    []byte            // 请求体
}
// 发送 post 请求
func FetchAPI_POST_TOKEN(url string, data string, token string) (string, error) {
	options := &RequestOptions{
        Method: "POST",
        Headers: map[string]string{
            "Accept": "application/json",
            "Authorization": "Bearer "+token,
        },
        Body: []byte(data),
    }
	return FetchAPI(url, options)
}
func FetchAPI_POST(url string, data string) (string, error) {
	options := &RequestOptions{
        Method: "POST",
        Headers: map[string]string{
            "Accept": "application/json",
        },
        Body: []byte(data),
    }
	return FetchAPI(url, options)
}
// fetchAPI 根据给定的URL和选项请求API
func FetchAPI(url string, options *RequestOptions) (string, error) {
    // 创建一个带有40秒超时的http.Client
    client := &http.Client{
        Timeout: 40 * time.Second,
    }

    // 创建请求
    req, err := http.NewRequest(options.Method, url, bytes.NewBuffer(options.Body))
    if err != nil {
        return "", err
    }

    // 设置请求头
    for key, value := range options.Headers {
        req.Header.Set(key, value)
    }

    // 发送请求并重试
    var lastErr error
    for retries := 0; retries < 2; retries++ {
        resp, err := client.Do(req)
        if err != nil {
            lastErr = err
            // 如果是超时错误，继续重试，否则跳出循环
            // if err, ok := err.(net.Error); !ok || !err.Timeout() {
             //   break
            // }
            continue
        }

        // 确保在读取完毕后立即关闭resp.Body
        body, err := ioutil.ReadAll(resp.Body)
        resp.Body.Close() // 不使用defer，确保在每次循环结束时关闭
        if err != nil {
            lastErr = err
            continue
        }

        // 检查返回的内容是否为空
        if len(body) == 0 {
            return "", errors.New("response is empty")
        }

        // 校验返回的内容是否为JSON格式
        var js json.RawMessage
        if err := json.Unmarshal(body, &js); err != nil {
            fmt.Println("Body:", string(body))
            return "", errors.New("response is not valid JSON")
        }

        // 成功获取响应，返回结果
        return string(body), nil
    }

    // 如果重试结束仍然失败，返回最后一个错误
    if lastErr != nil {
        return "", lastErr
    }

    return "", errors.New("maximum retries exceeded")
}