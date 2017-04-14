package tool

//点开通知弹窗下载模板
type NotyPopload struct {
	NotyIcon    string `json:"notyicon"`    //通知栏图标
	NotyTitle   string `json:"notytitle"`   //	通知标题
	NotyContent string `json:"notycontent"` //	通知内容
	//	LogoUrl       string `json:"logourl,omitempty"`        //	通知的网络图标地址 可选
	PopTitle      string      `json:"poptitle"`                 //	弹出框标题
	PopContent    string      `json:"popcontent"`               //弹出框内容
	PopImage      string      `json:"popimage"`                 //弹出框图标
	PopButton1    string      `json:"PopButton1"`               //弹出框左边按钮名称
	PopButton2    string      `json:"PopButton2"`               //弹出框左边按钮名称
	LoadUrl       string      `json:"loadurl"`                  //	下载文件地址
	LoadIcon      string      `json:"loadicon,omitempty"`       // 下载图标  可选
	LoadTitle     string      `json:"loadtitle,omitempty"`      //	下载标题  可选
	IsAutoinstall bool        `json:"is_autoinstall,omitempty"` //	是否自动安装，默认值false 可选
	IsActived     bool        `json:"is_actived,omitempty"`     //安装完成后是否自动启动应用程序，默认值false 可选
	AndroidMark   string      `json:"androidmark,omitempty"`    //安卓标识 可选
	SymbianMark   string      `json:"symbianmark,omitempty"`    //塞班标识 可选
	IphoneMark    string      `json:"iphonemark,omitempty"`     //苹果标志 可选
	DurationBegin string      `json:"duration_begin,omitempty"` //	设定展示开始时间，格式为yyyy-MM-dd HH:mm:ss 可选
	DurationEnd   string      `json:"duration_end,omitempty"`   //设定展示结束时间，格式为yyyy-MM-dd HH:mm:ss  可选
	NotifyStyle   int         `json:"notify_style,omitempty"`   //	通知栏消息布局样式(0 系统样式 1 个推样式) 默认为0  可选
	Style         interface{} `json:"style,omitempty"`          //通知栏消息布局样式(0 系统样式 1 个推样式) 默认为0  可选
}

func GetNotyPopload(notyIcon string, notyTitle string, notyContent string, popTitle string, popContent string, popImage string, popButton1 string, popButton2 string, loadUrl string) *NotyPopload {
	notyPopload := &NotyPopload{
		NotyIcon:      notyIcon,
		NotyTitle:     notyTitle,
		NotyContent:   notyContent,
		PopTitle:      popTitle,
		PopContent:    popContent,
		PopImage:      popImage,
		PopButton1:    popButton1,
		PopButton2:    popButton2,
		LoadUrl:       loadUrl,
		IsActived:     false,
		IsAutoinstall: false,
	}
	return notyPopload
}

func (this *NotyPopload) SetNotyIcon(str string) {
	this.NotyIcon = str
}

func (this *NotyPopload) SetNotyTitle(str string) {
	this.NotyTitle = str
}

func (this *NotyPopload) SetNotyContent(s string) {
	this.NotyContent = s
}

func (this *NotyPopload) SetPopTitle(str string) {
	this.PopTitle = str
}

func (this *NotyPopload) SetPopContent(str string) {
	this.PopContent = str
}

func (this *NotyPopload) SetPopImage(str string) {
	this.PopImage = str
}

func (this *NotyPopload) SetPopButton1(str string) {
	this.PopButton1 = str
}

func (this *NotyPopload) SetPopButton2(str string) {
	this.PopButton2 = str
}

func (this *NotyPopload) SetLoadUrl(str string) {
	this.LoadUrl = str
}

func (this *NotyPopload) SetLoadIcon(str string) {
	this.LoadIcon = str
}

func (this *NotyPopload) SetLoadTitle(str string) {
	this.LoadTitle = str
}

func (this *NotyPopload) SetIsAutoinstall(is bool) {
	this.IsAutoinstall = is
}

func (this *NotyPopload) SetIsActived(is bool) {
	this.IsActived = is
}

func (this *NotyPopload) SetDurationBegin(str string) {
	this.DurationBegin = str
}

func (this *NotyPopload) SetDurationEnd(str string) {
	this.DurationEnd = str
}

func (this *NotyPopload) SetNotifyStyle(notifyStyle interface{}) {
	this.Style = notifyStyle
}
