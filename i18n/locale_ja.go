//go:build i18n_ja || i18n_all

package i18n

// 注册日语配置
func init() {
	RegisterLocale(Japanese, &Locale{
		Language:     Japanese,
		Region:       "JP",
		Name:         "日本語",
		EnglishName:  "Japanese",
		Messages: map[string]string{
			// 通用消息
			"error":   "エラー",
			"warning": "警告",
			"info":    "情報",
			"success": "成功",
			"failed":  "失敗",
			"loading": "読み込み中...",
			"saving":  "保存中...",
			"done":    "完了",
			"cancel":  "キャンセル",
			"confirm": "確認",
			"yes":     "はい",
			"no":      "いいえ",
			"ok":      "OK",
			
			// 時間相関
			"just_now":      "今",
			"seconds_ago":   "%d秒前",
			"minutes_ago":   "%d分前",
			"hours_ago":     "%d時間前",
			"days_ago":      "%d日前",
			"weeks_ago":     "%d週間前",
			"months_ago":    "%dヶ月前",
			"years_ago":     "%d年前",
			"seconds_later": "%d秒後",
			"minutes_later": "%d分後",
			"hours_later":   "%d時間後",
			"days_later":    "%d日後",
			"weeks_later":   "%d週間後",
			"months_later":  "%dヶ月後",
			"years_later":   "%d年後",
			
			// 検証エラーメッセージ
			"required":        "%sは必須です",
			"email":           "%sは有効なメールアドレスである必要があります",
			"url":             "%sは有効なURLである必要があります",
			"min":             "%sの最小値は%sです",
			"max":             "%sの最大値は%sです",
			"len":             "%sの長さは%s文字である必要があります",
			"mobile":          "%sは有効な携帯電話番号である必要があります",
			"idcard":          "%sは有効な身分証明書番号である必要があります",
			"bankcard":        "%sは有効な銀行カード番号である必要があります",
			"chinese_name":    "%sは有効な中国名である必要があります",
			"strong_password": "%sは強力なパスワードである必要があります（8文字以上、大文字、小文字、数字、特殊文字を含む）",
			
			// ネットワーク関連
			"network_error":     "ネットワークエラー",
			"connection_failed": "接続に失敗しました",
			"timeout":           "タイムアウト",
			"server_error":      "サーバーエラー",
			"not_found":         "見つかりません",
			"unauthorized":      "認証されていません",
			"forbidden":         "アクセス禁止",
			
			// ファイル操作
			"file_not_found":    "ファイルが見つかりません",
			"file_read_error":   "ファイル読み取りエラー",
			"file_write_error":  "ファイル書き込みエラー",
			"file_delete_error": "ファイル削除エラー",
			"file_create_error": "ファイル作成エラー",
			
			// データベース関連
			"db_connection_error": "データベース接続エラー",
			"db_query_error":      "データベースクエリエラー",
			"db_insert_error":     "データベース挿入エラー",
			"db_update_error":     "データベース更新エラー",
			"db_delete_error":     "データベース削除エラー",
			
			// 一般的な操作
			"create": "作成",
			"read":   "読み取り",
			"update": "更新",
			"delete": "削除",
			"list":   "リスト",
			"search": "検索",
			"filter": "フィルタ",
			"sort":   "並び替え",
		},
		Formats: &Formats{
			DateFormat:        "2006年01月02日",
			TimeFormat:        "15:04:05",
			DateTimeFormat:    "2006年01月02日 15:04:05",
			NumberFormat:      "%.2f",
			CurrencyFormat:    "¥%.0f",
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			Units: &Units{
				ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
				SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
				BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},
				TimeUnits: map[string]string{
					"nanosecond":  "ナノ秒",
					"microsecond": "マイクロ秒",
					"millisecond": "ミリ秒",
					"second":      "秒",
					"minute":      "分",
					"hour":        "時間",
					"day":         "日",
					"week":        "週",
					"month":       "月",
					"year":        "年",
					"seconds":     "秒",
					"minutes":     "分",
					"hours":       "時間",
					"days":        "日",
					"weeks":       "週",
					"months":      "月",
					"years":       "年",
				},
				DistanceUnits: []string{"mm", "cm", "m", "km"},
				WeightUnits:   []string{"g", "kg", "t"},
			},
		},
	})
}