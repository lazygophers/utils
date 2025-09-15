//go:build i18n_ar || i18n_all

package i18n

// 注册阿拉伯语配置
func init() {
	RegisterLocale(Arabic, &Locale{
		Language:     Arabic,
		Region:       "SA",
		Name:         "العربية",
		EnglishName:  "Arabic",
		Messages: map[string]string{
			// رسائل عامة
			"error":   "خطأ",
			"warning": "تحذير",
			"info":    "معلومات",
			"success": "نجح",
			"failed":  "فشل",
			"loading": "جاري التحميل...",
			"saving":  "جاري الحفظ...",
			"done":    "تم",
			"cancel":  "إلغاء",
			"confirm": "تأكيد",
			"yes":     "نعم",
			"no":      "لا",
			"ok":      "موافق",
			
			// الوقت النسبي
			"just_now":      "الآن",
			"seconds_ago":   "منذ %d ثواني",
			"minutes_ago":   "منذ %d دقائق",
			"hours_ago":     "منذ %d ساعات",
			"days_ago":      "منذ %d أيام",
			"weeks_ago":     "منذ %d أسابيع",
			"months_ago":    "منذ %d شهور",
			"years_ago":     "منذ %d سنوات",
			"seconds_later": "بعد %d ثواني",
			"minutes_later": "بعد %d دقائق",
			"hours_later":   "بعد %d ساعات",
			"days_later":    "بعد %d أيام",
			"weeks_later":   "بعد %d أسابيع",
			"months_later":  "بعد %d شهور",
			"years_later":   "بعد %d سنوات",
			
			// رسائل خطأ التحقق
			"required":        "%s مطلوب",
			"email":           "%s يجب أن يكون عنوان بريد إلكتروني صحيح",
			"url":             "%s يجب أن يكون رابط صحيح",
			"min":             "الحد الأدنى لـ %s هو %s",
			"max":             "الحد الأقصى لـ %s هو %s",
			"len":             "طول %s يجب أن يكون %s حرف",
			"mobile":          "%s يجب أن يكون رقم هاتف محمول صحيح",
			"idcard":          "%s يجب أن يكون رقم هوية صحيح",
			"bankcard":        "%s يجب أن يكون رقم بطاقة بنكية صحيح",
			"chinese_name":    "%s يجب أن يكون اسم صيني صحيح",
			"strong_password": "%s يجب أن تكون كلمة مرور قوية (8 أحرف على الأقل، تحتوي على أحرف كبيرة وصغيرة وأرقام ورموز خاصة)",
			
			// الشبكة
			"network_error":     "خطأ في الشبكة",
			"connection_failed": "فشل الاتصال",
			"timeout":           "انتهت المهلة الزمنية",
			"server_error":      "خطأ في الخادم",
			"not_found":         "غير موجود",
			"unauthorized":      "غير مخول",
			"forbidden":         "ممنوع الوصول",
			
			// عمليات الملفات
			"file_not_found":    "الملف غير موجود",
			"file_read_error":   "خطأ في قراءة الملف",
			"file_write_error":  "خطأ في كتابة الملف",
			"file_delete_error": "خطأ في حذف الملف",
			"file_create_error": "خطأ في إنشاء الملف",
			
			// قاعدة البيانات
			"db_connection_error": "خطأ اتصال قاعدة البيانات",
			"db_query_error":      "خطأ استعلام قاعدة البيانات",
			"db_insert_error":     "خطأ إدراج قاعدة البيانات",
			"db_update_error":     "خطأ تحديث قاعدة البيانات",
			"db_delete_error":     "خطأ حذف قاعدة البيانات",
			
			// العمليات العامة
			"create": "إنشاء",
			"read":   "قراءة",
			"update": "تحديث",
			"delete": "حذف",
			"list":   "قائمة",
			"search": "بحث",
			"filter": "تصفية",
			"sort":   "ترتيب",
		},
		Formats: &Formats{
			DateFormat:        "2006/01/02",
			TimeFormat:        "15:04:05",
			DateTimeFormat:    "2006/01/02 15:04:05",
			NumberFormat:      "%.2f",
			CurrencyFormat:    "%.2f ر.س",
			DecimalSeparator:  ".",
			ThousandSeparator: ",",
			Units: &Units{
				ByteUnits:     []string{"ب", "ك.ب", "م.ب", "ج.ب", "ت.ب", "ب.ب"},
				SpeedUnits:    []string{"ب/ث", "ك.ب/ث", "م.ب/ث", "ج.ب/ث", "ت.ب/ث", "ب.ب/ث"},
				BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},
				TimeUnits: map[string]string{
					"nanosecond":  "ن.ث",
					"microsecond": "μث",
					"millisecond": "م.ث",
					"second":      "ثانية",
					"minute":      "دقيقة",
					"hour":        "ساعة",
					"day":         "يوم",
					"week":        "أسبوع",
					"month":       "شهر",
					"year":        "سنة",
					"seconds":     "ثواني",
					"minutes":     "دقائق",
					"hours":       "ساعات",
					"days":        "أيام",
					"weeks":       "أسابيع",
					"months":      "شهور",
					"years":       "سنوات",
				},
				DistanceUnits: []string{"مم", "سم", "م", "كم"},
				WeightUnits:   []string{"ج", "كج", "ط"},
			},
		},
	})
}