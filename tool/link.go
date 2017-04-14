package tool

//点开通知打开网页模板
type Link struct {
	Url           string      `json:"url,omitempty"`            //	打开网址 可选
	DurationBegin string      `json:"duration_begin,omitempty"` //	设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss  可选
	DurationEnd   string      `json:"duration_end,omitempty"`   //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	Style         interface{} `json:"style,omitempty"`          //	通知栏消息布局样式(0 系统样式 1 个推样式) 默认为0  可选
}

func GetLink(url string) *Link {
	link := &Link{
		Url: url,
	}
	return link
}

func (this *Link) SetUrl(str string) {
	this.Url = str
}

func (this *Link) SetDurationBegin(str string) {
	this.DurationBegin = str
}

func (this *Link) SetDurationEnd(str string) {
	this.DurationEnd = str
}

func (this *Link) SetNotifyStyle(notifyStyle interface{}) {
	this.Style = notifyStyle
}
