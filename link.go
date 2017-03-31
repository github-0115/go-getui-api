package getui

//点开通知打开网页模板
type Link struct {
	Text          string `json:"text"`                     //消息展示正文
	Title         string `json:"title"`                    //	消息展示标题
	Logo          string `json:"logo,omitempty"`           //		消息展示logo 可选
	LogoUrl       string `json:"logourl,omitempty"`        //	消息展示logourl 可选
	Url           string `json:"url,omitempty"`            //	打开网址 可选
	IsRing        bool   `json:"is_ring,omitempty"`        //是否声音提示，默认值true  可选
	IsVibrate     bool   `json:"is_vibrate,omitempty"`     //	是否振动提示，默认值true  可选
	IsClearable   bool   `json:"is_clearable,omitempty"`   //是否可消除，默认值true   可选
	DurationBegin string `json:"duration_begin,omitempty"` //	设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss  可选
	DurationEnd   string `json:"duration_end,omitempty"`   //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	NotifyStyle   int    `json:"notify_style,omitempty"`   //	通知栏消息布局样式(0 系统样式 1 个推样式) 默认为0  可选
}

func GetLink(text string, title string, url string) *Link {
	notification := &Link{
		Text:        text,
		Title:       title,
		Url:         url,
		IsRing:      true,
		IsVibrate:   true,
		IsClearable: true,
		NotifyStyle: 0,
	}
	return notification
}

func (this *Link) SetText(str string) {
	this.Text = str
}

func (this *Link) SetTitle(str string) {
	this.Title = str
}

func (this *Link) SetLogo(s string) {
	this.Logo = s
}

func (this *Link) SetLogoUrl(str string) {
	this.LogoUrl = str
}

func (this *Link) SetUrl(str string) {
	this.Url = str
}

func (this *Link) SetIsRing(is bool) {
	this.IsRing = is
}

func (this *Link) SetIsVibrate(is bool) {
	this.IsVibrate = is
}

func (this *Link) SetIsClearable(is bool) {
	this.IsClearable = is
}

func (this *Link) SetDurationBegin(str string) {
	this.DurationBegin = str
}

func (this *Link) SetDurationEnd(str string) {
	this.DurationEnd = str
}

func (this *Link) SetNotifyStyle(notifyStyle int) {
	this.NotifyStyle = notifyStyle
}
