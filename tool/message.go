package tool

type Message struct {
	AppKey            string `json:"appkey"`                        //	注册应用时生成的appkey
	IsOffline         bool   `json:"is_offline,omitempty"`          //	是否离线推送 可选 默认true
	OfflineExpireTime int    `json:"offline_expire_time,omitempty"` //	消息离线存储有效期，单位：ms 默认24小时
	PushNetWorkType   int    `json:"push_network_type,omitempty"`   //	选择推送消息使用网络类型，0：不限制，1：wifi 默认0
	MsgType           string `json:"msgtype"`                       //	消息应用类型，可选项：notification、link、notypopload、transmission
}

func GetMessage() *Message {
	message := &Message{
		IsOffline:         true,
		OfflineExpireTime: 24 * 60 * 60 * 1000,
		PushNetWorkType:   0,
	}
	return message
}

func (this *Message) SetAppKey(appKey string) {
	this.AppKey = appKey
}

func (this *Message) SetIsOffline(isOffline bool) {
	this.IsOffline = isOffline
}

func (this *Message) SetOfflineExpireTime(offlineExpireTime int) {
	this.OfflineExpireTime = offlineExpireTime
}

func (this *Message) SetPushNetWorkType(pushNetWorkType int) {
	this.PushNetWorkType = pushNetWorkType
}

func (this *Message) SetMsgType(msgType string) {
	this.MsgType = msgType
}
