package style

//open application templates
type OnlyImageStyle struct {
	Type        int    `json:"type"`                   //固定为4
	Logo        string `json:"logo,omitempty"`         //通知的图标名称，包含后缀名（需要在客户端开发时嵌入），如“push.png” 可选
	BannerUrl   string `json:"banner_url"`             //通过url方式指定动态banner图片作为通知背景图
	IsRing      bool   `json:"is_ring,omitempty"`      //收到通知是否响铃：true响铃，false不响铃。默认响铃  可选
	IsVibrate   bool   `json:"is_vibrate,omitempty"`   //收到通知是否振动：true振动，false不振动。默认振动  可选
	IsClearable bool   `json:"is_clearable,omitempty"` //通知是否可清除： true可清除，false不可清除。默认可清除  可选
}

func GetOnlyImageStyle() *OnlyImageStyle {
	style := &OnlyImageStyle{
		Type:        4,
		IsRing:      true,
		IsVibrate:   true,
		IsClearable: true,
	}
	return style
}

func (this *OnlyImageStyle) SetType(t int) {
	this.Type = t
}

func (this *OnlyImageStyle) SetLogo(s string) {
	this.Logo = s
}

func (this *OnlyImageStyle) SetBannerUrl(s string) {
	this.BannerUrl = s
}

func (this *OnlyImageStyle) SetIsRing(is bool) {
	this.IsRing = is
}

func (this *OnlyImageStyle) SetIsVibrate(is bool) {
	this.IsVibrate = is
}

func (this *OnlyImageStyle) SetIsClearable(is bool) {
	this.IsClearable = is
}
