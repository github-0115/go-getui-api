package tool

//open application templates
type Apns struct {
	Alert            *Alert `json:"alert"`                       //
	AutoBadge        string `json:"autoBadge,omitempty"`         //用于计算icon上显示的数字，还可以实现显示数字的自动增减，如“+1”、 “-1”、 “1” 等，计算结果将覆盖badge
	Sound            string `json:"sound,omitempty"`             //通知铃声文件名，无声设置为“com.gexin.ios.silence”
	ContentAvailable int    `json:"content-available,omitempty"` //推送直接带有透传数据
	Category         string `json:"category,omitempty"`          //在客户端通知栏触发特定的action和button显示
}

func GetApns() *Apns {
	apns := &Apns{}
	return apns
}

func (this *Apns) SetAlert(alert *Alert) {
	this.Alert = alert
}

func (this *Apns) SetAutoBadge(str string) {
	this.AutoBadge = str
}

func (this *Apns) SetSound(str string) {
	this.Sound = str
}

func (this *Apns) SetContentAvailable(is int) {
	this.ContentAvailable = is
}

func (this *Apns) SetCategory(str string) {
	this.Category = str
}
