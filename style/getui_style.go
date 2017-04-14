package style

//open application templates
type GeTuiStyle struct {
	Type        int    `json:"type"`                   //固定为1
	Text        string `json:"text"`                   //通知正文
	Title       string `json:"title"`                  //通知标题
	Logo        string `json:"logo,omitempty"`         //通知的图标名称，包含后缀名（需要在客户端开发时嵌入），如“push.png” 可选
	LogoUrl     string `json:"logourl,omitempty"`      //通知的图标logourl 可选
	IsRing      bool   `json:"is_ring,omitempty"`      //收到通知是否响铃：true响铃，false不响铃。默认响铃  可选
	IsVibrate   bool   `json:"is_vibrate,omitempty"`   //收到通知是否振动：true振动，false不振动。默认振动  可选
	IsClearable bool   `json:"is_clearable,omitempty"` //通知是否可清除： true可清除，false不可清除。默认可清除  可选
}

func GetGeTuiStyle(text string, title string) *GeTuiStyle {
	style := &GeTuiStyle{
		Type:        1,
		Text:        text,
		Title:       title,
		IsRing:      true,
		IsVibrate:   true,
		IsClearable: true,
	}
	return style
}

func (this *GeTuiStyle) SetType(t int) {
	this.Type = t
}

func (this *GeTuiStyle) SetText(str string) {
	this.Text = str
}

func (this *GeTuiStyle) SetTitle(str string) {
	this.Title = str
}

func (this *GeTuiStyle) SetLogo(s string) {
	this.Logo = s
}

func (this *GeTuiStyle) SetLogoUrl(s string) {
	this.LogoUrl = s
}

func (this *GeTuiStyle) SetIsRing(is bool) {
	this.IsRing = is
}

func (this *GeTuiStyle) SetIsVibrate(is bool) {
	this.IsVibrate = is
}

func (this *GeTuiStyle) SetIsClearable(is bool) {
	this.IsClearable = is
}
