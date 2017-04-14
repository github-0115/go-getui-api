package tool

type Condition struct {
	Key     string `json:"key"`      //筛选条件类型名称(省市region,手机类型phonetype,用户标签tag)
	Values  string `json:"values"`   //筛选参数
	OptType int    `json:"opt_type"` //筛选参数的组合，0:取参数并集or，1：交集and，2：相当与not in {参数1，参数2，....}
}

//key:筛选条件类型名称(省市region,手机类型phonetype,用户标签tag)
//values:筛选参数
//OptType:筛选参数的组合，0:取参数并集or，1：交集and，2：相当与not in {参数1，参数2，....}
func GetCondition(key string, values string, optType int) *Condition {
	condition := &Condition{
		Key:     key,
		Values:  values,
		OptType: optType,
	}
	return condition
}

func (this *Condition) SetKey(key string) {
	this.Key = key
}

func (this *Condition) SetValues(values string) {
	this.Values = values
}

func (this *Condition) SetOptType(optType int) {
	this.OptType = optType
}
