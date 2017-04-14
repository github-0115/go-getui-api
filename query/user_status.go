package query

import (
	util "GetuiDemo/getui/util"
	"encoding/json"
)

type UserStatusResult struct {
	Result    string `json:"result"`    //操作结果 无效用户返回no_user
	Cid       string `json:"cid"`       //查询状态的用户cid
	LastLogin string `json:"lastlogin"` //	sdk上次登陆时间戳
	Status    string `json:"status"`    //用户状态 online在线 offline离线
}

func UserStatus(appId string, auth_token string, cid string) (*UserStatusResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/user_status/" + cid

	result, err := util.Get(url, auth_token)
	if err != nil {
		return nil, err
	}

	userStatusResult := new(UserStatusResult)
	if err := json.Unmarshal([]byte(result), &userStatusResult); err != nil {
		return nil, err
	}

	return userStatusResult, err
}
