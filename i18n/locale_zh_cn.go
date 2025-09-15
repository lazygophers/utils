//go:build i18n_zh_cn || i18n_all

package i18n

// 注册简体中文配置
func init() {
	RegisterLocale(ChineseSimplified, &Locale{
		Language:     ChineseSimplified,
		Region:       "CN",
		Name:         "简体中文",
		EnglishName:  "Chinese Simplified",
		Messages: map[string]string{
			// 通用消息
			"error":   "错误",
			"warning": "警告",
			"info":    "信息",
			"success": "成功",
			"failed":  "失败",
			"loading": "加载中...",
			"saving":  "保存中...",
			"done":    "完成",
			"cancel":  "取消",
			"confirm": "确认",
			"yes":     "是",
			"no":      "否",
			"ok":      "确定",
			
			// 时间相关
			"just_now":     "刚刚",
			"seconds_ago":  "%d 秒前",
			"minutes_ago":  "%d 分钟前", 
			"hours_ago":    "%d 小时前",
			"days_ago":     "%d 天前",
			"weeks_ago":    "%d 周前",
			"months_ago":   "%d 个月前",
			"years_ago":    "%d 年前",
			"seconds_later": "%d 秒后",
			"minutes_later": "%d 分钟后",
			"hours_later":   "%d 小时后", 
			"days_later":    "%d 天后",
			"weeks_later":   "%d 周后",
			"months_later":  "%d 个月后",
			"years_later":   "%d 年后",
			
			// 验证错误消息
			"required":             "%s不能为空",
			"email":                "%s必须是有效的邮箱地址",
			"url":                  "%s必须是有效的URL",
			"min":                  "%s最小值为%s",
			"max":                  "%s最大值为%s",
			"len":                  "%s长度必须为%s个字符",
			"mobile":               "%s必须是有效的手机号码",
			"idcard":               "%s必须是有效的身份证号码",
			"bankcard":             "%s必须是有效的银行卡号",
			"chinese_name":         "%s必须是有效的中文姓名",
			"strong_password":      "%s必须是强密码（至少8位，包含大写字母、小写字母、数字和特殊字符）",
			
			// 网络相关
			"network_error":     "网络错误",
			"connection_failed": "连接失败",
			"timeout":           "超时",
			"server_error":      "服务器错误",
			"not_found":         "未找到",
			"unauthorized":      "未授权",
			"forbidden":         "禁止访问",
			
			// 文件操作
			"file_not_found":    "文件未找到",
			"file_read_error":   "文件读取错误", 
			"file_write_error":  "文件写入错误",
			"file_delete_error": "文件删除错误",
			"file_create_error": "文件创建错误",
			
			// 数据库相关
			"db_connection_error": "数据库连接错误",
			"db_query_error":      "数据库查询错误",
			"db_insert_error":     "数据库插入错误",
			"db_update_error":     "数据库更新错误",
			"db_delete_error":     "数据库删除错误",
			
			// 通用操作
			"create": "创建",
			"read":   "读取", 
			"update": "更新",
			"delete": "删除",
			"list":   "列表",
			"search": "搜索",
			"filter": "筛选",
			"sort":   "排序",
		},
		Formats: &Formats{
			DateFormat:        "2006年01月02日",
			TimeFormat:        "15:04:05", 
			DateTimeFormat:    "2006年01月02日 15:04:05",
			NumberFormat:      "%.2f",
			CurrencyFormat:    "¥%.2f",
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			Units: &Units{
				ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
				SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
				BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},
				TimeUnits: map[string]string{
					"nanosecond":  "纳秒",
					"microsecond": "微秒", 
					"millisecond": "毫秒",
					"second":      "秒",
					"minute":      "分钟",
					"hour":        "小时",
					"day":         "天",
					"week":        "周",
					"month":       "月",
					"year":        "年",
					"seconds":     "秒",
					"minutes":     "分钟",
					"hours":       "小时",
					"days":        "天",
					"weeks":       "周", 
					"months":      "月",
					"years":       "年",
				},
				DistanceUnits: []string{"毫米", "厘米", "米", "千米"},
				WeightUnits:   []string{"克", "千克", "吨"},
			},
		},
	})
}