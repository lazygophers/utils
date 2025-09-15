//go:build i18n_zh_tw || i18n_all

package i18n

// 注册繁体中文配置
func init() {
	RegisterLocale(ChineseTraditional, &Locale{
		Language:     ChineseTraditional,
		Region:       "TW",
		Name:         "繁體中文",
		EnglishName:  "Chinese Traditional",
		Messages: map[string]string{
			// 通用消息
			"error":   "錯誤",
			"warning": "警告",
			"info":    "資訊",
			"success": "成功",
			"failed":  "失敗",
			"loading": "載入中...",
			"saving":  "儲存中...",
			"done":    "完成",
			"cancel":  "取消",
			"confirm": "確認",
			"yes":     "是",
			"no":      "否",
			"ok":      "確定",
			
			// 時間相關
			"just_now":      "剛剛",
			"seconds_ago":   "%d 秒前",
			"minutes_ago":   "%d 分鐘前",
			"hours_ago":     "%d 小時前",
			"days_ago":      "%d 天前",
			"weeks_ago":     "%d 週前",
			"months_ago":    "%d 個月前",
			"years_ago":     "%d 年前",
			"seconds_later": "%d 秒後",
			"minutes_later": "%d 分鐘後",
			"hours_later":   "%d 小時後",
			"days_later":    "%d 天後",
			"weeks_later":   "%d 週後",
			"months_later":  "%d 個月後",
			"years_later":   "%d 年後",
			
			// 驗證錯誤消息
			"required":        "%s不能為空",
			"email":           "%s必須是有效的電子郵件地址",
			"url":             "%s必須是有效的URL",
			"min":             "%s最小值為%s",
			"max":             "%s最大值為%s",
			"len":             "%s長度必須為%s個字符",
			"mobile":          "%s必須是有效的手機號碼",
			"idcard":          "%s必須是有效的身份證號碼",
			"bankcard":        "%s必須是有效的銀行卡號",
			"chinese_name":    "%s必須是有效的中文姓名",
			"strong_password": "%s必須是強密碼（至少8位，包含大寫字母、小寫字母、數字和特殊字符）",
			
			// 網路相關
			"network_error":     "網路錯誤",
			"connection_failed": "連接失敗",
			"timeout":           "逾時",
			"server_error":      "伺服器錯誤",
			"not_found":         "未找到",
			"unauthorized":      "未授權",
			"forbidden":         "禁止存取",
			
			// 檔案操作
			"file_not_found":    "檔案未找到",
			"file_read_error":   "檔案讀取錯誤",
			"file_write_error":  "檔案寫入錯誤",
			"file_delete_error": "檔案刪除錯誤",
			"file_create_error": "檔案建立錯誤",
			
			// 資料庫相關
			"db_connection_error": "資料庫連接錯誤",
			"db_query_error":      "資料庫查詢錯誤",
			"db_insert_error":     "資料庫插入錯誤",
			"db_update_error":     "資料庫更新錯誤",
			"db_delete_error":     "資料庫刪除錯誤",
			
			// 通用操作
			"create": "建立",
			"read":   "讀取",
			"update": "更新",
			"delete": "刪除",
			"list":   "清單",
			"search": "搜尋",
			"filter": "篩選",
			"sort":   "排序",
		},
		Formats: &Formats{
			DateFormat:        "2006年01月02日",
			TimeFormat:        "15:04:05",
			DateTimeFormat:    "2006年01月02日 15:04:05",
			NumberFormat:      "%.2f",
			CurrencyFormat:    "NT$%.2f",
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			Units: &Units{
				ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
				SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
				BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},
				TimeUnits: map[string]string{
					"nanosecond":  "奈秒",
					"microsecond": "微秒",
					"millisecond": "毫秒",
					"second":      "秒",
					"minute":      "分鐘",
					"hour":        "小時",
					"day":         "天",
					"week":        "週",
					"month":       "月",
					"year":        "年",
					"seconds":     "秒",
					"minutes":     "分鐘",
					"hours":       "小時",
					"days":        "天",
					"weeks":       "週",
					"months":      "月",
					"years":       "年",
				},
				DistanceUnits: []string{"毫米", "公分", "公尺", "公里"},
				WeightUnits:   []string{"公克", "公斤", "公噸"},
			},
		},
	})
}