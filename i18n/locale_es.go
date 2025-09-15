//go:build i18n_es || i18n_all

package i18n

// 注册西班牙语配置
func init() {
	RegisterLocale(Spanish, &Locale{
		Language:     Spanish,
		Region:       "ES",
		Name:         "Español",
		EnglishName:  "Spanish",
		Messages: map[string]string{
			// Mensajes generales
			"error":   "Error",
			"warning": "Advertencia",
			"info":    "Información",
			"success": "Éxito",
			"failed":  "Falló",
			"loading": "Cargando...",
			"saving":  "Guardando...",
			"done":    "Completado",
			"cancel":  "Cancelar",
			"confirm": "Confirmar",
			"yes":     "Sí",
			"no":      "No",
			"ok":      "OK",
			
			// Tiempo relativo
			"just_now":      "Ahora mismo",
			"seconds_ago":   "Hace %d segundos",
			"minutes_ago":   "Hace %d minutos",
			"hours_ago":     "Hace %d horas",
			"days_ago":      "Hace %d días",
			"weeks_ago":     "Hace %d semanas",
			"months_ago":    "Hace %d meses",
			"years_ago":     "Hace %d años",
			"seconds_later": "En %d segundos",
			"minutes_later": "En %d minutos",
			"hours_later":   "En %d horas",
			"days_later":    "En %d días",
			"weeks_later":   "En %d semanas",
			"months_later":  "En %d meses",
			"years_later":   "En %d años",
			
			// Mensajes de error de validación
			"required":        "%s es requerido",
			"email":           "%s debe ser una dirección de correo electrónico válida",
			"url":             "%s debe ser una URL válida",
			"min":             "El valor mínimo de %s es %s",
			"max":             "El valor máximo de %s es %s",
			"len":             "La longitud de %s debe ser de %s caracteres",
			"mobile":          "%s debe ser un número de teléfono móvil válido",
			"idcard":          "%s debe ser un número de documento de identidad válido",
			"bankcard":        "%s debe ser un número de tarjeta bancaria válido",
			"chinese_name":    "%s debe ser un nombre chino válido",
			"strong_password": "%s debe ser una contraseña fuerte (al menos 8 caracteres, conteniendo mayúsculas, minúsculas, números y caracteres especiales)",
			
			// Red
			"network_error":     "Error de red",
			"connection_failed": "Falló la conexión",
			"timeout":           "Tiempo de espera agotado",
			"server_error":      "Error del servidor",
			"not_found":         "No encontrado",
			"unauthorized":      "No autorizado",
			"forbidden":         "Acceso prohibido",
			
			// Operaciones de archivos
			"file_not_found":    "Archivo no encontrado",
			"file_read_error":   "Error al leer el archivo",
			"file_write_error":  "Error al escribir el archivo",
			"file_delete_error": "Error al eliminar el archivo",
			"file_create_error": "Error al crear el archivo",
			
			// Base de datos
			"db_connection_error": "Error de conexión a la base de datos",
			"db_query_error":      "Error de consulta de la base de datos",
			"db_insert_error":     "Error de inserción en la base de datos",
			"db_update_error":     "Error de actualización de la base de datos",
			"db_delete_error":     "Error de eliminación de la base de datos",
			
			// Operaciones generales
			"create": "Crear",
			"read":   "Leer",
			"update": "Actualizar",
			"delete": "Eliminar",
			"list":   "Lista",
			"search": "Buscar",
			"filter": "Filtrar",
			"sort":   "Ordenar",
		},
		Formats: &Formats{
			DateFormat:        "02/01/2006",
			TimeFormat:        "15:04:05",
			DateTimeFormat:    "02/01/2006 15:04:05",
			NumberFormat:      "%.2f",
			CurrencyFormat:    "%.2f €",
			DecimalSeparator:  ",",
			ThousandSeparator: ".",
			Units: &Units{
				ByteUnits:     []string{"B", "KB", "MB", "GB", "TB", "PB"},
				SpeedUnits:    []string{"B/s", "KB/s", "MB/s", "GB/s", "TB/s", "PB/s"},
				BitSpeedUnits: []string{"bps", "Kbps", "Mbps", "Gbps", "Tbps", "Pbps"},
				TimeUnits: map[string]string{
					"nanosecond":  "ns",
					"microsecond": "μs",
					"millisecond": "ms",
					"second":      "segundo",
					"minute":      "minuto",
					"hour":        "hora",
					"day":         "día",
					"week":        "semana",
					"month":       "mes",
					"year":        "año",
					"seconds":     "segundos",
					"minutes":     "minutos",
					"hours":       "horas",
					"days":        "días",
					"weeks":       "semanas",
					"months":      "meses",
					"years":       "años",
				},
				DistanceUnits: []string{"mm", "cm", "m", "km"},
				WeightUnits:   []string{"g", "kg", "t"},
			},
		},
	})
}