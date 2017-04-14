package token

import (
	util "GetuiDemo/getui/util"
	"encoding/json"
)

type CloseTokenResult struct {
	Result string `json:"result"`
}

func SetAuthClose(appId string, auth_token string) (*CloseTokenResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/auth_close"
	result, err := util.BytePost(url, auth_token, nil)
	if err != nil {
		return nil, err
	}

	closeTokenResult := new(CloseTokenResult)
	if err := json.Unmarshal([]byte(result), &closeTokenResult); err != nil {
		return nil, err
	}

	return closeTokenResult, err
}
