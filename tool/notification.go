package tool

//open application templates
type Notification struct {
	TransmissionType    bool        `json:"transmission_type,omitempty"`    //收到消息是否立即启动应用，true为立即启动，false则广播等待启动，默认是否  可选
	TransmissionContent string      `json:"transmission_content,omitempty"` //透传内容  可选
	DurationBegin       string      `json:"duration_begin,omitempty"`       //设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss  可选
	DurationEnd         string      `json:"duration_end,omitempty"`         //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	Style               interface{} `json:"style"`                          //通知栏消息布局样式(0 系统样式 1 个推样式) 默认为0  可选
}

func GetNotification() *Notification {
	notification := &Notification{
		TransmissionType: true,
	}
	return notification
}

func (this *Notification) SetTransmissionContent(str string) {
	this.TransmissionContent = str
}

func (this *Notification) SetDurationBegin(str string) {
	this.DurationBegin = str
}

func (this *Notification) SetDurationEnd(str string) {
	this.DurationEnd = str
}

func (this *Notification) SetTransmissionType(is bool) {
	this.TransmissionType = is
}

func (this *Notification) SetNotifyStyle(notifyStyle interface{}) {
	this.Style = notifyStyle
}
