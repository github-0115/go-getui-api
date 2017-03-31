package getui

const (
	CHARSET                    = "UTF-8"
	CONTENT_TYPE_JSON          = "application/json"
	DEFAULT_CONNECTION_TIMEOUT = 10 //seconds
	TOKEN_DOMAIN               = "https://restapi.getui.com/v1/"
)

var RESULT_MAP = map[string]string{
	"ok":                   "成功",
	"no_msg":               "没有消息体",
	"alias_error":          "找不到别名",
	"black_ip":             "黑名单ip",
	"sign_error":           "鉴权失败",
	"pushnum_overlimit":    "推送次数超限",
	"no_appid":             "找不到appid",
	"no_user":              "找不到对应用户",
	"too_frequent":         "推送过于频繁",
	"sensitive_word":       "有敏感词出现",
	"appid_notmatch":       "appid与cid或者appkey不匹配",
	"not_auth":             "用户没有鉴权",
	"black_appid":          "黑名单app",
	"invalid_param":        "参数检验不通过",
	"alias_notbind":        "别名没有绑定cid",
	"tag_over_limit":       "tag个数超限",
	"successed_online":     "在线下发",
	"successed_offline":    "离线下发",
	"taginvalid_or_noauth": "tag无效或者没有使用权限",
	"no_valid_push":        "没有 有效下发",
	"successed_ignore":     "忽略非活跃用户",
	"no_taskid":            "找不到taskid",
	"other_error":          "其他错误",
}
