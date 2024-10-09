package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"os"
)

const chatGPTAPI = "https://api.openai.com/v1/chat/completions"

// ContentCheckResponse API 响应结构
type ContentCheckResponse struct {
	IsSafe bool   `json:"is_safe"`
	Reason string `json:"reason"`
}

func CheckContent(content string) (bool, string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return false, "", errors.New("API key is missing")
	}

	// 创建请求体
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo", // 使用的模型
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": "请检查以下内容是否合规并解释原因：" + content,
			},
		},
	})
	if err != nil {
		return false, "", err
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", chatGPTAPI, bytes.NewBuffer(requestBody))
	if err != nil {
		return false, "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, "", err
	}
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusOK {
		return false, "", errors.New("failed to check content")
	}

	// 解析响应
	var responseBody struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&responseBody); err != nil {
		return false, "", err
	}

	// 假设响应内容是合规性判断
	// 这里你需要根据 ChatGPT 的回复解析合规性
	// 示例：如果返回内容包含“不合规”则返回 false
	if contains(responseBody.Choices[0].Message.Content, "不合规") {
		return false, responseBody.Choices[0].Message.Content, nil
	}

	return true, "", nil
}

// contains 检查字符串中是否包含特定子字符串
func contains(s, substr string) bool {
	return bytes.Contains([]byte(s), []byte(substr))
}
