package getui

import (
	"encoding/json"
)

type ToAppResult struct {
	Result string `json:"result"`
	TaskId string `json:"taskid"` //	任务标识号
	Desc   string `json:"desc"`   //	错误信息描述
}

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type ToAppParmar struct {
	Message      *Message      `json:"message"` //消息内容
	Notification *Notification `json:"notification,omitempty"`
	Link         *Link         `json:"link,omitempty"`
	Notypopload  *NotyPopload  `json:"notypopload,omitempty"`
	Transmission *Transmission `json:"transmission,omitempty"`
	Condition    []*Condition  `json:"condition,omitempty"` //	筛选目标用户条件，参考下面的condition说明  可选
	Requestid    string        `json:"requestid"`           //请求唯一标识
}

func ToApp(appId string, auth_token string, toAppParmar *ToAppParmar) (*ToAppResult, error) {

	url := TOKEN_DOMAIN + appId + "/push_app"
	bodyByte, err := GetBody(toAppParmar)
	if err != nil {
		return nil, err
	}

	result, err := BytePost(url, auth_token, bodyByte)
	if err != nil {
		return nil, err
	}

	toAppResult := new(ToAppResult)
	if err := json.Unmarshal([]byte(result), &toAppResult); err != nil {
		return nil, err
	}

	return toAppResult, err
}
