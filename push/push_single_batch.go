package push

import (
	tool "GetuiDemo/getui/tool"
	util "GetuiDemo/getui/util"
	"encoding/json"
)

type PushSigleBatchResult struct {
	Result string `json:"result"`
	TaskId string `json:"taskid"` //	任务标识号
	Desc   string `json:"desc"`   //	错误信息描述
}

type PushSigleBatchListParmar struct {
	MsgList    []*PushSigleBatchParmar `json:"msg_list"`              //
	NeedDetail bool                    `json:"need_detail,omitempty"` //默认值:false，是否需要返回每个CID的状态
}

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type PushSigleBatchParmar struct {
	Message      *tool.Message      `json:"message"` //消息内容
	Notification *tool.Notification `json:"notification,omitempty"`
	Link         *tool.Link         `json:"link,omitempty"`
	Notypopload  *tool.NotyPopload  `json:"notypopload,omitempty"`
	Transmission *tool.Transmission `json:"transmission,omitempty"`
	Cid          string             `json:"cid"`   //cid为cid list，与alias list二选一
	Alias        string             `json:"alias"` //	alias为alias list，与cid list二选一
	RequestId    string             `json:"requestid"`
}

func PushSigleBatch(appId string, auth_token string, parmar *PushSigleBatchListParmar) (*PushSigleBatchResult, error) {

	url := util.TOKEN_DOMAIN + appId + "/push_single_batch"
	bodyByte, err := util.GetBody(parmar)
	if err != nil {
		return nil, err
	}

	result, err := util.BytePost(url, auth_token, bodyByte)
	if err != nil {
		return nil, err
	}

	pushSigleBatchResult := new(PushSigleBatchResult)
	if err := json.Unmarshal([]byte(result), &pushSigleBatchResult); err != nil {
		return nil, err
	}

	return pushSigleBatchResult, err
}
