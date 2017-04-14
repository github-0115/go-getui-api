package push

import (
	tool "GetuiDemo/getui/tool"
	util "GetuiDemo/getui/util"
	"encoding/json"
)

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type PushSingleParmar struct {
	Message      *tool.Message      `json:"message"`
	Notification *tool.Notification `json:"notification,omitempty"`
	Link         *tool.Link         `json:"link,omitempty"`
	Notypopload  *tool.NotyPopload  `json:"notypopload,omitempty"`
	Transmission *tool.Transmission `json:"transmission,omitempty"`
	Cid          string             `json:"cid,omitempty"`
	Alias        string             `json:"alias,omitempty"`
	RequestId    string             `json:"requestid"`
}

type PushSingleResult struct {
	Result string `json:"result"` //ok 鉴权成功
	TaskId string `json:"taskid"` //任务标识号
	Desc   string `json:"desc"`   //错误信息描述
	Status string `json:"status"` //推送结果successed_offline 离线下发successed_online 在线下发successed_ignore 非活跃用户不下发
}

func PushSingle(appId string, auth_token string, parmar *PushSingleParmar) (*PushSingleResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/push_single"
	bodyByte, err := util.GetBody(parmar)
	if err != nil {
		return nil, err
	}

	result, err := util.BytePost(url, auth_token, bodyByte)
	if err != nil {
		return nil, err
	}

	pushSingleResult := new(PushSingleResult)
	if err := json.Unmarshal([]byte(result), &pushSingleResult); err != nil {
		return nil, err
	}

	return pushSingleResult, err
}
