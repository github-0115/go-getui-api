package getui

//open application templates
type Notification struct {
	Text                string `json:"text"`                           //消息展示正文
	Title               string `json:"title"`                          //	消息展示标题
	Logo                string `json:"logo,omitempty"`                 //		消息展示logo 可选
	LogoUrl             string `json:"logourl,omitempty"`              //	消息展示logourl 可选
	TransmissionType    bool   `json:"transmission_type,omitempty"`    //	收到消息是否立即启动应用，true为立即启动，false则广播等待启动，默认是否  可选
	TransmissionContent string `json:"transmission_content,omitempty"` //透传内容  可选
	IsRing              bool   `json:"is_ring,omitempty"`              //是否声音提示，默认值true  可选
	IsVibrate           bool   `json:"is_vibrate,omitempty"`           //	是否振动提示，默认值true  可选
	IsClearable         bool   `json:"is_clearable,omitempty"`         //是否可消除，默认值true  可选
	DurationBegin       string `json:"duration_begin,omitempty"`       //	设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss  可选
	DurationEnd         string `json:"duration_end,omitempty"`         //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	NotifyStyle         int    `json:"notify_style,omitempty"`         //	通知栏消息布局样式(0 系统样式 1 个推样式) 默认为0  可选
}

func GetNotification(text string, title string) *Notification {
	notification := &Notification{
		Text:        text,
		Title:       title,
		IsRing:      true,
		IsVibrate:   true,
		IsClearable: true,
		NotifyStyle: 0,
	}
	return notification
}

func (this *Notification) SetText(str string) {
	this.Text = str
}

func (this *Notification) SetTitle(str string) {
	this.Title = str
}

func (this *Notification) SetLogo(s string) {
	this.Logo = s
}

func (this *Notification) SetLogoUrl(str string) {
	this.LogoUrl = str
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

func (this *Notification) SetIsRing(is bool) {
	this.IsRing = is
}

func (this *Notification) SetIsVibrate(is bool) {
	this.IsVibrate = is
}

func (this *Notification) SetIsClearable(is bool) {
	this.IsClearable = is
}

func (this *Notification) SetNotifyStyle(notifyStyle int) {
	this.NotifyStyle = notifyStyle
}
