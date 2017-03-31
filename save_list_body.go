package getui

import (
	"encoding/json"
)

type SaveListBodyResult struct {
	Result string `json:"result"`
	TaskId string `json:"taskid"` //	任务标识号
	Desc   string `json:"desc"`   //	错误信息描述
}

//消息应用模板 notification、link、notypopload、transmission 四种类型选其一该属性与message下面的msgtype一致
type SaveListBodyParmar struct {
	Message      *Message      `json:"message"` //消息内容
	Notification *Notification `json:"notification,omitempty"`
	Link         *Link         `json:"link,omitempty"`
	Notypopload  *NotyPopload  `json:"notypopload,omitempty"`
	Transmission *Transmission `json:"transmission,omitempty"`
	TaskName     string        `json:"task_name,omitempty"` //	任务名称 可以给多个任务指定相同的task_name，后面用task_name查询推送结果能得到多个任务的结果  可选
}

func SaveListBody(appId string, auth_token string, parmar *SaveListBodyParmar) (*SaveListBodyResult, error) {

	url := TOKEN_DOMAIN + appId + "/save_list_body"
	bodyByte, err := GetSaveListBodyBody(parmar)
	if err != nil {
		return nil, err
	}

	result, err := BytePost(url, auth_token, bodyByte)
	if err != nil {
		return nil, err
	}

	saveListBodyResult := new(SaveListBodyResult)
	if err := json.Unmarshal([]byte(result), &saveListBodyResult); err != nil {
		return nil, err
	}

	return saveListBodyResult, err
}

func GetSaveListBodyBody(parmar *SaveListBodyParmar) ([]byte, error) {

	body, err := json.Marshal(parmar)
	if err != nil {
		return nil, err
	}

	return body, nil
}
