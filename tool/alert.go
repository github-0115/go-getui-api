package tool

type Alert struct {
	Title string `json:"title"` //通知标题
	Body  string `json:"body"`  //通知内容
}

func GetAlert(title string, body string) *Alert {
	alert := &Alert{
		Title: title,
		Body:  body,
	}
	return alert
}

func (this *Alert) SetTitle(str string) {
	this.Title = str
}

func (this *Alert) SetBody(str string) {
	this.Body = str
}
