package query

import (
	util "GetuiDemo/getui/util"
	"encoding/json"
)

type QueryAppUserResult struct {
	Result           string `json:"result"`           //操作结果 无效用户返回no_user
	Data             string `json:"data"`             //查询数据对象
	AppId            string `json:"appId"`            //	请求的AppId
	Date             string `json:"date"`             //查询的日期（格式：yyyy-MM-dd）
	NewRegistCount   int64  `json:"newRegistCount"`   //	新注册用户数
	RegistTotalCount int64  `json:"registTotalCount"` //累计注册用户数
	ActiveCount      int64  `json:"activeCount"`      //活跃用户数
	OnlineCount      int64  `json:"onlineCount"`      //在线用户数使用方法
}

func QueryAppUser(appId string, auth_token string, timeStr string) (*QueryAppUserResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/query_app_user/" + timeStr //20160404

	result, err := util.Get(url, auth_token)
	if err != nil {
		return nil, err
	}

	queryAppUserResult := new(QueryAppUserResult)
	if err := json.Unmarshal([]byte(result), &queryAppUserResult); err != nil {
		return nil, err
	}

	return queryAppUserResult, err
}
