package tool

//传递消息内容给客户端，客户端自行定义通知栏展现形式
type Transmission struct {
	TransmissionType    bool                   `json:"transmission_type"`        //	收到消息是否立即启动应用，true为立即启动，false则广播等待启动，默认是否  可选
	TransmissionContent string                 `json:"transmission_content"`     //透传内容
	DurationBegin       string                 `json:"duration_begin,omitempty"` //	设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss  可选
	DurationEnd         string                 `json:"duration_end,omitempty"`   //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	PushInfo            map[string]interface{} `json:"push_info,omitempty"`      //APNs消息内容  可选
}

func GetTransmission(content string) *Transmission {
	transmission := &Transmission{
		TransmissionContent: content,
		TransmissionType:    false,
	}
	return transmission
}

func (this *Transmission) SetTransmissionContent(str string) {
	this.TransmissionContent = str
}

func (this *Transmission) SetTransmissionType(is bool) {
	this.TransmissionType = is
}

func (this *Transmission) SetDurationBegin(str string) {
	this.DurationBegin = str
}

func (this *Transmission) SetDurationEnd(str string) {
	this.DurationEnd = str
}

func (this *Transmission) SetPushInfo(info map[string]interface{}) {
	this.PushInfo = info
}
