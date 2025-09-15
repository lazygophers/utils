//go:build i18n_de || i18n_all

package i18n

// 注册德语配置
func init() {
	RegisterLocale(German, &Locale{
		Language:     German,
		Region:       "DE",
		Name:         "Deutsch",
		EnglishName:  "German",
		Messages: map[string]string{
			// Allgemeine Nachrichten
			"error":   "Fehler",
			"warning": "Warnung",
			"info":    "Information",
			"success": "Erfolg",
			"failed":  "Fehlgeschlagen",
			"loading": "Lädt...",
			"saving":  "Speichern...",
			"done":    "Fertig",
			"cancel":  "Abbrechen",
			"confirm": "Bestätigen",
			"yes":     "Ja",
			"no":      "Nein",
			"ok":      "OK",
			
			// Relative Zeit
			"just_now":      "Gerade eben",
			"seconds_ago":   "Vor %d Sekunden",
			"minutes_ago":   "Vor %d Minuten",
			"hours_ago":     "Vor %d Stunden",
			"days_ago":      "Vor %d Tagen",
			"weeks_ago":     "Vor %d Wochen",
			"months_ago":    "Vor %d Monaten",
			"years_ago":     "Vor %d Jahren",
			"seconds_later": "In %d Sekunden",
			"minutes_later": "In %d Minuten",
			"hours_later":   "In %d Stunden",
			"days_later":    "In %d Tagen",
			"weeks_later":   "In %d Wochen",
			"months_later":  "In %d Monaten",
			"years_later":   "In %d Jahren",
			
			// Validierungsfehlermeldungen
			"required":        "%s ist erforderlich",
			"email":           "%s muss eine gültige E-Mail-Adresse sein",
			"url":             "%s muss eine gültige URL sein",
			"min":             "Der Mindestwert von %s ist %s",
			"max":             "Der Höchstwert von %s ist %s",
			"len":             "Die Länge von %s muss %s Zeichen betragen",
			"mobile":          "%s muss eine gültige Handynummer sein",
			"idcard":          "%s muss eine gültige Ausweisnummer sein",
			"bankcard":        "%s muss eine gültige Bankkartennummer sein",
			"chinese_name":    "%s muss ein gültiger chinesischer Name sein",
			"strong_password": "%s muss ein starkes Passwort sein (mindestens 8 Zeichen, enthält Groß- und Kleinbuchstaben, Zahlen und Sonderzeichen)",
			
			// Netzwerk
			"network_error":     "Netzwerkfehler",
			"connection_failed": "Verbindung fehlgeschlagen",
			"timeout":           "Zeitüberschreitung",
			"server_error":      "Serverfehler",
			"not_found":         "Nicht gefunden",
			"unauthorized":      "Nicht autorisiert",
			"forbidden":         "Zugang verweigert",
			
			// Dateioperationen
			"file_not_found":    "Datei nicht gefunden",
			"file_read_error":   "Datei-Lesefehler",
			"file_write_error":  "Datei-Schreibfehler",
			"file_delete_error": "Datei-Löschfehler",
			"file_create_error": "Datei-Erstellungsfehler",
			
			// Datenbank
			"db_connection_error": "Datenbankverbindungsfehler",
			"db_query_error":      "Datenbankabfragefehler",
			"db_insert_error":     "Datenbank-Einfügefehler",
			"db_update_error":     "Datenbank-Aktualisierungsfehler",
			"db_delete_error":     "Datenbank-Löschfehler",
			
			// Allgemeine Operationen
			"create": "Erstellen",
			"read":   "Lesen",
			"update": "Aktualisieren",
			"delete": "Löschen",
			"list":   "Liste",
			"search": "Suchen",
			"filter": "Filtern",
			"sort":   "Sortieren",
		},
		Formats: &Formats{
			DateFormat:        "02.01.2006",
			TimeFormat:        "15:04:05",
			DateTimeFormat:    "02.01.2006 15:04:05",
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
					"second":      "Sekunde",
					"minute":      "Minute",
					"hour":        "Stunde",
					"day":         "Tag",
					"week":        "Woche",
					"month":       "Monat",
					"year":        "Jahr",
					"seconds":     "Sekunden",
					"minutes":     "Minuten",
					"hours":       "Stunden",
					"days":        "Tage",
					"weeks":       "Wochen",
					"months":      "Monate",
					"years":       "Jahre",
				},
				DistanceUnits: []string{"mm", "cm", "m", "km"},
				WeightUnits:   []string{"g", "kg", "t"},
			},
		},
	})
}