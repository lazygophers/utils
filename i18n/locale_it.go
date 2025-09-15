//go:build i18n_it || i18n_all

package i18n

// 注册意大利语配置
func init() {
	RegisterLocale(Italian, &Locale{
		Language:     Italian,
		Region:       "IT",
		Name:         "Italiano",
		EnglishName:  "Italian",
		Messages: map[string]string{
			// Messaggi generali
			"error":   "Errore",
			"warning": "Avviso",
			"info":    "Informazione",
			"success": "Successo",
			"failed":  "Fallito",
			"loading": "Caricamento...",
			"saving":  "Salvataggio...",
			"done":    "Fatto",
			"cancel":  "Annulla",
			"confirm": "Conferma",
			"yes":     "Sì",
			"no":      "No",
			"ok":      "OK",
			
			// Tempo relativo
			"just_now":      "Proprio ora",
			"seconds_ago":   "%d secondi fa",
			"minutes_ago":   "%d minuti fa",
			"hours_ago":     "%d ore fa",
			"days_ago":      "%d giorni fa",
			"weeks_ago":     "%d settimane fa",
			"months_ago":    "%d mesi fa",
			"years_ago":     "%d anni fa",
			"seconds_later": "Tra %d secondi",
			"minutes_later": "Tra %d minuti",
			"hours_later":   "Tra %d ore",
			"days_later":    "Tra %d giorni",
			"weeks_later":   "Tra %d settimane",
			"months_later":  "Tra %d mesi",
			"years_later":   "Tra %d anni",
			
			// Messaggi di errore di validazione
			"required":        "%s è richiesto",
			"email":           "%s deve essere un indirizzo email valido",
			"url":             "%s deve essere un URL valido",
			"min":             "Il valore minimo di %s è %s",
			"max":             "Il valore massimo di %s è %s",
			"len":             "La lunghezza di %s deve essere di %s caratteri",
			"mobile":          "%s deve essere un numero di telefono cellulare valido",
			"idcard":          "%s deve essere un numero di carta d'identità valido",
			"bankcard":        "%s deve essere un numero di carta bancaria valido",
			"chinese_name":    "%s deve essere un nome cinese valido",
			"strong_password": "%s deve essere una password forte (almeno 8 caratteri, contenente lettere maiuscole, minuscole, numeri e caratteri speciali)",
			
			// Rete
			"network_error":     "Errore di rete",
			"connection_failed": "Connessione fallita",
			"timeout":           "Timeout",
			"server_error":      "Errore del server",
			"not_found":         "Non trovato",
			"unauthorized":      "Non autorizzato",
			"forbidden":         "Accesso vietato",
			
			// Operazioni sui file
			"file_not_found":    "File non trovato",
			"file_read_error":   "Errore nella lettura del file",
			"file_write_error":  "Errore nella scrittura del file",
			"file_delete_error": "Errore nell'eliminazione del file",
			"file_create_error": "Errore nella creazione del file",
			
			// Database
			"db_connection_error": "Errore di connessione al database",
			"db_query_error":      "Errore nella query del database",
			"db_insert_error":     "Errore nell'inserimento nel database",
			"db_update_error":     "Errore nell'aggiornamento del database",
			"db_delete_error":     "Errore nell'eliminazione dal database",
			
			// Operazioni generali
			"create": "Crea",
			"read":   "Leggi",
			"update": "Aggiorna",
			"delete": "Elimina",
			"list":   "Lista",
			"search": "Cerca",
			"filter": "Filtra",
			"sort":   "Ordina",
		},
		Formats: &Formats{
			DateFormat:        "02/01/2006",
			TimeFormat:        "15:04:05",
			DateTimeFormat:    "02/01/2006 15:04:05",
			NumberFormat:      "%.2f",
			CurrencyFormat:    "€ %.2f",
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
					"second":      "secondo",
					"minute":      "minuto",
					"hour":        "ora",
					"day":         "giorno",
					"week":        "settimana",
					"month":       "mese",
					"year":        "anno",
					"seconds":     "secondi",
					"minutes":     "minuti",
					"hours":       "ore",
					"days":        "giorni",
					"weeks":       "settimane",
					"months":      "mesi",
					"years":       "anni",
				},
				DistanceUnits: []string{"mm", "cm", "m", "km"},
				WeightUnits:   []string{"g", "kg", "t"},
			},
		},
	})
}