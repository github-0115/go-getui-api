package getui

import (
	"encoding/json"
)

type PushRESResult struct {
	Result     string      `json:"result"`
	TaskId     string      `json:"taskid"`     //任务标识号
	Data       string      `json:"data"`       //查询数据对象
	MsgTotal   int64       `json:"msgTotal"`   //有效可下发总数
	MsgProcess int64       `json:"msgProcess"` //消息回执总数
	ClickNum   int64       `json:"clickNum"`   //用户点击数
	PushNum    int64       `json:"pushNum"`    //im下发总量
	GT         interface{} `json:"msgProcess"` //个推推送结果数据sent 成功下发数feedback 回执数clicked 点击数displayed 展示数
	APN        interface{} `json:"clickNum"`   //iOS推送结果数据，详细字段参考GT
}

type PushRESParmar struct {
	TaskIdList []string `json:"taskIdList"` //查询的任务尖列表
}

func PushResult(appId string, auth_token string, parmar *PushRESParmar) (*PushRESResult, error) {

	url := TOKEN_DOMAIN + appId + "/push_result"
	bodyByte, err := GetBody(parmar)
	if err != nil {
		return nil, err
	}

	result, err := BytePost(url, auth_token, bodyByte)
	if err != nil {
		return nil, err
	}

	pushRESResult := new(PushRESResult)
	if err := json.Unmarshal([]byte(result), &pushRESResult); err != nil {
		return nil, err
	}

	return pushRESResult, err
}
