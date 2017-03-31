package getui

import (
	"encoding/json"
)

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type PushSingleParmar struct {
	Message      *Message      `json:"message"`
	Notification *Notification `json:"notification,omitempty"`
	Link         *Link         `json:"link,omitempty"`
	Notypopload  *NotyPopload  `json:"notypopload,omitempty"`
	Transmission *Transmission `json:"transmission,omitempty"`
	Cid          string        `json:"cid,omitempty"`
	Alias        string        `json:"alias,omitempty"`
	RequestId    string        `json:"requestid"`
}

type PushSingleResult struct {
	Result string `json:"result"` //ok 鉴权成功
	TaskId string `json:"taskid"` //任务标识号
	Desc   string `json:"desc"`   //错误信息描述
	Status string `json:"status"` //推送结果successed_offline 离线下发successed_online 在线下发successed_ignore 非活跃用户不下发
}

func PushSingle(appId string, auth_token string, parmar *PushSingleParmar) (*PushSingleResult, error) {

	url := TOKEN_DOMAIN + appId + "/push_single"
	bodyByte, err := GetPushSingleBody(parmar)
	if err != nil {
		return nil, err
	}

	result, err := BytePost(url, auth_token, bodyByte)
	if err != nil {
		return nil, err
	}

	pushSingleResult := new(PushSingleResult)
	if err := json.Unmarshal([]byte(result), &pushSingleResult); err != nil {
		return nil, err
	}

	return pushSingleResult, err
}

func GetPushSingleBody(parmar *PushSingleParmar) ([]byte, error) {

	body, err := json.Marshal(parmar)
	if err != nil {
		return nil, err
	}

	return body, nil
}
