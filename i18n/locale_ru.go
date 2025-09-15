//go:build i18n_ru || i18n_all

package i18n

// 注册俄语配置
func init() {
	RegisterLocale(Russian, &Locale{
		Language:     Russian,
		Region:       "RU",
		Name:         "Русский",
		EnglishName:  "Russian",
		Messages: map[string]string{
			// Общие сообщения
			"error":   "Ошибка",
			"warning": "Предупреждение",
			"info":    "Информация",
			"success": "Успех",
			"failed":  "Неудача",
			"loading": "Загрузка...",
			"saving":  "Сохранение...",
			"done":    "Готово",
			"cancel":  "Отмена",
			"confirm": "Подтвердить",
			"yes":     "Да",
			"no":      "Нет",
			"ok":      "ОК",
			
			// Относительное время
			"just_now":      "Сейчас",
			"seconds_ago":   "%d секунд назад",
			"minutes_ago":   "%d минут назад",
			"hours_ago":     "%d часов назад",
			"days_ago":      "%d дней назад",
			"weeks_ago":     "%d недель назад",
			"months_ago":    "%d месяцев назад",
			"years_ago":     "%d лет назад",
			"seconds_later": "Через %d секунд",
			"minutes_later": "Через %d минут",
			"hours_later":   "Через %d часов",
			"days_later":    "Через %d дней",
			"weeks_later":   "Через %d недель",
			"months_later":  "Через %d месяцев",
			"years_later":   "Через %d лет",
			
			// Сообщения об ошибках валидации
			"required":        "%s обязательно",
			"email":           "%s должно быть действительным адресом электронной почты",
			"url":             "%s должно быть действительным URL",
			"min":             "Минимальное значение %s равно %s",
			"max":             "Максимальное значение %s равно %s",
			"len":             "Длина %s должна быть %s символов",
			"mobile":          "%s должно быть действительным номером мобильного телефона",
			"idcard":          "%s должно быть действительным номером удостоверения личности",
			"bankcard":        "%s должно быть действительным номером банковской карты",
			"chinese_name":    "%s должно быть действительным китайским именем",
			"strong_password": "%s должно быть надежным паролем (не менее 8 символов, содержащим заглавные и строчные буквы, цифры и специальные символы)",
			
			// Сеть
			"network_error":     "Ошибка сети",
			"connection_failed": "Сбой подключения",
			"timeout":           "Тайм-аут",
			"server_error":      "Ошибка сервера",
			"not_found":         "Не найдено",
			"unauthorized":      "Неавторизован",
			"forbidden":         "Доступ запрещен",
			
			// Операции с файлами
			"file_not_found":    "Файл не найден",
			"file_read_error":   "Ошибка чтения файла",
			"file_write_error":  "Ошибка записи файла",
			"file_delete_error": "Ошибка удаления файла",
			"file_create_error": "Ошибка создания файла",
			
			// База данных
			"db_connection_error": "Ошибка подключения к базе данных",
			"db_query_error":      "Ошибка запроса к базе данных",
			"db_insert_error":     "Ошибка вставки в базу данных",
			"db_update_error":     "Ошибка обновления базы данных",
			"db_delete_error":     "Ошибка удаления из базы данных",
			
			// Общие операции
			"create": "Создать",
			"read":   "Читать",
			"update": "Обновить",
			"delete": "Удалить",
			"list":   "Список",
			"search": "Поиск",
			"filter": "Фильтр",
			"sort":   "Сортировка",
		},
		Formats: &Formats{
			DateFormat:        "02.01.2006",
			TimeFormat:        "15:04:05",
			DateTimeFormat:    "02.01.2006 15:04:05",
			NumberFormat:      "%.2f",
			CurrencyFormat:    "%.2f ₽",
			DecimalSeparator:  ",",
			ThousandSeparator: " ",
			Units: &Units{
				ByteUnits:     []string{"Б", "КБ", "МБ", "ГБ", "ТБ", "ПБ"},
				SpeedUnits:    []string{"Б/с", "КБ/с", "МБ/с", "ГБ/с", "ТБ/с", "ПБ/с"},
				BitSpeedUnits: []string{"бит/с", "Кбит/с", "Мбит/с", "Гбит/с", "Тбит/с", "Пбит/с"},
				TimeUnits: map[string]string{
					"nanosecond":  "нс",
					"microsecond": "мкс",
					"millisecond": "мс",
					"second":      "секунда",
					"minute":      "минута",
					"hour":        "час",
					"day":         "день",
					"week":        "неделя",
					"month":       "месяц",
					"year":        "год",
					"seconds":     "секунды",
					"minutes":     "минуты",
					"hours":       "часы",
					"days":        "дни",
					"weeks":       "недели",
					"months":      "месяцы",
					"years":       "годы",
				},
				DistanceUnits: []string{"мм", "см", "м", "км"},
				WeightUnits:   []string{"г", "кг", "т"},
			},
		},
	})
}