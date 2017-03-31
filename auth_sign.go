package getui

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

type Parmar struct {
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
	AppKey    string `json:"appkey"`
}

type TokenResult struct {
	Result    string `json:"result"`
	AuthToken string `json:"auth_token"`
}

func GetAuthSign(appId string, appKey string, masterSecret string) (*TokenResult, error) {

	bodyByte, err := GetPostBody(appKey, masterSecret)
	if err != nil {
		return nil, err
	}

	url := TOKEN_DOMAIN + appId + "/auth_sign"
	result, err := BytePost(url, "", bodyByte)
	if err != nil {
		return nil, err
	}

	tokenResult := new(TokenResult)
	if err := json.Unmarshal([]byte(result), &tokenResult); err != nil {
		return nil, err
	}

	return tokenResult, err
}

func GetPostBody(appKey string, masterSecret string) ([]byte, error) {

	signStr, timestamp := HmacSha256(appKey, masterSecret)

	parmar := &Parmar{
		Sign:      signStr,
		Timestamp: timestamp,
		AppKey:    appKey,
	}

	body, err := json.Marshal(parmar)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func HmacSha256(appKey string, masterSecret string) (string, string) {
	timestamp := strconv.FormatInt((time.Now().UnixNano() / 1000000), 10) //签名开始生成毫秒时间
	original := appKey + timestamp + masterSecret

	hash := sha256.New()
	hash.Write([]byte(original))
	sum := hash.Sum(nil)

	return fmt.Sprintf("%x", sum), timestamp
}
