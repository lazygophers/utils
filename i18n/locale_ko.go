//go:build i18n_ko || i18n_all

package i18n

// 注册韩语配置
func init() {
	RegisterLocale(Korean, &Locale{
		Language:     Korean,
		Region:       "KR",
		Name:         "한국어",
		EnglishName:  "Korean",
		Messages: map[string]string{
			// 통용 메시지
			"error":   "오류",
			"warning": "경고",
			"info":    "정보",
			"success": "성공",
			"failed":  "실패",
			"loading": "로딩 중...",
			"saving":  "저장 중...",
			"done":    "완료",
			"cancel":  "취소",
			"confirm": "확인",
			"yes":     "예",
			"no":      "아니오",
			"ok":      "확인",
			
			// 시간 관련
			"just_now":      "방금",
			"seconds_ago":   "%d초 전",
			"minutes_ago":   "%d분 전",
			"hours_ago":     "%d시간 전",
			"days_ago":      "%d일 전",
			"weeks_ago":     "%d주 전",
			"months_ago":    "%d개월 전",
			"years_ago":     "%d년 전",
			"seconds_later": "%d초 후",
			"minutes_later": "%d분 후",
			"hours_later":   "%d시간 후",
			"days_later":    "%d일 후",
			"weeks_later":   "%d주 후",
			"months_later":  "%d개월 후",
			"years_later":   "%d년 후",
			
			// 검증 오류 메시지
			"required":        "%s은(는) 필수입니다",
			"email":           "%s은(는) 유효한 이메일 주소여야 합니다",
			"url":             "%s은(는) 유효한 URL이어야 합니다",
			"min":             "%s의 최소값은 %s입니다",
			"max":             "%s의 최대값은 %s입니다",
			"len":             "%s의 길이는 %s자여야 합니다",
			"mobile":          "%s은(는) 유효한 휴대폰 번호여야 합니다",
			"idcard":          "%s은(는) 유효한 신분증 번호여야 합니다",
			"bankcard":        "%s은(는) 유효한 은행 카드 번호여야 합니다",
			"chinese_name":    "%s은(는) 유효한 중국 이름이어야 합니다",
			"strong_password": "%s은(는) 강력한 비밀번호여야 합니다(8자 이상, 대문자, 소문자, 숫자, 특수문자 포함)",
			
			// 네트워크 관련
			"network_error":     "네트워크 오류",
			"connection_failed": "연결 실패",
			"timeout":           "시간 초과",
			"server_error":      "서버 오류",
			"not_found":         "찾을 수 없음",
			"unauthorized":      "인증되지 않음",
			"forbidden":         "접근 금지",
			
			// 파일 조작
			"file_not_found":    "파일을 찾을 수 없습니다",
			"file_read_error":   "파일 읽기 오류",
			"file_write_error":  "파일 쓰기 오류",
			"file_delete_error": "파일 삭제 오류",
			"file_create_error": "파일 생성 오류",
			
			// 데이터베이스 관련
			"db_connection_error": "데이터베이스 연결 오류",
			"db_query_error":      "데이터베이스 쿼리 오류",
			"db_insert_error":     "데이터베이스 삽입 오류",
			"db_update_error":     "데이터베이스 업데이트 오류",
			"db_delete_error":     "데이터베이스 삭제 오류",
			
			// 일반적인 조작
			"create": "생성",
			"read":   "읽기",
			"update": "업데이트",
			"delete": "삭제",
			"list":   "목록",
			"search": "검색",
			"filter": "필터",
			"sort":   "정렬",
		},
		Formats: &Formats{
			DateFormat:        "2006년 01월 02일",
			TimeFormat:        "15:04:05",
			DateTimeFormat:    "2006년 01월 02일 15:04:05",
			NumberFormat:      "%.2f",
			CurrencyFormat:    "₩%.0f",
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			Units: &Units{
				ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
				SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
				BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},
				TimeUnits: map[string]string{
					"nanosecond":  "나노초",
					"microsecond": "마이크로초",
					"millisecond": "밀리초",
					"second":      "초",
					"minute":      "분",
					"hour":        "시간",
					"day":         "일",
					"week":        "주",
					"month":       "월",
					"year":        "년",
					"seconds":     "초",
					"minutes":     "분",
					"hours":       "시간",
					"days":        "일",
					"weeks":       "주",
					"months":      "월",
					"years":       "년",
				},
				DistanceUnits: []string{"mm", "cm", "m", "km"},
				WeightUnits:   []string{"g", "kg", "t"},
			},
		},
	})
}