package tool

type MultiMedia struct {
	Url      string `json:"url"`       //多媒体资源地址
	Type     int    `json:"type"`      //资源类型（1.图片，2.音频， 3.视频）
	OnlyWifi bool   `json:"only_wifi"` //是否只在wifi环境下加载，如果设置成true,但未使用wifi时，会展示成普通通知
}

func GetMultiMedia(url string, t int, onlyWifi bool) *MultiMedia {
	multiMedia := &MultiMedia{
		Url:      url,
		Type:     t,
		OnlyWifi: onlyWifi,
	}
	return multiMedia
}

func (this *MultiMedia) SetUrl(str string) {
	this.Url = str
}

func (this *MultiMedia) SetType(t int) {
	this.Type = t
}

func (this *MultiMedia) SetOnlyWifi(onlyWifi bool) {
	this.OnlyWifi = onlyWifi
}
